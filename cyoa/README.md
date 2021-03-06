
# Choose your own adventure  ![language](https://img.shields.io/github/go-mod/go-version/koioannis/Gophercises-Solutions)


An HTML text based adventure game based on [Choose Your Own Adventure](https://en.wikipedia.org/wiki/Choose_Your_Own_Adventure).

  

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
go build .\cmd\cyoa\
./cyoa.exe
```
or

```sh
go run .\cmd\cyoa\
```

  

## :rocket: Usage
### Make sure you have the yson file under the `data` folder.

The adventure game is based on a `JSON` file. The basic structure is:
```jsonc  
{
  // Each story arc will have a unique key that represents
  // the name of that particular arc.
  "story-arc": {
    "title": "A title for that story arc. Think of it like a chapter title.",
    "story": [
      "A series of paragraphs, each represented as a string in a slice.",
      "This is a new paragraph in this particular story arc."
    ],
    // Options will be empty if it is the end of that
    // particular story arc. Otherwise it will have one or
    // more JSON objects that represent an "option" that the
    // reader has at the end of a story arc.
    "options": [
      {
        "text": "the text to render for this option. eg 'venture down the dark passage'",
        "arc": "the name of the story arc to navigate to. This will match the story-arc key at the very root of the JSON document"
      }
    ]
  },
  ...
}
```

### Run the app
If you were to use the go run command.
```sh
go run .\cmd\cyoa\ -i youfile.json
```
And with the build command
```sh
go build .\cmd\cyoa\
.\cyoa.exe -i yourfile.json
```



  

