package sort

type Sort interface {
	BubbleSort()
	MergeSort(int, int)
	QuickSort(int, int)
}

type ArrInt []int

func (arr *ArrInt) BubbleSort() {
	for i := 0; i < len(*arr)-1; i++ {
		for j := i + 1; j < len(*arr)-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

func (arr *ArrInt) QuickSort(left, right int) {
	right -= 1
	i, j := left, right
	pivot := (*arr)[(left+right)/2]
	for i <= j {
		for (*arr)[i] < pivot {
			i++
		}

		for (*arr)[j] > pivot {
			j--
		}

		if i <= j {
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			i++
			j--
		}
	}

	if left < j {
		arr.QuickSort(left, j)
		arr.QuickSort(j, right)
	}

	if i < right {
		arr.QuickSort(left, i)
		arr.QuickSort(i, right)
	}
}

func (arr *ArrInt) MergeSort(left int, right int) {
	mid := (left + right) / 2
	if left < right-1 {
		arr.MergeSort(left, mid)
		arr.MergeSort(mid, right)

		result := make([]int, 0)
		leftArr := (*arr)[left:mid]
		rightArr := (*arr)[mid:right]

		for len(leftArr) > 0 && len(rightArr) > 0 {
			if leftArr[0] <= rightArr[0] {
				result = append(result, leftArr[0])
				leftArr = leftArr[1:]
			} else {
				result = append(result, rightArr[0])
				rightArr = rightArr[1:]
			}
		}

		if len(leftArr) != 0 {
			result = append(result, leftArr...)
		}

		if len(rightArr) != 0 {
			result = append(result, rightArr...)
		}
		copy((*arr)[left:right], result)
	}
}
