package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to the database",
	Long: `The 'add' method should be used to create new tasks in the underlying data store.
It should take a positional argument with the task description. For example:

$ task add "My new task"`,
	Run: add,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func add(cmd *cobra.Command, args []string) {
	fmt.Println("add command")
}
