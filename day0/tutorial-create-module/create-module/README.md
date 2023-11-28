# Documentation: Tutorial - Create a Go module

This is the first part of a tutorial that introduces a few fundamental features of the Go language. If you're just getting started with Go, be sure to take a look at Tutorial: Get started with Go, which introduces the go command, Go modules, and very simple Go code.

> source of the code **[here](https://go.dev/doc/tutorial/create-module)**

## How to run the projects?

To run the projects, you need to have the Go installed on your machine. If you don't have it, you can download it [here](https://golang.org/dl/).

This example is not complete so we cannot run yet.

## What I studied here?

- [x] Preface
  - [x] Create a new module
  - [x] Declare a greetings package to collect related functions.
  - [x] Implement a Hello function to return the greeting (Exported name).
  - [x] Declaring and initializing a variable (:= operator).
- [x] Commands
  - [x] `go mod init`
  - [x] `go run`
- [x] Program Structure
  - [x] Defining an exported name function `func Hello (){}`
  - [x] Defining params types `func Hello(name string){}`
  - [x] Defining return types `func Hello(name string) string {}`
  - [x] Declaring and initializing a variable
  ```go
        var message string
        message = fmt.Sprintf("Hi, %v. Welcome!", name)
  ```
  - [x] Declaring and initializing a variable shortcut `:=`
