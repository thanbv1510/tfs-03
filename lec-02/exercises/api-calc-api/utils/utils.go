package utils

import (
	"errors"
	"strconv"
)

func GetParamValue(paramKey string, mapPrams map[string][]string) (int, error) {
	if len(mapPrams) != 2 {
		return 0, errors.New("Param not valid")
	}

	nums, ok := mapPrams[paramKey]
	if !ok || len(nums) != 1 {
		return 0, errors.New("Param not valid")
	}

	return strconv.Atoi(nums[0])
}
