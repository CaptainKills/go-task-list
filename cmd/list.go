package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"

	"github.com/mergestat/timediff"
	"github.com/spf13/cobra"

	"github.com/CaptainKills/go-task-list/database"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the tasks in the database",
	Long: `The 'list' should return a list of all of the uncompleted tasks. 
It also has the option to return all tasks regardless of whether or not they are completed.

$ task list`,
	Run: list,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func list(cmd *cobra.Command, args []string) {
	records := []database.Task{
		{Id: 0, Name: "My first Task", CreatedAt: time.Now(), IsComplete: false},
		{Id: 1, Name: "My second Task", CreatedAt: time.Now(), IsComplete: true},
	}

	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 0, '\t', 0)

	fmt.Fprintln(writer, "ID\tTask\tCreated\tDone")
	for i := 0; i < len(records); i++ {
		task := records[i]
		timeText := timediff.TimeDiff(task.CreatedAt)
		text := fmt.Sprintf("%d\t%s\t%s\t%t", task.Id, task.Name, timeText, task.IsComplete)

		fmt.Fprintln(writer, text)
	}

	writer.Flush()
}
