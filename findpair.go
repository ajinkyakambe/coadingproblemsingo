package main

import "fmt"

/*


Find a pair with the given sum in an array

*/

func main() {

	inputArr := []int{2, 5, 7, 8, 10, 15, 26}
	targetSum := 25
	findingThePair(inputArr, targetSum)

}

// Brute force approach

func findingThePair(input []int, sum int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if input[i]+input[j] == sum && i != j {
				fmt.Printf("indexs are %d - %d and values are %d and %d \n", i, j, input[i], input[j])
			}
		}
	}
}
