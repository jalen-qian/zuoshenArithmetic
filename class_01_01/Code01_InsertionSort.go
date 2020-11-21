package main

import (
	"fmt"
	"github.com/zuoshenArithmetic/tools"
)

/**
插入排序
[0         N-1] n个数，0 - N-1的索引位置
1. 先做到0-0位置有序（废话，肯定有序，只有一个数）
2. 再做到0-1位置有序（如果arr[1] < arr[0]则交换）
3. 再做到0-2位置有序（2位置看前一个，如果比前一个小，则交换，否则停止）
4. 做到  0-i位置有序

时间复杂度：O(n^2)
*/

func InsertionSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := 0; i < n; i++ { //0 ~ i位置做到有序
		for j := i - 1; j >= 0 && arr[j] > arr[j+1]; j-- {
			tools.Swap(arr, j, j+1)
		}
	}
}

func main() {
	//写对数器
	//arr := []int{2, 9, 1, 5, 3, 0, -2, 3}
	//InsertionSort(arr)
	//fmt.Println(arr)

	testTimes := 100000
	maxNum, maxSize := 1000, 1000
	fmt.Println("开始测试")
	for i := 0; i < testTimes; i++ {
		arr1 := tools.GetRandomArray(maxNum, maxSize)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		//对arr1调用插入排序
		InsertionSort(arr1)
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
