package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task from the database",
	Long: `The 'delete' method should be implemented to delete a task from the data store:

$ task delete <taskId>`,
	Run: delete,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}

func delete(cmd *cobra.Command, args []string) {
	fmt.Println("delete command")
}
