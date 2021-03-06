package tools

import (
	"sort"
)

/**
交换
*/
func Swap(arr []int, i, j int) {
	arr[i] = arr[i] ^ arr[j]
	arr[j] = arr[i] ^ arr[j]
	arr[i] = arr[i] ^ arr[j]
}

/**
 * 校验器
 */
func Validator(arr1, arr2 []int) bool {
	if arr1 == nil && arr2 != nil || arr1 != nil && arr2 == nil {
		return false
	}
	if arr1 == nil && arr2 == nil {
		return true
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

/**
 * 比较函数，通过go sdk里面的排序方法进行排序
 */
func CompareSort(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
}

/**
 * 获取随机数组
 * @param maxNum  数组中数的最大值
 * @param maxSize 数组最大长度
 */
func GetRandomArray(maxNum int, maxSize int) []int {
	//比如maxSize = 100
	//tools.GetRandom(int64(maxSize-1)) 的范围是 [0 99]
	//tools.GetRandom(int64(maxSize-1)) + 1 的范围是 [1 100] 也就是说数组长度至少是1，最大是maxSize
	size := GetRandom(maxSize-1) + 1
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = GetRandom(maxNum) - GetRandom(maxNum)
	}
	return arr
}

/**
 * 随机样本产生器，能随机生成一个有序数组，和一个可能在此范围内的数
 */
func GenerateRandomSortedArrayAndNum(maxSize int, step int) ([]int, int) {
	size := GetRandom(maxSize)
	arr := make([]int, size)
	tmp := 0
	for i := 0; i < size; i++ {
		arr[i] = tmp + GetRandom(step)
		tmp = arr[i]
	}
	//最后生成一个随机数
	if tmp == 0 {
		tmp = 1
	}
	num := GetRandom(tmp << 1)
	return arr, num
}
