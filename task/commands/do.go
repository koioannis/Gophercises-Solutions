package task

import (
	"fmt"
	"koioannis/gophercises/task/services"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a TODO as completed",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			cmd.Usage()
			return nil
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			cmd.Usage()
			return nil
		}

		ts, err := services.NewTaskService()
		if err != nil {
			return err
		}
		defer ts.Close()

		task, err := ts.Remove(id)
		if err != nil {
			return err
		}

		if task.Content == "" {
			fmt.Printf("Task not found. Try:\n$ task list\n")
			return nil
		}

		fmt.Printf("You have completed the \"%v\" task", task.Content)

		return nil
	},
}
