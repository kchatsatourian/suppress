package tag

import (
	"log/slog"
	"os"
)

var (
	path = "/suppress/configuration/tag"
)

func Read() string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		slog.Warn("Could not read tag.", "error", err)
		return ""
	}
	return string(bytes)
}

func Write(tag string) {
	bytes := []byte(tag)
	err := os.WriteFile(path, bytes, 0400)
	if err != nil {
		slog.Warn("Could not write tag.", "error", err)
	}
}
