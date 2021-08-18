package middleware

import (
	"awesomebook/helpers"
	"fmt"
	"net/http"
)

var sugar = helpers.GetSugar()

const JsonContentType = "application/json"

func ContentTypeCheckingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		reqContentType := request.Header.Get("Content-Type")
		if JsonContentType != reqContentType {
			fmt.Fprintf(writer, "request only allow content type %s", JsonContentType)
			sugar.Errorf("conten type %s not valid", reqContentType)
			return
		}

		next.ServeHTTP(writer, request)
	})
}
