/*
<p>The sum of the squares of the first ten natural numbers is,</p>
$$1^2 + 2^2 + ... + 10^2 = 385.$$
<p>The square of the sum of the first ten natural numbers is,</p>
$$(1 + 2 + ... + 10)^2 = 55^2 = 3025.$$
<p>Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is $3025 - 385 = 2640$.</p>
<p>Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.</p>
*/

package main

import ("fmt")

func main() {

	var max_value int 
	var sum_of_squares, sum, square_of_sum int	
	var difference int


	fmt.Println("Enter max value: ")
	fmt.Scan(&max_value)
	
	
	fmt.Println(max_value)
	
	for index := 1 ; index <= max_value ; index ++  {
		fmt.Println(index)
		
		sum += index
		sum_of_squares += index*index
	
	}

	square_of_sum = sum * sum 
	difference = square_of_sum - sum_of_squares


	fmt.Println("the sqauare of the sum :")
	fmt.Println(square_of_sum)
	fmt.Println("the sum of squares:")
	fmt.Println(sum_of_squares)

	fmt.Println("the difference:")
	fmt.Println(difference)


}

