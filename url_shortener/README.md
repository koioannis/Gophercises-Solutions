
# URL Shortener  ![language](https://img.shields.io/github/go-mod/go-version/koioannis/Gophercises-Solutions?filename=url_shortener%2Fgo.mod)


A simple url shortener that that gets urls from a yson or yaml file. For more detailed information you should check [Gophercises](https://gophercises.com/).

  

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
go build .\cmd\url_shortener\
./url_shortener.exe
```
or

```sh
go run .\cmd\url_shortener\
```

  

## :rocket: Usage
### Make sure you have the yson or yaml file under the `data` folder.


Yaml file should look like:
```yaml
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /example
  url: https://example.com
```
Json file should look like:
```json
[
    {
        "path": "/urlshort",
        "url": "https://github.com/gophercises/urlshort"
    },
    {
        "path": "/example",
        "url": "https://example.com"
    }
]
```
### Run the app
**Note:** You can use either json or yaml files with the same `-i` flag, just make sure that you specify the file extension.

If you were to use the go run command.
```sh
go run .\cmd\url_shortener\ -i youfile.json
```
or 
```sh
go run .\cmd\url_shortener\ -i youfile.yaml
```
And with the build command
```sh
go build .\cmd\url_shortener\
.\url_shortener.exe -i yourfile.json
```
If you do not pass the `-i` flag or the file could not be found, 
a default (hard coded) map with some urls will be used.



  

