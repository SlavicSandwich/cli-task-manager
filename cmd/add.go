package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
	"task/db"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		res := strings.Join(args, " ")
		_, err := db.CreateTask(res)
		if err != nil {
			fmt.Printf("Something went wrong: %q\n", err)
			return
		}
		fmt.Printf("Added %q\n", res)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
