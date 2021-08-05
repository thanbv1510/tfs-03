package main

import (
	"fmt"
	"sort-data/sort"
)

func main() {
	arrInt := sort.ArrInt{11, 12, 9, 2, 12, 21, 10, 15, 13, 51, 11, 12, 34}
	fmt.Println("Before sort:", arrInt)

	arrInt.BubbleSort()
	arrInt.QuickSort(0, len(arrInt))
	arrInt.MergeSort(0, len(arrInt))

	fmt.Println("After sort:", arrInt)

}
