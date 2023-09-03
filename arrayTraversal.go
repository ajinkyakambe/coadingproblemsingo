package main

import "fmt"

//common ways to traverse an array in Go

func main() {

	input := []int{2, 4, 5, 6, 7, 8, 9, 10, 11, 13, 12}

	//Single Pass (Left to Right or Right to Left):

	// Left to right
	for i := 0; i < len(input); i++ {
		fmt.Print(input[i], " ")
	}
	fmt.Println()
	// Right to left
	for i := len(input) - 1; i >= 0; i-- {
		fmt.Print(input[i], " ")
	}

	fmt.Println()
	fmt.Println()

	// Two pointers

	// Two pointers moving in the same direction
	left := 0
	right := len(input) - 1

	for left <= right {
		fmt.Print("left=", left, "right=", right, "\n")
		left++
		right--
	}

	fmt.Println()

	// Two pointers moving in opposite directions

	left = 0
	right = len(input) - 1
	for left < right {
		fmt.Print("left=", left, "right=", right, "\n")
		left++
		right--
	}

	fmt.Println()
	// Using range -- in other languages you can use recursive apporach but in go it is not idiomatic

	for index, value := range input {
		fmt.Print("Index and value is", index, value, "\n")
	}

}
