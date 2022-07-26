package main

import (
	"fmt"
	"sort"
)

func numberOfPairs(nums []int) []int {
	arr := map[int]int{}
	for _, v := range nums {
		arr[v]++
	}
	res := 0
	for _, v := range arr {
		res += (v / 2)
	}
	return []int{res, len(nums) - 2*res}
}

func maximumSum(nums []int) int {
	sort.Ints(nums)
	arr := map[int][]int{}
	exist := false
	for k, v := range sumNum(nums) {
		c := arr[v]
		c = append(c, k)
		arr[v] = c
		if len(arr[v]) > 1 {
			exist = true
		}
	}
	if !exist {
		return -1
	}
	res := 0
	for _, v := range arr {
		if len(v) >= 2 {
			res = max(res, nums[v[len(v)-1]]+nums[v[len(v)-2]])
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func sumNum(nums []int) (arr []int) {
	// 将传来的数组进行转为数位的和, 顺序不改变
	for _, v := range nums {
		res := 0
		for v > 0 {
			res += v % 10
			v /= 10
		}
		arr = append(arr, res)
	}
	return arr
}
func main() {
	fmt.Println(maximumSum([]int{11, 3, 2, 1, 3, 2, 2}))

}
