package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"task/db"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Competes and removes selected tasks from your list.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg)
			} else {
				ids = append(ids, id)
			}
		}
		tasks, err := db.ListTasks()
		if err != nil {
			fmt.Printf("Something went wrong: %q\n", err)
			return
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number", id)
				continue
			}

			task := tasks[id-1]
			err := db.RemoveTask(task.Key)
			if err != nil {
				fmt.Printf("Something went wrong: %q\n", err)
				return
			} else {
				fmt.Printf("Marked \"%v\" as completed.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
