# Notes on Week One
## Go modules
- We're only using modules to avoid gopath stuff. If that means nothing to you, rejoice: you missed some terrible stuff.
- `go.mod` is like a `Gemfile`
- `go.sum` is like a `Gemfile.lock` with checksums for your vendored packages
##  Go tools
- `mkdir -p cmd/blagsrv`
- `touch cmd/blagsrv/main.go`
- `go mod init github.com/jcc333/golunch # or your fork, whatever`
- `go build ./... # build all of the .go files you can find in this dir and its children`
- `mkdir internal # where your library files live`
- ``
## Main
- a minimal go program
```
package main

import "fmt"

func main() {
    fmt.Println("Helgo, world")
}
```
## fmt.Println
- `fmt` is a package we imported. To refer to `Println` we have to index it to the package: `fmt.Println`
- [fmt.Println](https://golang.org/pkg/fmt/#Println)
- "This is a function named Printlin which takes any amount of arguments of any type at all"
- Things which start in Uppercase runes are package-public, things which do not are package-private
- The `...` means, a comma delimited list of whatever type
- To cast a `[]int` into a `...int`, do `ints...`, where `ints []int`
## Interfaces
- The `interface{}` means, any type which satisfies all of the methods between the braces. In the case of `interface{}`, there are no methods to satisfy so all types satisfy `interface{}`
- One `interface` that's more interesting is [`error`](https://golang.org/ref/spec#Errors)
- `interface error { Error() string }` is a built-in, so it can be called `error` and still be public to every package
- Types don't inherit interfaces, they satisfy them
- For example, `type MyCoolError struct { ErrorMessage string; LineNumber int }` does not satisfy error until we define:
- `func (mce MyCoolError) Error() string { return fmt.Sprintln("error at %d: %s", mce.LineNumber, mce.ErrorMessage) }`
- This is function on `MyCoolError`s called `Error` which takes no arguments and returns a string
- When we need an error type, we can now use `MyCoolError` and the compiler will look at the struct we provide and see it has a function on it called `Error() string`. The compiler knows what public values are defined for each type and in each package and uses them to figure out if the type satisfies an interface.
## Multiple Returns, Named Return Values
- So `fmt.Println` returns `(n int, err error)`: `func Println(a ...interface{}) (n int, err error)`
- That means, this is a function from an arbitrary number of arbitrary arguments to an `int` called `n` and an `error` called `err`
- This is just like "out" arguments in c/c++/etc. (_kind of_ [chrono::time](https://en.cppreference.com/w/c/chrono/time) but there are better examples)
- In practice, when you see `(something int, somethingElse string, err error)`, you can usually assume that `int` and `string` are only valid when `err == nil`
## Types are Names
- If we define `type MyCoolerError MyCoolError`, we're telling the compiler, `MyCoolerError` looks exactly like `MyCoolError` but it's different.
- They'll have the same properties but different available methods
- You have to cast them to one another to get to the functions one or the other has: `MyCoolerError(MyCoolError{ ErrorMessage: "something bad", LineNumber: 0 })` has all of the `MyCoolerError` functions and none of the `MyCoolError` functions.
- You can also define different versions of the same function for type aliases:
```
import "encoding/json"

// this is a type called MyCoolError which is a struct with a field, ErrorMessage that is a string, and a field LineNumber that is an int
// a `struct` is a sized array of named fields of varrying types
// an `int` is *platform dependent* but in practice is almost always a signed 64-bit integer
type  MyCoolError struct { ErrorMessage string; LineNumber int }

type MyCoolerError MyCoolError

// A
func (mce MyCoolError) Error() string {
    return fmt.Sprintln("error at %d: %s", mce.LineNumber, mce.ErrorMessage)
}

// B
func (mce MyCoolerError) Error() string {
    bytes, __ := json.Marshal(mce) // returns ([]byte, error) but we're ignoring the error for now
    return string(bytes)
}

func main() {
    mce := MyCoolError{ ErrorMessage: "oh no", LineNumber: 42 }
    fmt.Println(mce.Error()) // calls A. If the 'A' version of `Error` weren't defined this would be a compiler error
    fmt.Println(MyCoolerError(mce).Error()) // calls B. If the 'B' version of `Error` weren't defined this would be a compiler error
}
```

## Functions are Values
- Functions are passable as values and enclose their stackframe-local variables
- Functions can be fields in structs or arguments to functions or local variables themselves:
```
func main() {
    mce := MyCoolError{ ErrorMessage: "oh no", LineNumber: 42 }
    fmt.Println(mce.Error()) // calls A. If the 'A' version of `Error` weren't defined this would be a compiler error
    fmt.Println(MyCoolerError(mce).Error()) // calls B. If the 'B' version of `Error` weren't defined this would be a compiler error
    apply = func(m MyCoolError, f func(MyCoolError) string) string { return f(m) } // calls `f` on `m`
    fmt.Println(apply(mce, MyCoolError.Error)) // calls the version of `Error` that belongs to the type `MyCoolError`

}
```
