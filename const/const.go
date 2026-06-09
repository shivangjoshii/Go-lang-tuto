package main

import "fmt"

//const can be also  assigned outside the main 

const age =19
var name = "Pawan"
func main() {

	//basically const is value that can not be changed once assigned

	name := "Pawan"

	const student = "Kumar"

	fmt.Println(name + student)


	//multiple constants

	const (
		port = 5000 
		server_text="Your server is running in port :"
	)

	fmt.Println(server_text, port)

}