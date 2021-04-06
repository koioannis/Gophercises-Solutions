package task

import (
	"fmt"
	"github.com/koioannis/gophercises-solutions/task/services"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:                "add",
	Short:              "Adds a new task to the list",
	DisableFlagParsing: true,
	SilenceUsage:       true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
			return nil
		}
		ts, err := services.NewTaskService()
		if err != nil {
			return err
		}
		defer ts.Close()

		content := strings.Join(args, " ")
		task := services.Task{
			Content: content,
		}

		if err := ts.Create(&task); err != nil {
			return err
		}

		fmt.Printf("Added \"%v\" to your task list", task.Content)

		return nil
	},
}
