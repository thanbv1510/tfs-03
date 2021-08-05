package helpers

import (
	"bufio"
	"io/ioutil"
	"os"
	"strconv"
)

func ReadFile(path string) []int {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return make([]int, 0)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var result []int

	for scanner.Scan() {
		intValue, err := strconv.Atoi(scanner.Text())

		if err == nil {
			result = append(result, intValue)
		}
	}

	return result
}

func WriteFile(path, data string) error {
	byteData := []byte(data)
	return ioutil.WriteFile(path, byteData, 0777)
}
