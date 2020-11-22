package main

import (
	"fmt"
	"github.com/zuoshenArithmetic/tools"
)

/**
 * 二分题目1：
 *
 * 给定一个有序数组，指定一个数，判断这个数在数组中是否存在
 * 思路：
 * 先打到中点 mid ，判断此位置的数是否 == num
 *     是：返回true
 *     否：判断是否 > num
 *         是：在mid左边继续找中点二分
 *         否：在mid右边继续找中点二分
 * 如果没有数了，还没有找到num，则返回false
 */
func exist(arr []int, num int) bool {
	if len(arr) == 0 {
		return false
	}
	L, R := 0, len(arr)-1
	for L < R {
		//先找到中点
		mid := L + ((R - L) >> 1)
		//先判断中点位置是否就是num
		if arr[mid] == num {
			return true
		}
		//否则判断中点位置的数是小了还是大了
		if arr[mid] > num {
			//如果中点位置的数 > num，说明num只可能在L~mid的区间，所以把mid赋值给R
			R = mid - 1
		} else {
			//如果中点位置的数 < num，说明num只可能在mid ~ R的区间，所以把mid赋值给L
			L = mid + 1
		}
	}
	//当最后只剩一个数 (L >= R) 时，会跳出循环，但是这个数还没有验证
	//所以最后要验证一下，这个下标一定在L的位置，因为L一定>0，而R可能为 - 1
	//还有，如果数组本身长度就是1，则不会进入循环，此时也要额外判断一次
	return arr[L] == num
}

//for test 比较函数
func existComparator(arr []int, num int) bool {
	if len(arr) == 0 {
		return false
	}
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

func main() {
	//对数器
	//测试次数
	testTimes := 500000
	maxSize := 100
	step := 100
	success := true
	fmt.Println("开始测试...")
	for i := 0; i < testTimes; i++ {
		arr, num := tools.GenerateRandomSortedArrayAndNum(maxSize, step)
		arr1 := make([]int, len(arr))
		copy(arr1, arr)
		res1 := exist(arr, num)
		res2 := existComparator(arr1, num)
		if res1 != res2 {
			success = false
			fmt.Println("测试失败：", arr, num)
			break
		}
	}
	if !success {
		fmt.Println("fucked!")
	} else {
		fmt.Println("succeed!")
	}
}
