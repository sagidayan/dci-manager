package cmd

import (
	"fmt"
	"log"
	"os"

	v1 "github.com/rh-ecosystem-edge/dci-manager/config/v1"
	"github.com/spf13/cobra"
)

var generate = &cobra.Command{
	Use:   "generate <name>",
	Short: "Generate a config file",
	Run:   generateHandler,
	Args:  cobra.MinimumNArgs(1),
}

var format string

func init() {
	generate.Flags().StringVarP(&format, "format", "f", "yaml", "Output format. Can be json/yaml/yml. Defaults to yaml")
	rootCmd.AddCommand(generate)
}

func generateHandler(cmd *cobra.Command, args []string) {
	fileName := fmt.Sprintf("%s.%s", args[0], format)
	if format != "json" && format != "yaml" && format != "yml" {
		log.Fatal("Error: invalid output format.")
	}
	isJSON := format == "json"
	c := &v1.Config{}
	b, err := c.Sample(isJSON)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(fileName, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Success. Generated config file %s", fileName)

}
