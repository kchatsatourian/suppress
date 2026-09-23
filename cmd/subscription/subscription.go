package subscription

import "github.com/spf13/cobra"

// Command returns a cobra command for `subscription` subcommands
func Command() *cobra.Command {
	command := &cobra.Command{
		Use:   "subscription",
		Short: "Manage subscriptions",
	}

	command.AddCommand(
		addCommand(),
		listCommand(),
		removeCommand(),
	)

	return command
}
