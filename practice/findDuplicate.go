package main

import "fmt"

func findDuplicate(nums []int) int {

	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i] == nums[j] && i != j {
				return nums[i]
				//break
			}
		}
	}
	return 0
}

func main() {
	res := findDuplicate([]int{1, 3, 4, 2, 2})
	fmt.Println(res)
}
