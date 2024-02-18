package main

import "fmt"

func main() {
	fmt.Println(ValidaBin([]int{1, 2, 3, 4, 5, 20}, 20))
	fmt.Println(ValidaBin([]int{1, 2, 3, 4, 5, 7000}, 7000))
	fmt.Println(ValidaBin([]int{1, 2, 3, 4, 5, 7000}, 380000))
}

func ValidaBin(list []int, item int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		shot := list[mid]

		if shot == item {
			return mid
		}
		if shot < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
