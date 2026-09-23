package subscription

import "github.com/spf13/cobra"

func listCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "list",
		Short: "List subscriptions",
		RunE: func(command *cobra.Command, arguments []string) error {
			// TODO: list subscriptions
			return nil
		},
	}

	return command
}
