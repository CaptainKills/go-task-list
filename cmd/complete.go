package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Complete a task in the database",
	Long: `The 'complete' method should be complete tasks in the underlying data store.
It should take a positional argument with the task ID. For example:

$ task complete <taskId>`,
	Run: complete,
}

func init() {
	rootCmd.AddCommand(completeCmd)
}

func complete(cmd *cobra.Command, args []string) {
	fmt.Println("complete command")
}
