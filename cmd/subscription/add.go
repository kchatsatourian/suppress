package subscription

import (
	"github.com/spf13/cobra"
)

type addFeedOptions struct {
	channel []int64
}

type addGitHubOptions struct {
	channel []int64
	pattern string
}

func addCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "add",
		Short: "Add a subscription",
	}

	command.AddCommand(
		addFeedCommand(),
		addGitHubCommand(),
	)

	return command
}

func addFeedCommand() *cobra.Command {
	options := addFeedOptions{}

	command := &cobra.Command{
		Use:   "feed [URL]",
		Short: "Add a feed subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(command *cobra.Command, arguments []string) error {
			// TODO: create feed subscription
			return nil
		},
	}

	flags := command.Flags()

	flags.Int64SliceVar(&options.channel, "channel", nil, "Notification channel")

	return command
}

func addGitHubCommand() *cobra.Command {
	options := addGitHubOptions{}

	command := &cobra.Command{
		Use:   "github [OWNER]/[REPOSITORY]",
		Short: "Add a GitHub subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(command *cobra.Command, arguments []string) error {
			// TODO: create GitHub subscription
			return nil
		},
	}

	flags := command.Flags()

	flags.Int64SliceVar(&options.channel, "channel", nil, "Notification channel")
	flags.StringVar(&options.pattern, "pattern", "", "Regular expression for matching releases")

	return command
}
