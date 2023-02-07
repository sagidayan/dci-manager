package cmd

import (
	"fmt"

	"github.com/rh-ecosystem-edge/dci-manager/config"
	"github.com/rh-ecosystem-edge/dci-manager/internal"
	"github.com/spf13/cobra"
)

var state = &cobra.Command{
	Use:   "state",
	Short: "Print the current matrix",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := config.Parse(internal.StateFilePath)
		if err != nil {
			fmt.Println("No state found.")
			return nil
		}
		fmt.Println(config.MatrixString())
		return nil
	},
}

func init() {
	rootCmd.AddCommand(state)
}
