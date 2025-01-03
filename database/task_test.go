package database

import (
	"log"
	"testing"
	"time"
)

func TestString(t *testing.T) {
	t.Run("Task To String", func(t *testing.T) {
		currentTime := time.Now()
		task := Task{0, "New Task", currentTime, false}

		got := task.String()
		want := "0,New Task," + currentTime.Format(TimeFormatString) + ",false\n"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}

func assertString(t testing.TB, got, want []string) {
	t.Helper()

	if len(got) == 0 || len(want) == 0 {
		log.Fatalf("String Array Empty: got %q want %q", got, want)
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			log.Fatalf("got %q want %q", got, want)
		}
	}
}
