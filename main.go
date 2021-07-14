package main

import (
	"fmt"
	t "time"
	// creates a alias for time, also imports can be done inside (), no need for , or ;
)
// standard library for packages to import can be found at https://pkg.go.dev/std

func main () {
	// main function does not need to be called, it will run when the program is called
	fmt.Println("Let's see if this works, also know as Hello World!")
	fmt.Println(t.Now())
}