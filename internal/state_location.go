package internal

import (
	"log"
)

func init() {
	// Ctreate state folder if not exists
	err := CreateFolderIfNotExists(StateFolderLocation)
	if err != nil {
		log.Fatalf("Error. failed to create state folder. %s", err.Error())
	}
}
