package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"task/db"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Competes and removes selected tasks from your list.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Printf("Something went wrong: %q\n", err)
			return
		}
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
