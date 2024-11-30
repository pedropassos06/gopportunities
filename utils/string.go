package utils

import (
	"strconv"
)

func StringToUint(s string) (uint, error) {
	uint64, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(uint64), nil
}
