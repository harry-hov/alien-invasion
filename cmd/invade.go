package cmd

import (
	"fmt"
	"os"

	cmderror "github.com/harry-hov/alien-invasion/error"
	"github.com/harry-hov/alien-invasion/invasion"
	"github.com/harry-hov/alien-invasion/worldmap"
	"github.com/spf13/cobra"
)

func CmdInvade() *cobra.Command {
	var alienCount uint
	cmd := &cobra.Command{
		Use:   "invade [world-file]",
		Short: "Invade a World",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			filename := args[0]
			if filename == "" {
				return cmderror.Wrap(cmderror.ErrInvalidFileName, "")
			}
			if alienCount == 0 {
				return cmderror.Wrap(cmderror.ErrInvalidAlienCount, "invalid value (0) for [-a | --aliens] flag")
			}

			fp, err := os.Open(filename)
			if err != nil {
				return err
			}

			worldMap, err := worldmap.InitWorldMap(fp)
			if err != nil {
				return err
			}
			invasion := invasion.InitInvasion(worldMap, alienCount)
			fmt.Println(invasion.GetWorldMap())

			return nil
		},
	}

	cmd.Flags().UintVarP(&alienCount, "aliens", "a", 0, "Alien Count")

	return cmd
}
