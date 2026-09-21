package tag

import (
	"database/sql"
	"log/slog"

	"github.com/kchatsatourian/suppress/internal/state"
)

func Read() string {
	var value string

	err := state.SQLite.QueryRow(
		"SELECT value FROM state WHERE key = 'tag'",
	).Scan(&value)

	if err == sql.ErrNoRows {
		return ""
	}

	if err != nil {
		slog.Warn("Could not read tag.", "error", err)
		return ""
	}

	return value
}

func Write(tag string) {
	_, err := state.SQLite.Exec(`
		INSERT INTO state (key, value)
		VALUES ('tag', ?)
		ON CONFLICT(key) DO UPDATE SET value = excluded.value
	`, tag)

	if err != nil {
		slog.Warn("Could not write tag.", "error", err)
	}
}
