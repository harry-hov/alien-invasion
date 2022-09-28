package cmd

import (
	"fmt"
	"os"

	cmderror "github.com/harry-hov/alien-invasion/error"
	"github.com/harry-hov/alien-invasion/worldmap"
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
				return cmderror.Wrap(cmderror.ErrInvalidFileName, "")
			}

			fp, err := os.Open(filename)
			if err != nil {
				return err
			}

			worldMap, err := worldmap.InitWorldMap(fp)
			if err != nil {
				return err
			}

			fmt.Println(worldMap)

			return nil
		},
	}

	return cmd
}
