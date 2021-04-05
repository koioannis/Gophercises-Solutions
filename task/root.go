package task

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task [command]",
	Short: "Tasks is a CLI adfor managing your TODOs",
	Long: `Available Commands
	add         Add a new task to your TODO list
	do          Mark a task on your TODO list as complete
	list        List all of your incomplete tasks

Use "task [command] --help" for more information about a command.`,

	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var testCmd = &cobra.Command{
	Use:   "test [test]",
	Short: "test",
	Long: `test`,

	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello!")
	},
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	rootCmd.AddCommand(testCmd)
}
