package cmd

import (
	"fmt"
	"log"

	"github.com/rh-ecosystem-edge/dci-manager/config"
	"github.com/rh-ecosystem-edge/dci-manager/internal"
	"github.com/spf13/cobra"
)

var plan = &cobra.Command{
	Use:   "plan",
	Short: "Run diff on current state and see what will change",
	Run:   planHandler,
}

func init() {
	rootCmd.AddCommand(plan)
}

func planHandler(cmd *cobra.Command, args []string) {
	config, err := config.Parse(internal.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	diff, err := config.Plan()
	if err != nil {
		log.Fatal(fmt.Errorf("Planing Error: %s", err))
	}
	fmt.Println(diff)
}
