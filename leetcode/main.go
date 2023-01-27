package main

import (
	"fmt"
	"sort"
)

/**
* 代码描述: 公共前缀
* 作者:小大白兔
* 创建时间:2022/10/05 11:21:12
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	pre := strs[0]

	for i := 1; i < len(strs); i++ {
		pre = ll(pre, strs[i])
		if len(pre) == 0 {
			return ""
		}
	}

	return pre
}
func ll(str1, str2 string) string {
	length := min(len(str1), len(str2))

	index := 0

	for index < length && str1[index] == str2[index] {
		index++
	}

	return str1[:index]
}
func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

// func main() {
// 	dd := longestCommonPrefix([]string{"flower", "flow", "flight"})
// 	fmt.Println(dd)
// }

// JAVA_HOME=/usr/local/java/jdk-19
// JRE_HOME=/usr/local/java/jdk-19/jre
// CLASS_PATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar:$JRE_HOME/lib
// PATH=$PATH:$JAVA_HOME/bin:$JRE_HOME/bin
// export JAVA_HOME JRE_HOME CLASS_PATH PATH

/**
* 代码描述:三数之和
* 作者:小大白兔
* 创建时间:2023/01/02 20:41:39
 */
func threeSum(nums []int) (arr [][]int) {
	// 特判
	if len(nums) < 3 {
		return arr
	}

	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		n1 := nums[i]

		if nums[i] > 0 {
			return arr
		}

		// 去重
		for i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			n2, n3 := nums[l], nums[r]

			if n1+n2+n3 == 0 {
				arr = append(arr, []int{n1, n2, n3})
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[i] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return
}
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	dd := threeSum(nums)
	fmt.Println(dd)
}
