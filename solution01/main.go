package main

import "fmt"

const size = 10

func insertionSort(input [size]int) [size]int {
	for i := 1; i < size; i++ {
		for j := i; j > 0; j-- {
			if input[j-1] > input[j] {
				input[j-1], input[j] = input[j], input[j-1]
			}
		}
	}
	return input
}

func main() {
	massive := [size]int{2, 8, 5, 1, 13, 20, 1, 3, 12, 0}
	fmt.Println("Unsorted massive:", massive)
	fmt.Println("Sorted massive:", insertionSort(massive))
}
