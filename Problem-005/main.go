/*
<p>$2520$ is the smallest number that can be divided by each of the numbers from $1$ to $10$ without any remainder.</p>
<p>What is the smallest positive number that is <dfn class="tooltip">evenly divisible<span class="tooltiptext">divisible with no remainder</span></dfn> by all of the numbers from $1$ to $20$?</p>
*/

package main

import ("fmt")

func main() {

	var range_max, answer int 
	var number int = 1	
	var check bool = true

	fmt.Println("Enter max range: ")
	fmt.Scan(&range_max)
	
	
	fmt.Println(range_max)

	
	for check {
		for index := 1 ; index <= range_max ; index ++  {
			if number % index == 0 && index == range_max {
				answer = number 
				check = false	
				break
			} else if number % index != 0 {
				break
			}
		}
		number ++

	}
	fmt.Println("answer is :")	
	fmt.Println(answer)
}

