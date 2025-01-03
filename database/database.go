package database

import (
	"os"
)

var DatabaseFilePath string = "/tmp/database.csv"

func Create() error {
	file, err := os.Create(DatabaseFilePath)
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

func Destroy() error {
	err := os.Remove(DatabaseFilePath)
	if err != nil {
		return err
	}

	return nil
}

func Exists() bool {
	file, err := os.Open(DatabaseFilePath)
	file.Close()

	if err != nil {
		return false
	}

	return true
}
