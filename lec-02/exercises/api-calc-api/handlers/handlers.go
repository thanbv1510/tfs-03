package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

type ValueDTO struct {
	Value int `json:"value"`
}

func CalcHandler(num1, num2 int, w http.ResponseWriter, r *http.Request) {
	rawURLs := strings.Split(r.URL.String(), "/")
	if len(rawURLs) != 3 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	resultValue, err := getValue(rawURLs[2], num1, num2)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	valueJson, err := json.Marshal(ValueDTO{resultValue})
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

func getValue(operator string, num1, num2 int) (int, error) {
	switch {
	case strings.HasPrefix(operator, "add"):
		return num1 + num2, nil
	case strings.HasPrefix(operator, "sub"):
		return num1 - num2, nil
	case strings.HasPrefix(operator, "mul"):
		return num1 * num2, nil
	case strings.HasPrefix(operator, "div"):
		if num2 == 0 {
			return 0, errors.New("Second value must != 0")
		}
		return num1 / num2, nil
	default:
		return 0, nil
	}
}

/* OLD CODE
func AddHandler(num1, num2 int, w http.ResponseWriter, r *http.Request) {
	valueJson, err := json.Marshal(ValueDTO{num1 + num2})
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

func SubHandler(num1, num2 int, w http.ResponseWriter, r *http.Request) {
	valueJson, err := json.Marshal(ValueDTO{num1 + num2})
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

func MulHandler(num1, num2 int, w http.ResponseWriter, r *http.Request) {
	valueJson, err := json.Marshal(ValueDTO{num1 + num2})
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

func DivHandler(num1, num2 int, w http.ResponseWriter, r *http.Request) {
	valueJson, err := json.Marshal(ValueDTO{num1 + num2})
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
*/
