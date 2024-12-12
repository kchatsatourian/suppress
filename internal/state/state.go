package state

import (
	"log/slog"
	"os"
	"time"
)

var (
	ExecutedAt time.Time
	path       = "/suppress/configuration/updatedAt"
	UpdatedAt  time.Time
)

func Read() {
	bytes, err := os.ReadFile(path)
	if err != nil {
		slog.Warn("Could not read time.", "error", err)
		UpdatedAt = ExecutedAt
		return
	}
	UpdatedAt, err = time.Parse(time.RFC3339, string(bytes))
	if err != nil {
		slog.Warn("Could not deserialize time.", "error", err)
		UpdatedAt = ExecutedAt
	}
}

func Write() {
	bytes := []byte(ExecutedAt.Format(time.RFC3339))
	err := os.WriteFile(path, bytes, 0400)
	if err != nil {
		slog.Warn("Could not write time.", "error", err)
	}
}
