package main

import "fmt"

func tribonacci(n int) []int {
	// Handle base cases
	if n <= 0 {
		return []int{}
	} else if n == 1 {
		return []int{0}
	} else if n == 2 {
		return []int{0, 1}
	}

	// Initialize the sequence with the first three terms
	tribonacciSeq := []int{1, 1, 1}

	// Calculate the remaining terms
	for i := 3; i < n; i++ {
		nextNum := tribonacciSeq[i-1] + tribonacciSeq[i-2] + tribonacciSeq[i-3]
		tribonacciSeq = append(tribonacciSeq, nextNum)
	}

	return tribonacciSeq
}

// Example usage
func main() {
	result := tribonacci(20)
	fmt.Println(result)
	//Output of Code in terminal : [0, 1, 1, 2, 4, 7, 13, 24, 44, 81, 149, 274, 504, 927, 1705, 3136, 5768, 10609, 19513, 35890]
}

//Zachary Shaft Module 3 Project 3
