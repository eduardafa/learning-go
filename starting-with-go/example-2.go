package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "John"
	var age int = 24
	var version float32 = 1.1
	fmt.Println("Hello, Mr.", name, "and your age is", age)
	fmt.Println("This program is running on version", version)

	fmt.Println("The type of the variable 'name' is", reflect.TypeOf(name))
}
