package subscription

import (
	"strconv"

	"github.com/spf13/cobra"
)

func removeCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "remove [ID]",
		Short: "Remove a subscription",
		Args:  cobra.ExactArgs(1),
		RunE: func(command *cobra.Command, arguments []string) error {
			id, err := strconv.ParseInt(arguments[0], 10, 64)
			if err != nil {
				return err
			}

			// TODO: remove subscription
			_ = id

			return nil
		},
	}

	return command
}
