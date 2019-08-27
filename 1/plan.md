# Go Lunch Week One
## A Blogging Server
Some features we want include:
 - blog-post crud
 - blog-posts persistent between server restarts

## This Week
### Repo Setup
- `git clone https://github.com/jcc333/golunch # or go to github and fork the repo`
- `cd golunch/1`
### Project Setup
- `mkdir -p cmd/blagsrv`
- `touch cmd/blagsrv/main.go`
- `go mod init github.com/jcc333/golunch # or your fork, whatever`
- `go build ./... # build all of the .go files you can find in this dir and its children`
- `mkdir internal # where your library files live`
- `curl https://golang.org/doc/articles/wiki/final.go?m=text > cmd/blagsrv/main.go`
### Let's Refactor
#### Backend Only
- The `"encoding/json"` [package](https://golang.org/pkg/encoding/json/)
#### Organization
- a `blagsrv` package
- a `BlagHandler` type [http.Handler](https://golang.org/pkg/net/http/#Handler)
- what's an `interface`?
- what's a `struct`?
- adding logs with [logrus](https://github.com/Sirupsen/logrus)
