package server

import (
	"api-calc/handlers"
	"api-calc/utils"
	"net/http"
)

func MakeHandler(fn func(num1, num2 int, w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		mapParams := r.URL.Query()
		num1, err := utils.GetParamValue("num1", mapParams)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Param not valid"))
			return
		}

		num2, err := utils.GetParamValue("num2", mapParams)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("Param not valid"))
			return
		}

		fn(num1, num2, w, r)

	}
}

/*
	Sample API:
		+ http://localhost:8088/calc/add?num1=1&num2=2
		+ http://localhost:8088/calc/sub?num1=1&num2=2
		+ http://localhost:8088/calc/mul?num1=1&num2=2
		+ http://localhost:8088/calc/div?num1=1&num2=2
*/
func RunServer() {
	http.HandleFunc("/calc/add", MakeHandler(handlers.CalcHandler))
	http.HandleFunc("/calc/sub", MakeHandler(handlers.CalcHandler))
	http.HandleFunc("/calc/mul", MakeHandler(handlers.CalcHandler))
	http.HandleFunc("/calc/div", MakeHandler(handlers.CalcHandler))

	/* OLD CODE
	http.HandleFunc("/calc/add", MakeHandler(handlers.AddHandler))
	http.HandleFunc("/calc/sub", MakeHandler(handlers.SubHandler))
	http.HandleFunc("/calc/mul", MakeHandler(handlers.MulHandler))
	http.HandleFunc("/calc/div", MakeHandler(handlers.DivHandler))
	*/

	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic("Error when running server")
	}
}
