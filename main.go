package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			fmt.Println(n / i)
			return
		}
	}
}
