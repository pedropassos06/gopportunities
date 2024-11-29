package utils

import (
	"strconv"
)

func StringToUint(s string) (uint64, error) {
	return strconv.ParseUint(s, 10, 64)
}
