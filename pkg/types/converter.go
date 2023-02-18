package types

import (
	"goblog/pkg/logger"
	"strconv"
)

func UInt64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func StringToUInt64(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logger.LogError(err)
	}
	return i
}
