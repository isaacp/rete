package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of rete",
	Long:  `All software has versions. This is rete's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Rete Network Generator v0.1")
	},
}
