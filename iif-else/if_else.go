package main

import "fmt"

func main() {

	//taking input from the user  

	var age  int 
	var name string

	fmt.Print("Enter your age :")
	fmt.Scan(&age)

	fmt.Println("Enter Your name :")
	fmt.Scan(&name)

	if age > 18 {
		fmt.Println("Hii ", name , "Congrats ! You are eligible to go in")
		}else if age  >12{
			fmt.Println("Hey kid - " , name , "You should go home")
		}else{
			fmt.Println("Hii", name, "Baby , You are not allowed to go in")
		}


		// testing like middleware for RBAC

		var role int

		fmt.Println("Please Enter your role is 1, 2 and 3")
		fmt.Scan(&role)

		


		if role == 1 {
			fmt.Println ("welcome admin")
		}else if role ==2{
			fmt.Println("Hii user , Enjoy the app")

		}else{
			fmt.Println("Hii Guest , Please sign up to enjoy the app")
		}




}