package main

import (
	"fmt"
	"time"
)

func SwitchProgram(){

	switch time.Now().Weekday(){
	case time.Sunday:
		fmt.Println("Store is closed today")
	case time.Saturday:
		fmt.Println("Today is weekend, so Store will only open for 4 hours")


	default:
		fmt.Println("Store is live and serving from 9am to 9pm")

	}
}