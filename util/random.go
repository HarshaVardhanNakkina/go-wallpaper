package util

import "math/rand"

func GetRandomNum(n int) int {
	if n <= 0 {
		return -1
	}
	return rand.Intn(n)
}
