package main

import (
	"fmt"
	"github.com/zuoshenArithmetic/tools"
)

/**
冒泡排序策略：
假定有个数组，长度为n，如下表示0~n-1的索引位置
[0          n-1]
1)在0~N-1范围内，依次比较相邻的两个数，如果左边的数比右边大，则交换（最终最大的数会到N-1位置）
2)在0~N-2范围内，依次比较相邻的两个数，如果左边的数比右边大，则交换（最终第二大的数会到N-2位置）
...
3)在0-1范围内做同样的操作

时间复杂度：O(n^2)
*/

func BubbleSort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				tools.Swap(arr, j, j+1)
			}
		}
	}
}

func main() {
	//写对数器
	//测试次数
	testTimes := 100000
	maxNum, maxSize := 1000, 1000
	fmt.Println("开始测试")
	for i := 0; i < testTimes; i++ {
		arr1 := tools.GetRandomArray(maxNum, maxSize)
		arr2 := make([]int, len(arr1))
		copy(arr2, arr1)
		//对arr1调用冒泡排序
		BubbleSort(arr1)
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
