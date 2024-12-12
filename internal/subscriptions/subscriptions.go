package subscriptions

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"

	"github.com/kchatsatourian/suppress/internal/tag"
)

var (
	path          = "/suppress/configuration/subscriptions.json"
	Subscriptions map[string]Subscription
)

func fetch() {
	link, exists := os.LookupEnv("GITHUB_GIST")
	if !exists {
		return
	}

	request, err := http.NewRequest("GET", link, nil)
	if err != nil {
		slog.Warn("Could not create HTTP request.", "error", err)
	}
	request.Header.Set("If-None-Match", tag.Read())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		slog.Warn("Could not fetch remote subscriptions.", "error", err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		tag.Write(response.Header.Get("ETag"))

		err = json.NewDecoder(response.Body).Decode(&Subscriptions)
		if err != nil {
			slog.Warn("Could not deserialize remote subscriptions.", "error", err)
		}

		write()
	}
}

func Read() {
	fetch()
	bytes, err := os.ReadFile(path)
	if err != nil {
		slog.Error("Could not read subscriptions.", "error", err)
		os.Exit(1)
	}
	err = json.Unmarshal(bytes, &Subscriptions)
	if err != nil {
		slog.Error("Could not deserialize subscriptions.", "error", err)
		os.Exit(1)
	}
}

func write() {
	bytes, err := json.Marshal(Subscriptions)
	if err != nil {
		slog.Warn("Could not serialize subscriptions.", "error", err)
	}
	err = os.WriteFile(path, bytes, 0400)
	if err != nil {
		slog.Warn("Could not write subscriptions.", "error", err)
	}
}
