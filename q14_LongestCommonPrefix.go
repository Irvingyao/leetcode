package leecode

/*
q14

编写一个函数来查找字符串数组中的最长公共前缀。
如果不存在公共前缀，返回空字符串 ""。

eg1:
输入: ["flower","flow","flight"]
输出: "fl"

eg2:
输入: ["dog","racecar","car"]
输出: ""

说明:
所有输入只包含小写字母 a-z
*/

func LongestCommonPrefix(strs []string) string {
	// if input is empty, return empty
	if len(strs) == 0 {
		return ""
	}
	// if input only has one element, return the element
	if len(strs) == 1 {
		return strs[0]
	}

	i := 0
	for {
		var temp uint8
		for _, str := range strs {
			// if one element is "", return ""
			if len(str) == 0 {
				return ""
			}
			// this means the str is the common prefix
			if i >= len(str) {
				return str
			}
			// when the ith character is different, str[:i] is the common prefix
			if temp == 0 {
				temp = str[i]
			} else if temp != str[i] {
				return str[:i]
			}
		}
		i++
	}
}
