package main

import "fmt"

func main() {
	var name string
	var age int
	fmt.Print("Enter Your name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter Your Age: ")
	fmt.Scanln(&age)

	fmt.Println("Hello", name, "Your age is", age)
	hello()
}
