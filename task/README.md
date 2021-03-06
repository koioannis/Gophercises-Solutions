
# Choose your own adventure  ![language](https://img.shields.io/github/go-mod/go-version/koioannis/Gophercises-Solutions)


A CLI tool that manages TODOs. Created with Cobra and BoltDB

## :books: Table of Contents

  

- [Installation](#package-installation)

- [Usage](#rocket-usage)

  

## :package: Installation

  
### Clone the repo
```sh
git clone https://github.com/koioannis/Gophercises-Solutions
```


### Make sure you have [Go](https://golang.org) installed


```sh

go version

```

  

If you get an answer like this, it means that `Go` is installed.

```sh
go version go1.16.2 windows/amd64
```

### Build  the app
You can either build the app and run the executable or run the app via the go run command


```sh
go build .\cmd\task\
./task.exe
```
or

```sh
go run .\cmd\task\
```

  

## :rocket: Usage
Basic usage:

```bash
$ task
Tasks is a CLI for managing your TODOs


  task [command]


Available Commands:
  add         Adds a new task to the list
  do          Marks a TODO as completed
  help        Help about any command
  list        Lists your current Todos

Flags:
  -h, --help   help for task

Use "task [command] --help" for more information about a command.

$ task add review talk proposal
Added "review talk proposal" to your task list.

$ task add clean dishes
Added "clean dishes" to your task list.

$ task list
1. review talk proposal
2. clean dishes

$ task do 1
You have completed the "review talk proposal" task

$ task list 
2. clean dishes
```

### Run the app
If you were to use the go run command.
```sh
go run .\cmd\task [command]
```
And with the build command
```sh
go build .\cmd\task
.\task.exe [command]
```

You can also try:
```sh
go install .\cmd\task
task [command]
```



  

