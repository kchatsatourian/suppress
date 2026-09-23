package channel

import (
	"github.com/spf13/cobra"
)

func listCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "list",
		Short: "List notification channels",
		RunE: func(command *cobra.Command, arguments []string) error {
			// TODO: list channels
			return nil
		},
	}

	return command
}
