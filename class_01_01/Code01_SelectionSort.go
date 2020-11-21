package main

import (
	"fmt"
	"github.com/zuoshenArithmetic/tools"
)

/**
选择排序：
[                  ]
 0               N-1
选择排序策略：
1)从0~N-1位置中，找到最小的那个数，然后和0位置的数交换
2)从1~N-1位置中，找到最小的那个数，然后和1位置的数交换
...
3)从N-2~N-1位置中，找到最小的数，然后和N-1位置的数交换


时间复杂度：O(n^2)
*/
func SelectionSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	tmp := 0
	for i := 0; i < n; i++ {
		tmp = i
		//找最小的数的位置
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[tmp] {
				tmp = j
			}
		}
		//如果找到了最小数的位置，且这个位置不是当前的i位置，则交换
		if i != tmp {
			tools.Swap(arr, i, tmp)
		}
	}
}

func main() {
	//写对数器
	//测试次数
	testTimes := 10000
	maxNum, maxSize := 1000, 1000
	fmt.Println("开始测试")
	for i := 0; i < testTimes; i++ {
		arr1 := tools.GetRandomArray(maxNum, maxSize)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		//对arr1调用选择排序
		SelectionSort(arr1)
		//对arr2调用系统的排序
		tools.CompareSort(arr2)
		if !tools.Validator(arr1, arr2) {
			fmt.Println("出错了！")
			fmt.Println(arr1)
			fmt.Println(arr2)
			break
		}
	}
	fmt.Println("测试通过！")
}
