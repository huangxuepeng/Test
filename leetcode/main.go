package main

import "fmt"

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

func main() {
	dd := longestCommonPrefix([]string{"flower", "flow", "flight"})
	fmt.Println(dd)
}
JAVA_HOME=/usr/local/java/jdk-19       
JRE_HOME=/usr/local/java/jdk-19/jre     
CLASS_PATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar:$JRE_HOME/lib
PATH=$PATH:$JAVA_HOME/bin:$JRE_HOME/bin
export JAVA_HOME JRE_HOME CLASS_PATH PATH
