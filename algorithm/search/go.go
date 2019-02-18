package main

import (
	"fmt"
)

/*

	二分查找:


*/

func search(arr []int, value int) (result int) {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == value {
			return arr[mid]
			// 如果中间的值大于当前value
		} else if arr[mid] > value {
			right = mid - 1

		} else if arr[mid] < value {
			left = mid + 1
		}

	}
	return -1
}

func main() {

	result := search([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 6)
	fmt.Println(result)

}
