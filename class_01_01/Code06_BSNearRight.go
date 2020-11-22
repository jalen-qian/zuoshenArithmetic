package main

import (
	"fmt"
	"github.com/zuoshenArithmetic/tools"
)

/**
 * 二分题3：
 * 给定一个有序的数组arr和一个数num，找到arr中 <= num 的最右侧的位置
 * 如果不存在，则返回-1
 *
 * 示例：
 * arr = [1 1 1 2 2 2 2 2 3 3 3 4 5 5] num = 4
 * 返回 11
 *
 * 思路：
 * 指定一个变量 index := -1
 * 先打到中点 mid , 判断arr[mid] <= num
 *     是：记录当前位置到index，要找的值可能在 mid 右边还有 <= num的值，继续对右边二分
 *     否：说明要找的值如果有，也只可能在 mid 左边，继续对左边二分
 * 二分结束：当二分越界时，结束二分，返回index
 */

func nearestRightIndex(sortedArr []int, num int) int {
	n := len(sortedArr)
	if n == 0 {
		return -1
	}
	index := -1
	L, R := 0, n-1
	mid := 0
	for L <= R {
		mid = L + ((R - L) >> 1)
		//如果当前mid满足 >= num，则记录mid并继续右边二分
		if sortedArr[mid] <= num {
			index = mid
			L = mid + 1
		} else {
			//否则继续左边二分
			R = mid - 1
		}
	}
	return index
}

// for test 易于实现，但时间复杂度比较高的实现 O(N)
func nearestRightIndexComparator(sortedArr []int, num int) int {
	for i := len(sortedArr) - 1; i >= 0; i-- {
		if sortedArr[i] <= num {
			return i
		}
	}
	return -1
}

func main() {
	//测试次数
	testTimes := 500000
	maxSize := 100
	step := 10
	success := true
	fmt.Println("开始测试...")
	for i := 0; i < testTimes; i++ {
		arr, num := tools.GenerateRandomSortedArrayAndNum(maxSize, step)
		index1 := nearestRightIndex(arr, num)
		index2 := nearestRightIndexComparator(arr, num)
		if index1 != index2 {
			success = false
			fmt.Println("测试失败：", arr, num, index1, index2)
			break
		}
	}
	if !success {
		fmt.Println("fucked!")
	} else {
		fmt.Println("succeed!")
	}
}
