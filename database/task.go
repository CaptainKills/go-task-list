package database

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

const TimeFormatString string = time.RFC3339

type Task struct {
	Id         int64
	Name       string
	CreatedAt  time.Time
	IsComplete bool
}

func (t Task) String() string {
	// Task Elements
	id := strconv.FormatInt(t.Id, 10)
	name := t.Name
	createdAt := t.CreatedAt.Format(TimeFormatString)
	isComplete := strconv.FormatBool(t.IsComplete)

	text := []string{id, name, createdAt, isComplete}

	// Task to Text Conversion
	buffer := bytes.Buffer{}

	writer := csv.NewWriter(&buffer)
	err := writer.Write(text)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not convert Task to String!")
		return ""
	}

	writer.Flush()

	return buffer.String()
}
