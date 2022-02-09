# evolugo

[![Go Report Card](https://goreportcard.com/badge/github.com/jaroddev/evolugo)](https://goreportcard.com/report/github.com/jaroddev/evolugo)

evolugo is a new project and a genetic algorithm library in go. The goal is to provide
a starting point to solve problems with, few lines of code to write, pretty good performance in the long 
through parralelism anda lot more.

## Warning

The api will change as the project is currently being under its early development phase. 

## Installation

Using this package requires a working Go environment. [See the install instructions for Go](http://golang.org/doc/install.html). 
Go Modules are required when using this package. [See the go blog guide on using Go Modules](https://blog.golang.org/using-go-modules).

```
$ go get github.com/jaroddev/evolugo
```

```go
...
import (
  genetic "github.com/jaroddev/evolugo/genetic"
)
...
```

## Packages

You can get access to a basic genetic algorithm skeleton in the genetic package, but you can code your own one too.

You also get basic genetic algorithm components such as:
 - Selections
 - Mutations
 - Crossovers
 - Insertions

## Supported platforms

evolugo is tested against Linux, Windows, and against the 1.17
released version of Go.
