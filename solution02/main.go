package main

import "fmt"

func main() {
	massive := []int{2, 8, 5, 1, 13, 20, 1, 3, 12, 0}
	f := func(a []int) []int {
		for i := 0; i < len(a); i++ {
			for j := len(a) - 2; j >= 0; j-- {
				if a[j+1] > a[j] {
					a[j+1], a[j] = a[j], a[j+1]
				}
			}
		}
		return a
	}
	fmt.Println("Unsorted massive:", massive)
	fmt.Println("Sorted and inversed massive:", f(massive))
}
