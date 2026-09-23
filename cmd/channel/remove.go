package channel

import (
	"strconv"

	"github.com/spf13/cobra"
)

func removeCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "remove [ID]",
		Short: "Remove a notification channel",
		RunE: func(command *cobra.Command, arguments []string) error {
			id, err := strconv.ParseInt(arguments[0], 10, 64)
			if err != nil {
				return err
			}

			// TODO: remove channel
			_ = id

			return nil
		},
	}

	return command
}
