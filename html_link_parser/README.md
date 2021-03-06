
# HTML Link Parser  ![language](https://img.shields.io/github/go-mod/go-version/koioannis/Gophercises-Solutions)


Parses the HTML tree returned by [x/net/html](https://pkg.go.dev/golang.org/x/net/html) using DFS, and extracts all the `a` tags into a json. 
For example, given the code below
```html
...
<a href="/dog">
  <span>Something in a span</span>
  Text not in a span
  <b>Bold text!</b>
</a>
<a href="/koioannis">My username</a>
<a href="https://golang.org">Go is awesome!</a>
...
```
The output looks like
```json
[
  {
    "href": "/dog",
    "text": "Something in a span Text not in a span Bold text!"
  },
  {
    "href": "/koioannis",
    "text": "My username"
  },
  {
    "href": "https://golang.org",
    "text": "Go is awesome!"
  }
]

```

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
go build .\cmd\html_link_parser\
./html_link_parser.exe
```
or

```sh
go run .\cmd\html_link_parser\
```

  

## :rocket: Usage
### Make sure you have the html file under the `data` folder.

### Run the app
If you were to use the go run command.
```sh
go run .\cmd\html_link_parser\ -i input.html -o output.json
```
And with the build command
```sh
go build .\cmd\html_link_parser\
.\html_link_parser.exe -i input.html -o output.json
```

Each `a` tag is treated on it's own. Nested `a` tag's text won't be included on the parent's text. For example
```html
<a href="/parent">This is text<a href="/nested">This is nested text</a></a>
```
Should return
```json
[
  {
    "href":"/parent",
    "text": "This is text"
  },
  {
    "href":"/nested",
    "text":"This is nested text"
  }
]
```
You might argue it's a bug, and you'd probably be right, but I've decided to call it a feature. :sweat_smile:

  

