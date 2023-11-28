# Documentation: Tutorial - Return and handle an error

Handling errors is an essential feature of solid code. In this section, you'll add a bit of code to return an error from the greetings module, then handle it in the caller.

> source of the code **[here](https://go.dev/doc/tutorial/handle-errors)**

## How to run the projects?

To run the projects, you need to have the Go installed on your machine. If you don't have it, you can download it [here](https://golang.org/dl/).

After installing Go, you can run the projects using the command:

```bash
go run hello/hello.go
```

## What I studied here?

- [x] Preface
  - [x] Import Go standard library `errors` package.
  - [x] Change the function to return two values.
  - [x] Add validation to check the name.
  - [x] If empty throw an error.
  - [x] Return two values a string and an error if happens
  - [x] Add the `log` package and configure it to not show some information
  - [x] Change the main program to get the new params
  - [x] If an error happens log it and stop the program
  - [x] Otherwise, print the name
- [x] Commands
  - [x] `go run`
- [x] Program Structure
  - [x] Return multiple values `func x() (string, error) {}`
  - [x] Return an error `return "", errors.New("empty name")`
  - [x] If operator `if`
  - [x] Compare operator `==`
  - [x] Returning a no error value `nil`
  - [x] Throw an error `errors.New`
  - [x] Set multiple values in the same line `message, err := greetings.Hello("")`
- [x] Standard packages
  - [x] `errors`
  - [x] `log`
