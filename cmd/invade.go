package cmd

import (
	"fmt"

	cmderror "github.com/harry-hov/alien-invasion/error"
	"github.com/spf13/cobra"
)

func CmdInvade() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "invade [world-file]",
		Short: "Invade a World",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filename := args[0]
			if filename == "" {
				return cmderror.InvalidFileName
			}

			fmt.Println("Invaded")

			return nil
		},
	}

	return cmd
}
