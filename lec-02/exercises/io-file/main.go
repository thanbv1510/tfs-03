package main

import (
	"fmt"
	"iofile/helpers"
	"iofile/utils"
)

func main() {
	path := "data.txt"
	data := helpers.ReadFile(path)

	minNum := utils.FindMinNumber(data)
	maxNum := utils.FindMaxNumber(data)
	averageNum := utils.GetAverage(data)

	isExist := utils.IsExist(10, data)

	primeNums := make([]int, 0)
	for _, v := range data {
		if utils.IsPrimeNumber(v) {
			primeNums = append(primeNums, v)
		}
	}

	fmt.Printf("Min: %d, Max = %d, Average = %d, %d exist? %t, List prime numbers = %v",
		minNum, maxNum, averageNum, 10, isExist, primeNums)
}
