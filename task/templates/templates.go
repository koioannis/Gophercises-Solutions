package templates

const RootUsageTemplate = `{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}
{{if gt (len .Aliases) 0}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`

const SubCommandTemplate = `Usage:{{if .Runnable}}
{{.UseLine}}{{end}}
{{if .HasExample}}
Examples:
{{.Example}}{{end}}`

const AddExample = `
$ task add something important!
Added "review talk proposal" to your task list.

`

const DoExample = `
$ task do 1
You have completed the "review talk proposal" task.

`

const ListExample = `
$ task list
You have the following tasks:
1. some task description
2. some other task description

`

// `Usage:{{if .Runnable}}
//   {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
//   {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

// Aliases:
//   {{.NameAndAliases}}{{end}}{{if .HasExample}}

// Examples:
// {{.Example}}{{end}}{{if .HasAvailableSubCommands}}

// Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
//   {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

// Flags:
// {{.LocalFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

// Global Flags:
// {{.InheritedFlags.FlagUsages | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

// Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
//   {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

// Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
// `
