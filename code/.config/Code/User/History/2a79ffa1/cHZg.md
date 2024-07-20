# Golang

## What is Go?

- Go is an open-source programming language developed by Google
- Focus on simplicity, clarity and scalability
- High performance and focus on concurrency (running multiple tasks simultaneously)
- Batteries included, meaning that many core features are built-in
- Static typing

## Poupular use-cases

- Networking and APIs
- Microservices
- ClI tools

## Packages

- Used to organize your code and to export and import features across our files
- Every different package must be placed in a separate folder
- At least one package is required in a Go project
- `package`: a collection of golang files that are in the same directory
- `package main`: a reserved package name that tells golang where the execution should start
- Within a main package, there must be only one `main` function definition (the entry point of the Go project execution)
- Install third-party package: `go get <github-repo>`

### Exporting Functions/Variables

- To export a function/variable in Go to make it available in other packages, you need to name it with its first character in upper-case

## Modules

- `module`: a set of packages (the project itself)
- Module creation: `go mod init <path where the module can be found>`
- Example: `go mod init example.com/first-app`

## Execution

- `go run <gofile>`
- `go run .` for executing a module (a module is required)

## Compilation

- `go build` (a module is required)
- The command above will generate an executable that you can run with `./filename`

## Libraries

- `import "fmt"`: the format I/O goland libray
- `import "os"`: for interacting with the OS (write/read etc)

## Functions/Methods

- `main()`: the entry point function for golang
- `fmt.Printf(any)`: print "any" formatted to stdout: `fmt.Printf("The value is: %v\n", value)`
- `fmt.Print(any)`: print "any" to stdout
- `os.WriteFile(...)`
- `os.ReadFile(...)`

## Conversion

- `fmt.Sprint(any)`: convert "any" to string and **return it as a new string**
- `[]byte(string)`: convert "string" to a slice of bytes
- `string([]byte)`: convert a slice of bytes to a string
- `strconv.parseFloat(string, bitSize)`: convert a string to a float

## Install Third-Party Packages

- Find go packages here: https://pkg.go.dev/
- Install go package using: `go get <package-url>`
- The package will be installed globally and available in the `go.mod` file
- It is possibile to install all the packages listed in the `go.mod` file by using the `go get` command

## Pointers

- Pointers are variables that stores value addresses instead of values
- Why Pointers?
  - Avoid unnecessary value copies
  - Directly mutate values
- Get a pointer using `&<variable>`
- Define a type for a pointer using `*<type>`
- De-reference a pointer using `*<pointer>`
- Writing into a variable using a pointer is as easy as `*<pointer> = value`

## Rune

- A rune is an alias for the `int32` type and represents a single Unicode code point.
- Single quotes are used to define runes in Go

## Structs

- Data type that allows to structure data, to group related data together into a collection of data
- It is possible to attach methods to a struct by using a _Receiver_ (a type declared right after `func`)

  ```go
  type User struct {
  	firstName string
  	lastName  string
  	birthdate string
  	createdAt time.Time
  }

  func (u User) OutputUserDetails() {
  	fmt.Println(u.firstName, u.lastName, u.birthdate)
  }
  ```

- If you have a method that wants to **edit** a structs, you must pass a **pointer** to that struct to the method:
  ```go
  func (u *User) clearUserName() {
  	u.firstName = ""
  	u.lastName = ""
  }
  ```
- `tags`: metadata that you can add to your struct fields

  ```go
  // Example with the "encoding/json" package

  // The fields are renamed for the creations of the json
  type Note struct {
  	Title     string    `json:"title"`
  	Content   string    `json:"content"`
  	CreatedAt time.Time `json:"created_at"`
  }

  json, err := json.Marshal(note)
  ```

## Interfaces and Generics

### Interfaces

```go
// This interface includes objects that implements the Save method with the exact signature
type saver interface {
	Save() error
}

// This interface includes objects that implements the Display method with the exact signature
// and the saver interface
type outputtable interface {
	saver
	Display()
}
```

### Generics

```go
// A generic function that takes two args of type T (int, float64 or string) and returns a + b
func add[T int | float64 | string](a, b T) T {
	return a + b
}
```

## Data Structures

### Arrays

```go
// Arrays in Go are not dynamic. So you need to declare the length first
prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// newPrices is a slice (not an array) that points to the address of the prices array (not a copy)
newPrices := prices[:3] // from index 0 to 3 (excluded)
fmt.Println(prices)
fmt.Println(len(newPrices)) // the length of newPrices => 3
fmt.Println(cap(newPrices)) // the capacity of newPrices => 4


newPrices = prices[1:] // from index 1 to the end
fmt.Println(len(newPrices)) // the length of newPrices => 3
fmt.Println(cap(newPrices)) // the capacity of newPrices => 3 *** the left part is lost ***
```

### Slices

```go

// Slices in Go are dynamic. So you don't need to declare the length
prices := []float64{10.99, 9.99, 45.99, 20.0}
fmt.Println(prices)

// the `append` function will create a **new slice** that we use to overrite prices
prices = append(prices, 5.99) // append takes a slice as a first argument and some values to append

```

### make

`make` is a function in Go that allows to creates and initializes **slices**, **maps** and **channels** by defining the type, size and capacity in advance.
The example below helps get a better understing of the capacity applied to slices:

```go

package main

import "fmt"

func main() {
	nPrices := make([]float64, 3, 4) // len 3, cap 4, [0,0,0]
	printSlice("nPrices", nPrices)
	nPrices = nPrices[:4]            // len 4, cap 4, [0,0,0,0]
	nPrices[3] = 100
	printSlice("nPrices", nPrices)   // len 4, cap 4, [0,0,0,100]

}

func printSlice(s string, x []float64) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

```

## Concurrency & Goroutines

In Go it's possible to run code in parallel and execute functions concurrently by using a feature called **goroutines**

### Channel

A channel in Go is simply a communication device, used to transmit some kind of data when working with goroutines.
A channel can be used to wait for the completion of a goroutine, or multiple goroutines.

```go
// create a channel
make(chan)
```
