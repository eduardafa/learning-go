package main

import "fmt"

func main() {
	var name string = "John"
	var version float32 = 1.1
	fmt.Println("Hello, Mr.", name)
	fmt.Println("This program is running on version", version)

	fmt.Println("1 - Start monitoring")
	fmt.Println("2 - Show logs")
	fmt.Println("0 - Quit")

	var command int
	fmt.Scan(&command)
	fmt.Println("The address of the command variable is", &command)
	fmt.Println("The command chosen was", command)

	if command == 1 {
		fmt.Println("Monitoring...")
	} else if command == 2 {
		fmt.Println("Showing logs...")
	} else if command == 0 {
		fmt.Println("Quitting the program")
	} else {
		fmt.Println("Unknown command")
	}
}
