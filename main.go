package main

func main() {
}
func checkValid(matrix [][]int) bool {
	arr := map[int]bool{}
	for i := 1; i <= len(matrix); i++ {
		arr[i] = true
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if !arr[j+1] {
				return false
			}
		}
	}
	for i := 0; i < len(matrix[0]); i++ {
		for j := 0; j < len(matrix); j++ {
			if !arr[j+1] {
				return false
			}
		}
	}
	return true
}
