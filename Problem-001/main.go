package main

import ("fmt")

func main() {

	var range_max int 
	var sum int	

	fmt.Println("Enter max range: ")
	fmt.Scan(&range_max)
	
	
	fmt.Println(range_max)
	
	for index := 1 ; index < range_max ; index ++  {
		if index % 3 == 0 || index % 5 == 0 {
			sum += index	
			fmt.Println(index)
		}
	}
	println(sum)
}

