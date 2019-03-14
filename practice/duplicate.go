package main

import "fmt"






/*

// {2,3,1,0,2,5,3}
1:{1,3,2,0,2,5,3}
2:{3,1,2,0,2,5,3}
3:{0,1,2,3,2,5,3}
总结:数组中重复的数字

解法:
* 排列法 On(logn)
* hash法: O(1)
* 交换数据法

exmaple:
tmp = 2
nums[0] = nums[2]
nums[2] = 2

*/
func duplicate(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return 1
			}
			tmp := nums[i]
			nums[i] = nums[tmp]
			nums[tmp]=  tmp
		}

	}
	return 0
}

func main() {
	result := duplicate([]int{2, 3, 1, 0, 2, 5, 3})
	fmt.Println(result)
}
