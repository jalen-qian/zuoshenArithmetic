package main

import (
	"fmt"
	"github.com/zuoshenArithmetic/tools"
)

/**
 * 二分题2：给定一个有序的数组arr,和一个整数num,找到arr中 >= num的最左侧的位置
 * 如果不存在这个位置，返回-1
 * 策略：
 * 先打到中点，判断 arr[mid] 是否 >=num
 *    是：记录下这个位置，但是mid左边还有可能存在 >=num 的数，继续左边二分
 *    否：说明mid右边才可能有 >=num 的数，继续右边二分
 * 结束：当二分越界（二分区间左边 > 二分区间右边）时结束
 * 如果一直没有找到 >= num的数，二分就结束了，那么返回-1
 *
 * 时间复杂度分析：O(logN)
 */
func nearestIndex(sortedArr []int, num int) int {
	index := -1
	L, R := 0, len(sortedArr)-1
	mid := 0
	for L <= R {
		mid = L + ((R - L) >> 1)
		if sortedArr[mid] >= num {
			index = mid
			R = mid - 1
		} else {
			L = mid + 1
		}
	}
	return index
}

// for test 易于实现，但时间复杂度比较高的实现 O(N)
func nearestIndexComparator(sortedArr []int, num int) int {
	for i, v := range sortedArr {
		if v >= num {
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
		index1 := nearestIndex(arr, num)
		index2 := nearestIndexComparator(arr, num)
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
