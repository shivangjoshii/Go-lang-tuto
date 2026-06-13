package main

import "fmt"

func main() {

	var age int
	var name string

	fmt.Println("== inititialised the vars==")
	fmt.Println("== Checking if else conditions==")


	fmt.Println("Enter You name :")
	fmt.Scan(&name)
	fmt.Println("Enter Your age :")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println(name,", You are an adult")

	} else if age >12 && age < 18 {
		fmt.Println(name,", You are a teenager")
	}


}