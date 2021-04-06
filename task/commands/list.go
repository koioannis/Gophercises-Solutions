package task

import (
	"fmt"
	"github.com/koioannis/gophercises-solutions/task/services"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists your current Todos",
	RunE: func(cmd *cobra.Command, args []string) error {
		ts, err := services.NewTaskService()
		if err != nil {
			return err
		}
		defer ts.Close()

		tasks, err := ts.GetAll()
		if err != nil {
			return err
		}

		for _, task := range tasks {
			fmt.Printf("%v. %v\n", task.ID, task.Content)
		}

		return nil
	},
}
