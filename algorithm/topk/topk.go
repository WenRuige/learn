package main

import (
	"fmt"
	"sort"
)

type num struct {
	num int
	fre int
}

type numsArr []*num

// less
func (num numsArr) Less(i, j int) bool {
	if num[i].fre < num[j].fre {
		return false
	}
	return true
}

// len
func (num numsArr) Len() int {
	return len(num)
}

// swap
func (num numsArr) Swap(i, j int) {
	num[i], num[j] = num[j], num[i]
}

// 求TOPK
func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := numMap[nums[i]]
		if ok {
			numMap[nums[i]] = numMap[nums[i]] + 1
		} else {
			numMap[nums[i]] = 1
		}
	}
	fw := make(numsArr, 0, len(numMap))
	for k, v := range numMap {
		fmt.Println(k, v)
		fw = append(fw, &num{k, v})
	}

	sort.Sort(fw)
	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, fw[i].num)
	}

	return res

}

func main() {
	nums := []int{1, 2, 3, 2, 4, 2, 3, 2, 2, 12, 2}
	res := topKFrequent(nums, 2)
	fmt.Println(res)
}
