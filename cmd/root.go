package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func GetAlienInvasionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                "alien-invasion",
		Short:              `Mad aliens are about to invade the earth and this program is to simulate the invasion.`,
		DisableSuggestions: true,
	}

	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.AddCommand(CmdInvade())

	return cmd
}

func Execute() {
	if err := GetAlienInvasionCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
