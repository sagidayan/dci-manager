package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func YesNoPrompt(s string) bool {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%s [yes/N]: ", s)

	response, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	if response == "yes\n" {
		return true
	}

	return false
}
