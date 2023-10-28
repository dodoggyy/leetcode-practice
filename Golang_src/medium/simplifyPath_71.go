package medium

import "strings"

// Use stack:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 4.8 MB, less than 15.60%
func simplifyPath(path string) string {
	arr := []string{}
	str := ""
	for i := range path {
		switch path[i] {
		case '/':
			if len(str) > 0 {
				arr = append(arr, str)
				str = ""
			}
		default:
			str = str + string(path[i])
		}
	}
	if len(str) > 0 {
		arr = append(arr, str)
	}

	stack := []string{}
	for i := range arr {
		switch arr[i] {
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		case ".":
		default:
			stack = append(stack, arr[i])
		}
	}

	if len(stack) == 0 {
		return "/"
	}

	result := "/"
	for i := range stack {
		result = result + stack[i] + "/"
	}

	return result[:len(result)-1]
}

// Use stack 2:
// Time Complexity: O(n)
// Space Complexity:O(n)
// Runtime: 0 ms, faster than 100.00%
// Memory Usage: 3.15 MB, less than 35.84%
func simplifyPath2(path string) string {
	stack := []string{}
	pathArr := strings.Split(path, "/")
	for _, str := range pathArr {
		switch str {
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			if len(str) > 0 && str != "." {
				stack = append(stack, str)
			}
		}
	}

	return "/" + strings.Join(stack, "/")
}
