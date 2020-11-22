package main

/**
 * 二分题四：局部最小值
 * 给定一个无序的数组arr，这个arr中任意两个相邻的数都不相等，找到其中任意一个局部最小值的位置。如果数组为空，返回-1
 *
 * 局部最小值的定义：
 * 假定数组arr的长度为n [0    n-1]
 * 1.只有一个数：如果数组长度是1，那么这个数就是局部最小值。下面是数组长度大于1的情况。
 * 2.左边界：如果数组 arr[0] < arr[1]，那么说 0 位置是局部最小值
 * 3.右边界：如果数组 arr[n-2] > arr[n-1]，那么说 n-1 位置是局部最小值
 * 4.中间某处： 如果位置i满足 arr[i-1] > arr[i] && arr[i+1] > arr[i]，那么i位置是局部最小值
 *
 * 思路：不是一定要有序才能二分，只要找到一个规律，能够排除掉左边的所有数，或者右边的所有数，只在剩下的数里继续操作，
 *      就能二分
 * 1.先判断左右边界，如果左右边界中有1个是局部最小值，那么直接返回
 * 2.如果没有，说明arr[0] > arr[1] && arr[n-2] < arr[n-1]
 *   也就是说，arr数组开头是变小的趋势，结尾是变大的趋势
 *   这说明arr数组中间，一定至少有一个局部最小值，不然不可能出现开头变小，结尾变大
 *
 *   那么我们打到arr的中点mid，判断mid是否是局部最小值
 *       是：   直接返回
 *       不是： 说明要么arr[mid-1] < arr[mid] 要么 arr[mid+1] < arr[mid]（相邻不等）
 *                 如果是 arr[mid-1] < arr[mid]， 那么在 [0 ~ mid]区间内，又形成了开头变小结尾变大
 *                       那么可以排除掉 mid 右边的值，继续对mid左边的数二分
 *                 如果是 arr[mid+1] < arr[mid]， 同理，继续在mid右边的数二分。
 *
 */

func getLessIndex(arr []int) int {
	//如果数组长度为0，没有局部最小值，返回-1
	n := len(arr)
	if n == 0 {
		return -1
	}
	//判断边界值或者只有1个数的情况
	if n == 1 || arr[0] < arr[1] {
		return 0
	}
	if arr[n-2] > arr[n-1] {
		return n - 1
	}
	L, R, mid := 1, n-2, 0
	// [8 6 5 7]
	for L < R {
		mid = L + ((R - L) >> 1)
		//判断中点是否就是局部最小值
		if arr[mid-1] > arr[mid] && arr[mid] < arr[mid+1] {
			return mid
		}
		if arr[mid-1] < arr[mid] {
			R = mid - 1
		} else { // arr[mid] > arr[mid + 1]
			L = mid + 1
		}
	}
	//L == R时，跳出了循环，此时还没有在for循环中Return，说明L的位置一定是局部最小值
	//因为一定存在局部最小值，for循环二分没找到，说明这个局部最小值只能是L位置
	return L
}

func main() {
	//写
}
