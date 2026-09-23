package channel

import "github.com/spf13/cobra"

// Command returns a cobra command for `channel` subcommands
func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "channel",
		Short: "Manage notification channels",
	}

	command.AddCommand(
		addCommand(),
		listCommand(),
		removeCommand(),
	)

	return command
}
