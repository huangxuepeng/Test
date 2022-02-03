package main

import "fmt"

func swap(a string, b int) int {
	arr := []byte(a)
	res := 1
	for i := 0; i < len(arr); i++ {
		res = res*b + int(arr[i]-'0')
	}
	return res
}

func main() {
	var a, b string
	arr := map[int]bool{}
	fmt.Scan(&a, &b)
	aa := []byte(a)
	for i := 0; i < len(aa); i++ {
		if aa[i] == '0' {
			aa[i] = '1'
			arr[swap(string(aa), 2)] = true
			aa[i] = '0'
		} else {
			aa[i] = '0'
			arr[swap(string(aa), 2)] = true
			aa[i] = '1'
		}
	}
	bb := []byte(b)
	for i := 0; i < len(bb); i++ {
		dd := bb[i]
		for t := 0; t < 3; t++ {
			bb[i] = byte(t + '0')
			if arr[swap(string(bb), 3)] {
				fmt.Println(swap(string(bb), 3))
				return
			}
			bb[i] = dd
		}
	}
}
