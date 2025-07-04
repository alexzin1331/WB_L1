package main

import "fmt"

func BSCycle(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func BSRecursion(arr []int, target, left, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return BSRecursion(arr, target, mid+1, right)
	} else {
		return BSRecursion(arr, target, left, mid-1)
	}
}

func main() {
	m := []int{-13, -5, -2, -1, 2, 3, 3, 4, 4, 5, 5, 6, 6, 9, 11, 13, 56}
	fmt.Println("Recursion (element found): ", BSRecursion(m, 3, 0, len(m)-1))
	fmt.Println("Recursion (element doesn't found):", BSRecursion(m, 1012, 0, len(m)-1)+1)
	fmt.Println("Cycle (element found): ", BSCycle(m, 3))
	fmt.Println("Cycle (element doesn't found):", BSCycle(m, 1012))
}
