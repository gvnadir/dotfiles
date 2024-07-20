# Golang notes

## Structure

- `package`: a collection of golang files that are in the same directory
- `package main`: a reserved package name that tells golang where the execution should start
- `function main()`: the entry point function for golang

## Modules

- `module`: a set of packages (the project itself)
- module creation: `go mod init example.com/first-app`

## Execution

- `go run <gofile>`
- `go run .` for executing a module (a module is required)

## Compilation

- `go build` (a module is required)

## Libraries

- `import "fmt"`: the format I/O goland libray
- `import "os"`: for interacting with the OS (write/read etc)

## Functions

- `fmt.Print(something)`: print "something" to stdout
- `fmt.Sprint(something)`: convert "something" to string and return it (for example from Float64 to string)
- `string(something)`: convert "something" to string and return it (for example []byte to string)
