package task

import (
	"fmt"
	"os"

	"koioannis/gophercises/task/templates"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Tasks is a CLI for managing your TODOs",

	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.SetUsageTemplate(templates.RootUsageTemplate)


	addCmd.SetUsageTemplate(templates.SubCommandTemplate)
	addCmd.Example = templates.AddExample

	doCmd.SetUsageTemplate(templates.SubCommandTemplate)
	doCmd.Example = templates.DoExample

	listCmd.SetUsageTemplate(templates.SubCommandTemplate)
	listCmd.Example = templates.ListExample
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
