package database

import (
	"log"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	err := Create()
	if err != nil {
		log.Fatalf("Could not create database: %q", err)
	}

	file, err := os.Open(DatabaseFilePath)
	if err != nil {
		log.Fatalf("Could not create database: %q", err)
	}

	file.Close()
}

func TestDestroy(t *testing.T) {
	err := Destroy()
	if err != nil {
		log.Fatalf("Could not destroy database: %q", err)
	}

	file, err := os.Open(DatabaseFilePath)
	if err == nil {
		log.Fatalf("Could not destroy database: %q", err)
	}

	file.Close()
}

func TestExists(t *testing.T) {
	t.Run("Database Exists", func(t *testing.T) {
		err := Create()
		if err != nil {
			log.Fatalf("Could not create database: %q", err)
		}

		got := Exists()
		want := true
		if got != want {
			log.Fatalf("got %t want %t", got, want)
		}
	})

	t.Run("Database Does Not Exist", func(t *testing.T) {
		err := Destroy()
		if err != nil {
			log.Fatalf("Could not destroy database: %q", err)
		}

		got := Exists()
		want := false
		if got != want {
			log.Fatalf("got %t want %t", got, want)
		}
	})
}
