package util

import (
	"math/rand"
	"time"
)

func GetRandomNum(n int) int {
	rand.Seed(time.Now().UnixNano())

	if n <= 0 {
		return -1
	}
	return rand.Intn(n)
}
