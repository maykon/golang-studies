# Documentation: Tutorial - Call your code from another module

In the previous section, you created a greetings module. In this section, you'll write code to make calls to the Hello function in the module you just wrote. You'll write code you can execute as an application, and which calls code in the greetings module.

> source of the code **[here](https://go.dev/doc/tutorial/call-module-code)**

## How to run the projects?

To run the projects, you need to have the Go installed on your machine. If you don't have it, you can download it [here](https://golang.org/dl/).

After installing Go, you can run the projects using the command:

```bash
go run hello/hello.go
```

## What I studied here?

- [x] Preface
  - [x] Create a hello directory for your Go module source code.
  - [x] Enable dependency tracking for the code.
  - [x] Import the external package `mk/greetings`.
  - [x] Edit the module to apointing to correct path to `mk/greetings`.
  - [x] Synchronize the module's dependencies
  - [x] Run the example
- [x] Commands
  - [x] `go mod init`
  - [x] `go mod edit -replace`
  - [x] `go mod tidy`
  - [x] `go run`
- [x] Program Structure
  - [x] Call exported name func `greetings.Hello("Maykon")`
