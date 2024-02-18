package main

import "fmt"

func main() {
	fmt.Println(ValidaBin([]int{1, 2, 3, 4, 5}, 1))
	fmt.Println(ValidaBin([]int{1, 2, 3, 4, 5}, -1))
}

func ValidaBin(list []int, i int) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := (low + high) / 2
		if list[mid] == i {
			return mid
		}
		if list[mid] < i {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
