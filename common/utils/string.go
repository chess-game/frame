package utils

import (
	"fmt"
	"strconv"
	"strings"
)

//  其他类型转化为字符串
func Tostring(ceil interface{}) string {
	return fmt.Sprintf("%v", ceil)
}

// 字符串转换 int   没有位数
func Int(ceil string) int {
	i, err := strconv.Atoi(ceil)
	if err != nil {
		return 0
	}
	return i
}

// 字符串转int32
func Int8(ceil string) int8 {
	ib, err := strconv.ParseInt(ceil, 10, 8)
	if err == nil {
		return int8(ib)
	}
	return int8(0)
}

// 字符串转int32
func Int32(ceil string) int32 {
	ib, err := strconv.ParseInt(ceil, 10, 32)
	if err == nil {
		return int32(ib)
	}
	return int32(-1)
}

// 字符串转int64
func Int64(ceil string) int64 {
	ib, err := strconv.ParseInt(ceil, 10, 64)
	if err == nil {
		return ib
	}
	return 0
}

// 字符串转int32
func Uint8(ceil string) uint8 {
	ib, err := strconv.ParseUint(ceil, 10, 8)
	if err == nil {
		return uint8(ib)
	}
	return uint8(0)
}

// 字符串转int32
func Uint32(ceil string) uint32 {
	ib, err := strconv.ParseUint(ceil, 10, 32)
	if err == nil {
		return uint32(ib)
	}
	return uint32(0)
}

// 字符串转int64
func Uint64(ceil string) uint64 {
	ib, err := strconv.ParseUint(ceil, 10, 64)
	if err == nil {
		return ib
	}
	return 0
}

// 字符串转float32
func Float32(ceil string) float32 {
	fb, err := strconv.ParseFloat(ceil, 32)
	if err == nil {
		return float32(fb)
	}
	return float32(0)
}

// 字符串转float64
func Float64(ceil string) float64 {
	fb, err := strconv.ParseFloat(ceil, 64)
	if err == nil {
		return fb
	}
	return float64(0)
}

// 字符串转bool
func Bool(ceil string) bool {
	ceil = strings.ToLower(ceil)
	if ceil == "true" || ceil == "yes" || ceil == "on" || Int32(ceil) > 0 {
		return true
	}
	return false
}

/*
	字符串截取    模仿PHP的substr
	@param begin int 起始位置
	@param end  int  结束位置    0表示截取到末端    整数表示长度    负数表示从后面开始截取
*/
func SubStr(str string, begin, length int) (substr string) {
	// 将字符串的转换成[]rune
	rs := []rune(str)
	lth := len(rs)

	// 简单的越界判断
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	var end int
	if length == 0 {
		end = lth
	} else if length < 0 {
		// 从后面截取字符串
		end = lth + length
	} else {
		end = begin + length
	}
	if end > lth {
		end = lth
	}
	if begin >= end {
		return ""
	}
	// 返回子串
	return string(rs[begin:end])
}
