package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type ValueDTO struct {
	Value string `json:"value"`
}

func calcHandler(w http.ResponseWriter, r *http.Request) {
	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if len(r.URL.String()) < 6 { // /calc/
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Param not valid"))
		return
	}
	params := r.URL.String()[6:]
	fmt.Println("Params: ", params)

	resultValue, err := getValue(params)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Param not valid"))
		return
	}

	valueJson, err := json.Marshal(ValueDTO{Value: fmt.Sprintf("%.2f", resultValue)})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(valueJson)
	if err != nil {
		panic(err.Error())
	}
}

func getValue(param string) (float64, error) {
	params := make([]string, 0)
	var word string
	for _, char := range param {
		switch {
		case (char >= '0' && char <= '9') || char == '*' || char == '/' || char == '.':
			word += string(char)
		case char == '-' || char == '+':
			params = append(params, word)
			params = append(params, string(char))
			word = ""
		}
	}
	if word != "" {
		params = append(params, word)
	}
	fmt.Println(params)

	for i, v := range params {
		if strings.ContainsAny(v, "* | /") {
			params2 := make([]string, 0)
			word1 := ""
			for _, char := range v {
				switch {
				case (char >= '0' && char <= '9') || char == '.':
					word1 += string(char)
				case char == '*' || char == '/':
					params2 = append(params2, word1)
					params2 = append(params2, string(char))
					word1 = ""
				}
			}
			if word1 != "" {
				params2 = append(params2, word1)
				word1 = ""
			}

			d, _ := calcData(params2)
			fmt.Println(d)
			params[i] = fmt.Sprintf("%v", d)
		}
	}
	fmt.Println(params)
	v, _ := calcData(params)
	fmt.Println(v)

	return v, nil
}

func calcData(input []string) (float64, error) {
	result, err := strconv.ParseFloat(input[0], 64)
	if err != nil {
		return 0, err
	}
	for i := 1; i < len(input)-1; i += 2 {
		value, err := strconv.ParseFloat(input[i+1], 64)
		if err != nil {
			return 0, err
		}
		operator := input[i]

		switch operator {
		case "+":
			result += value
		case "-":
			result -= value
		case "*":
			result *= value
		case "/":
			result /= value
		default:
			return 0, nil

		}
	}
	return result, nil
}

func main() {
	http.HandleFunc("/calc/", calcHandler)

	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic("Error when running server")
	}
}
