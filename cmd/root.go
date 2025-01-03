package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A cli application for managing tasks in the terminal.",
	Long: `A cli application that is be able to perform crud operations on a data file of tasks. The operations should be as follows:

$ task add "My new task"
$ task list
$ task complete`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
