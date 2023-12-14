package cli

import (
	"github.com/spf13/cobra"
)

var cmdReadme = &cobra.Command{
	Use:   "readme",
	Short: "Generate all readme files",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(cmdReadme)
}
