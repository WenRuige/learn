package main

import "fmt"

// 快速排序是分治的算法
func quickSort(value []int) {

	quickSortInstance(value, 0, len(value)-1)
}

func quickSortInstance(value []int, left, right int) {
	// 最左元素当做基准元素

	if left < right {
		tmp := value[left]
		i, j := left, right
		for {
			//从右向左找比基准值大的数
			for i < j && value[j] >= tmp {
				j--
			}
			//从左向右找比基准小的数
			for i < j && value[i] <= tmp {
				i++
			}

			if i >= j {
				break
			}
			// 找到则交换
			value[i], value[j] = value[j], value[i]

		}

		value[left] = value[i]
		value[i] = tmp

		quickSortInstance(value, left, i-1)
		quickSortInstance(value, i+1, right)
	}
}

func main() {
	nums := []int{1, 2, 21, 3131231, 122, 2, 12, 221, 2, 2, 23}
	quickSort(nums)

	fmt.Println(nums)

}
