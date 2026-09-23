package cmd

import (
	"github.com/kchatsatourian/suppress/cmd/channel"
	"github.com/kchatsatourian/suppress/cmd/subscription"
	"github.com/kchatsatourian/suppress/internal/scheduler"
	"github.com/kchatsatourian/suppress/internal/state"
	"github.com/kchatsatourian/suppress/internal/telegram"
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	command := &cobra.Command{
		Use:     "suppress",
		Short:   "Suppress is a simple RSS application for Telegram written in Go.",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			state.Initialize()
			telegram.Initialize()
			scheduler.Initialize()
		},
	}

	command.SetVersionTemplate(`{{.Version}}` + "\n")

	command.AddCommand(
		channel.Command(),
		serveCommand(),
		subscription.Command(),
		versionCommand(),
	)

	return command
}
