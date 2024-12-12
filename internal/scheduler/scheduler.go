package scheduler

import (
	"log/slog"
	"os"

	"github.com/kchatsatourian/suppress/internal/suppress"
	CRON "github.com/robfig/cron/v3"
)

func Initialize() {
	schedule, exists := os.LookupEnv("SCHEDULE")
	if !exists {
		suppress.Execute()
		return
	}

	job := CRON.New()
	_, err := job.AddFunc(schedule, func() {
		suppress.Execute()
	})
	if err != nil {
		slog.Error("Could not create a job with schedule.", "schedule", schedule, "error", err)
		os.Exit(1)
	}

	slog.Info("Job created with schedule.", "schedule", schedule)
	job.Start()

	select {}
}
