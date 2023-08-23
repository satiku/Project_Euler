/*
<p>By listing the first six prime numbers: $2, 3, 5, 7, 11$, and $13$, we can see that the $6$th prime is $13$.</p>
<p>What is the $10\,001$st prime number?</p>
*/

package main

import (
	"fmt"
)

func isPrime(number int) bool {

	var is_prime bool = true

	for check := number / 2; check >= 2; check-- {

		if number%check == 0 {

			is_prime = false
			break

		}

	}

	return is_prime
}

func main() {

	var prime_number int
	var last_prime int = 1

	fmt.Println("please enter nubmber:")
	fmt.Scan(&prime_number)

	for index := 1; index <= prime_number; index++ {
		var prime_check bool = false
		fmt.Print(index, " ")
		for number := last_prime + 1; prime_check == false; number++ {

			if isPrime(number) {
				last_prime = number
				fmt.Println(number)
				prime_check = true
				break
			}
		}
	}

}
