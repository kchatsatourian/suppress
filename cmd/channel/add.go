package channel

import (
	"github.com/spf13/cobra"
)

type addOptions struct {
	subscriptionType string
}

func addCommand() *cobra.Command {
	options := addOptions{}

	command := &cobra.Command{
		Use:   "add [URL]",
		Short: "Add a notification channel",
		Args:  cobra.ExactArgs(1),
		RunE: func(command *cobra.Command, arguments []string) error {
			// TODO: create channel
			return nil
		},
	}

	flags := command.Flags()

	flags.StringVar(&options.subscriptionType, "type", "", "Subscription type")

	return command
}
