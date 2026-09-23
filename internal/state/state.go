package state

import (
	"database/sql"
	"log/slog"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

const expiration = 30 * 24 * time.Hour

var (
	ExecutedAt time.Time
	path       = "/suppress/configuration/suppress.db"
	SQLite     *sql.DB
	UpdatedAt  time.Time
)

func Initialize() {
	var err error

	SQLite, err = sql.Open("sqlite", path)
	if err != nil {
		slog.Error("Could not open database.", "error", err)
		os.Exit(1)
	}

	_, err = SQLite.Exec(schema)
	if err != nil {
		slog.Error("Could not create schema.", "error", err)
		os.Exit(1)
	}

	_, err = SQLite.Exec(`
		DELETE FROM deduplication
		WHERE expires_at <= ?
	`, time.Now().Unix())
	if err != nil {
		slog.Error("Could not remove expired entries.", "error", err)
		os.Exit(1)
	}
}

func IsUnique(link string) bool {
	now := time.Now().Unix()

	result, err := SQLite.Exec(`
		INSERT INTO deduplication (link, expires_at)
		VALUES (?, ?)
		ON CONFLICT(link) DO NOTHING
	`, link, now+int64(expiration.Seconds()))

	if err != nil {
		slog.Warn("Could not check deduplication.", "error", err)
		return true
	}

	rows, err := result.RowsAffected()
	if err != nil {
		slog.Warn("Could not check deduplication result.", "error", err)
		return true
	}

	return rows > 0
}

func Read() {
	var value string

	err := SQLite.QueryRow(`
		SELECT value
		FROM state
		WHERE key = 'updated_at'
	`).Scan(&value)

	if err == sql.ErrNoRows {
		UpdatedAt = ExecutedAt
		return
	}

	if err != nil {
		slog.Warn("Could not read time.", "error", err)
		UpdatedAt = ExecutedAt
		return
	}

	UpdatedAt, err = time.Parse(time.RFC3339, value)
	if err != nil {
		slog.Warn("Could not deserialize time.", "error", err)
		UpdatedAt = ExecutedAt
	}
}

func Write() {
	_, err := SQLite.Exec(`
		INSERT INTO state (key, value)
		VALUES ('updated_at', ?)
		ON CONFLICT(key) DO UPDATE SET value = excluded.value
	`, ExecutedAt.Format(time.RFC3339))

	if err != nil {
		slog.Warn("Could not write time.", "error", err)
	}
}
