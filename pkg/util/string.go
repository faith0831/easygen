package util

import "strings"

// ToSnake 转换成蛇形命名
func ToSnake(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			if i == 0 {
				result += strings.ToLower(string(s[i]))
			} else {
				result += "_" + strings.ToLower(string(s[i]))
			}
		} else {
			result += string(s[i])
		}
	}

	return result
}
