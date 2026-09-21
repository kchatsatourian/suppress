package state

import (
	"database/sql"
	"log/slog"
	"os"
	"time"

	_ "modernc.org/sqlite"
)

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

	_, err = SQLite.Exec(`
		CREATE TABLE IF NOT EXISTS state (
			key TEXT PRIMARY KEY,
			value TEXT NOT NULL
		)
	`)

	if err != nil {
		slog.Error("Could not initialize database.", "error", err)
		os.Exit(1)
	}
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
