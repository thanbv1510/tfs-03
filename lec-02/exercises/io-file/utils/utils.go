package utils

import "math"

func FindMaxNumber(arr []int) int {
	maxNum := arr[0]
	for _, v := range arr {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}

func FindMinNumber(arr []int) int {
	minNum := arr[0]
	for _, v := range arr {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}

func GetAverage(arr []int) int {
	total := 0
	for _, v := range arr {
		total += v
	}

	return total / len(arr)
}

func IsExist(num int, arr []int) bool {
	for _, v := range arr {
		if num == v {
			return true
		}
	}
	return false
}

func IsPrimeNumber(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
