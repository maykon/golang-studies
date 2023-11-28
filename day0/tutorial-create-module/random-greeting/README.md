# Documentation: Tutorial - Return a random greeting

In this section, you'll change your code so that instead of returning a single greeting every time, it returns one of several predefined greeting messages.

> source of the code **[here](https://go.dev/doc/tutorial/random-greeting)**

## How to run the projects?

To run the projects, you need to have the Go installed on your machine. If you don't have it, you can download it [here](https://golang.org/dl/).

After installing Go, you can run the projects using the command:

```bash
go run hello/hello.go
```

## What I studied here?

- [x] Preface
  - [x] Import Go standard library `math/rand` package.
  - [x] Change the function `Hello` to call another internal function `randomFormat`.
  - [x] Create a function `randomFormat` that define a `slice` of messages.
  - [x] Use the `rand.Intn` to return an integer based on the `slice` length.
  - [x] Return the `slice` length using `len` function.
- [x] Commands
  - [x] `go run`
- [x] Program Structure
  - [x] define an internal function (`func randomFormat()`)
  - [x] defining a dynamic slice `[]string{ }`
  - [x] Access a slice value by index `formats[0]`
  - [x] Get the slice length `len(format)`
- [x] Standard packages
  - [x] `math/rand`
  - [x] `log`
