package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/rh-ecosystem-edge/dci-manager/config"
	"github.com/rh-ecosystem-edge/dci-manager/dci"
	"github.com/rh-ecosystem-edge/dci-manager/internal"
	"github.com/spf13/cobra"
)

var (
	autoApprove = false
)

var apply = &cobra.Command{
	Use:   "apply",
	Short: "Run diff on current state and see what will change. Approve to apply",
	Run:   applyHandler,
}

func init() {
	apply.Flags().BoolVarP(&autoApprove, "yes", "y", false, "To auto approve")
	rootCmd.AddCommand(apply)
}

func applyHandler(cmd *cobra.Command, args []string) {
	config, err := config.Parse(internal.ConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	diff, err := config.Plan()
	if err != nil {
		log.Fatal(err)
	}
	if diff.Len() == 0 {
		fmt.Println(diff)
		return
	}
	fmt.Println("CAUTION: These changes will take affect once you approve!")
	fmt.Println(diff)
	if autoApprove || YesNoPrompt("Do you want to apply these changes?") {
		//TODO: Implement
		b, err := json.Marshal(config)
		if err != nil {
			log.Fatal(err)
		}
		err = os.WriteFile(internal.StateFilePath, b, 0644)
		if err != nil {
			log.Fatal(err)
		}
		for _, j := range diff.Added() {
			if err := dci.CreateJob(j); err != nil {
				log.Fatal(err)
			}
		}
		for _, j := range diff.Removed() {
			if err := dci.RemoveJob(j); err != nil {
				log.Fatal(err)
			}
		}
		for _, j := range diff.Modified() {
			if err := dci.CreateJob(j); err != nil {
				log.Fatal(err)
			}
		}
		added := color.New(color.FgGreen).Sprintf("%d Added", len(diff.Added()))
		modified := color.New(color.FgYellow).Sprintf("%d Modified", len(diff.Modified()))
		removed := color.New(color.FgRed).Sprintf("%d Removed", len(diff.Removed()))
		fmt.Printf("Success! %s | %s | %s.\n", added, removed, modified)
		return
	}
	fmt.Println("Aborting")
}
