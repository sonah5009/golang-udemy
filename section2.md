# Section 2: A Simple Start

## Q1. How do we run the code in our project?

go build: just compile, not execute
go run: compiles & executes
go fmt: formats all the code in each file in the current directory
go install: compiles and "installs" a package
go get: Downloads the raw source code of someone else's package
go test: Runs and tests associated with the current project

## Q2: What does 'pacakage main' mean?

Types of Packages

1. Executable - package main
   important: name of the package "package main". If not, it doesn't generate executable file.

2. Reusable - package claculator/package uploader...
   : kinda code dependancy, library, helper code

## Q3: What does 'import "fmt"' mean?

## Q4: What's that 'func' thing?

func is similar with ruby, python's
declare, set the name, for argument, code body
func, main, (), {}

## Q5: How is the main.go file organized?

package main: Package declaration
import "fmt": import other packages that we need
func main() {fmt.Println("hi there")}: Declare functions, tell Go to do things
