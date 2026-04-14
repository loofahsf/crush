package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Aliases: []string{"auth"},
	Use:     "login [platform]",
	Short:   "Login Crush to a platform",
	Long:    `Login Crush to a specified platform.`,
	Args:    cobra.MaximumNArgs(1),
	RunE: func(_ *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("unknown platform: %s", args[0])
		}
		return fmt.Errorf("no login platforms available")
	},
}
