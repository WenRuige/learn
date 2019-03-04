package main

import "fmt"

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	result := []int{}
	for i := 0; i < len(nums)-k+1; i++ {
		max := -1 << 32
		for j := 0; j < k; j++ {
			if max <= nums[i+j] {
				max = nums[i+j]
			}
		}
		result = append(result, max)

	}
	return result

}

func main() {
	result := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Println(result)
}
