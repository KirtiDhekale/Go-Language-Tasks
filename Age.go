package main
import "fmt"


func Age(){
	var name string
	var age int
	fmt.Print("Enter you name: ")
	fmt.Scanln(&name)


	fmt.Print("Enter Your age")
	fmt.Scanln(&age)

	fmt.Println("Hello", name, "Your Age is", age)
}