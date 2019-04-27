package utils

import (
	"math/rand"
	"time"
)

/**
  生成一个随机数
*/
func RandInt(min int, max int) int {

	if min == max {
		return min
	}

	var nan = int(0)
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
		nan += int(rand.Int31())
		nan = (nan % (max - min)) + (min)
	}

	return nan
}
