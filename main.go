package main

import "fmt"

func main() {
	var n int
	arr1, arr2 := map[int]int{}, map[int]int{}
	arr := []int{}
	a, b := 0, 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var c int
		fmt.Scan(&c)
		arr = append(arr, c)
		if (i+1)%2 == 1 {
			arr1[c]++
			a += c
		} else {
			b += c
			arr2[c]++
		}
	}
	res, res1 := 0, 0
	for _, v := range arr2 {
		res += v
	}
	for _, v := range arr1 {
		res1 += v
	}
	if a-b > 0 {
		fmt.Println(arr1[a-b])
	} else if b-a > 0 {
		fmt.Println(arr2[b-a])
	} else {
		res2, res3 := 0, 0
		for i := 0; i < n; i++ {
			if (i+1)%2 == 1 {
				arr2[arr[i]]--
			} else {
				arr1[arr[i]]--
			}
		}
		for _, v := range arr2 {
			res2 += v
		}
		res -= res2
		for _, v := range arr1 {
			res3 += v
		}
		res1 -= res3
		fmt.Println(min(res, res1))
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
