package utils

import "fmt"

// 数字转序数词
func NumToStr(num int) string {
	switch num % 10 {
	case 1:
		return fmt.Sprintf("%dst", num)
	case 2:
		return fmt.Sprintf("%dnd", num)
	case 3:
		return fmt.Sprintf("%drd", num)
	default:
		return fmt.Sprintf("%dth", num)
	}
}