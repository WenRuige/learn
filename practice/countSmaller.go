package main

import "fmt"

func countSmaller(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				count++

			}
		}
		result = append(result, count)

	}

	return result

}

func main() {
	result := countSmaller([]int{5, 2, 6, 1})
	fmt.Println(result)
}
