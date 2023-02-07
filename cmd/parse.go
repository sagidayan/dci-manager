package cmd

import (
	"fmt"
	"log"

	"github.com/rh-ecosystem-edge/dci-manager/config"
	"github.com/rh-ecosystem-edge/dci-manager/internal"
	"github.com/spf13/cobra"
)

var parse = &cobra.Command{
	Use:   "parse",
	Short: "Parse the given config file",
	Run:   parseHandler,
}

func init() {
	rootCmd.AddCommand(parse)
}

func parseHandler(cmd *cobra.Command, args []string) {
	config, err := config.Parse(internal.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Config version %s\nMatrices:\n", config.ConfigVersion())
	fmt.Println(config.MatrixString())
}
