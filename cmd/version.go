package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var version = "development"

func versionCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run: func(cmd *cobra.Command, arguments []string) {
			fmt.Printf("Version: %s\n", version)
		},
	}

	return command
}
