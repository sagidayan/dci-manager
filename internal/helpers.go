package internal

import (
	"fmt"
	"os"
)

func CreateFolderIfNotExists(location string) error {
	_, err := os.Stat(location)
	if os.IsNotExist(err) {
		err := os.MkdirAll(location, 0775)
		if err != nil {
			return fmt.Errorf("Error. failed to create state folder. %s", err.Error())
		}
	}
	return nil
}
