package main

import "fmt"

func main() {
	fmt.Println(isValid("jjj"))
}
func isValid(s string) bool {
	// arr := []rune{}
	a := []byte(s)
	for i := 0; i < len(a); i++ {
		fmt.Println(s[i])
	}
	return true
}
