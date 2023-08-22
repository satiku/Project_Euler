/*
<p>The prime factors of $13195$ are $5, 7, 13$ and $29$.</p>
<p>What is the largest prime factor of the number $600851475143$?</p>
*/


package main

import ("fmt")

func isPrime(number int) bool{
	
	var is_prime bool = true


	for check := number / 2 ; check >= 2 ; check-- {

		if number % check == 0 {
			
			is_prime = false
			break

		}

	}

	return is_prime
}

func factorsArray(number int) []int{

	factors := []int{}
//	var factored bool = false 


	
	for check := number / 2 ; check >= 2 ; check-- {
		if number % check == 0 {
			factors = append(factors, check)	
			factors = append(factors, number/check)	
			check = number / check	
		}

	}

	for _ , factor := range factors{
		if isPrime(factor) == false {
			factorsArray(factor)
			continue
		}	

		fmt.Println(factor)
	} 

	return factors
}

func main(){

	var number int 


	fmt.Println("please enter nubmber:")
	fmt.Scan(&number)

	fmt.Println(factorsArray(number))


}
