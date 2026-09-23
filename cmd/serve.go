package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func serveCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
	and usage of using your command. For example:

	Cobra is a CLI library for Go that empowers applications.
	This application is a tool to generate the needed files
	to quickly create a Cobra application.`,
		Run: func(command *cobra.Command, arguments []string) {
			fmt.Println("serve called")
		},
	}

	return command
}
