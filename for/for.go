package main

import "fmt"

// there is only one loop in go and that is for loop, no while loop or do while loop
func main(){


	//checking while loop using for 

	i:=1

	for i<=5{
		fmt.Println(i)
		i++
		
	}

	//infinite loop in for 

	// for {
	// 	fmt.Println("Pawan Kumar")
	// }


	// classic for loop , worked as same which we learnn in C++

	for i :=0; i<=10 ;i++{

		if i ==1{

			continue
		}


		if i == 3 {

			break
			
		}
			fmt.Println(i)
	
	}

	//range check  

	for i:= range 3{
		fmt.Println(i)
	}


}