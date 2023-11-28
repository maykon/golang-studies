# Documentation: Tutorial - Return greetings for multiple people

In the last changes you'll make to your module's code, you'll add support for getting greetings for multiple people in one request. In other words, you'll handle a multiple-value input, then pair values in that input with a multiple-value output. To do this, you'll need to pass a set of names to a function that can return a greeting for each of them.

> source of the code **[here](https://go.dev/doc/tutorial/greetings-multiple-people)**

## How to run the projects?

To run the projects, you need to have the Go installed on your machine. If you don't have it, you can download it [here](https://golang.org/dl/).

After installing Go, you can run the projects using the command:

```bash
go run hello/hello.go
```

## What I studied here?

- [x] Preface
  - [x] Hints for backward compatibility.
  - [x] Add a `Hellos` function whose parameter is a slice of names rather than a single name.
  - [x] Return types from a string to a map so you can return names mapped to greeting messages.
  - [x] Reuse the old `Hello` function to avoid duplication.
  - [x] Create a messages map.
  - [x] Populate the map using `for loop`.
  - [x] Using `range` to return the names from the `slice`.
  - [x] Create a `slice` with the names and call the new `Hellos` function.
  - [x] Print the map of messages.
- [x] Commands
  - [x] `go run`
- [x] Program Structure
  - [x] Define an exported function (`func Hellos()`)
  - [x] Create a new map (`make(map[string]string)`)
  - [x] using the range function to return the index and value from some `slice` (`range names`)
  - [x] Use for loop (`for _, name := range names {}`)
  - [x] Associate some value in a map (`messages[name] = message`)
  - [x] Get the slice length (`len(format)`)
  - [x] Create a slice of strings (`[]string{"Gladys", "Samantha", "Darrin"}`)
- [x] Standard packages
  - [x] `math/rand`
  - [x] `log`
- [x] Learning things
  - [x] `slice` is sorted by index
  - [x] `map` is sorted by keys (key strings will be considered in alphabetic order)
