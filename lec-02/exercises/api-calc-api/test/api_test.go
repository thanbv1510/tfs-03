package test

import (
	"api-calc/handlers"
	"api-calc/server"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddNumber(t *testing.T) {
	req, err := http.NewRequest("GET", "/calc/add?num1=10&num2=11", nil)
	fmt.Println(req.URL.String())

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.MakeHandler(handlers.CalcHandler))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Fail, code = %v", status)
	}

	// Check response
	expected := `{"value":21}`
	if rr.Body.String() != expected {
		t.Errorf("Fail with body = %s, expected body = %s", rr.Body.String(), expected)
	}
}

func TestSubNumber(t *testing.T) {
	req, err := http.NewRequest("GET", "/calc/sub?num1=10&num2=11", nil)
	fmt.Println(req.URL.String())

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.MakeHandler(handlers.CalcHandler))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Fail, code = %v", status)
	}

	// Check response
	expected := `{"value":-1}`
	if rr.Body.String() != expected {
		t.Errorf("Fail with body = %s, expected body = %s", rr.Body.String(), expected)
	}
}
func TestMulNumber(t *testing.T) {
	req, err := http.NewRequest("GET", "/calc/mul?num1=10&num2=11", nil)
	fmt.Println(req.URL.String())

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.MakeHandler(handlers.CalcHandler))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Fail, code = %v", status)
	}

	// Check response
	expected := `{"value":110}`
	if rr.Body.String() != expected {
		t.Errorf("Fail with body = %s, expected body = %s", rr.Body.String(), expected)
	}
}
func TestDivNumber(t *testing.T) {
	req, err := http.NewRequest("GET", "/calc/div?num1=10&num2=11", nil)
	fmt.Println(req.URL.String())

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(server.MakeHandler(handlers.CalcHandler))
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Fail, code = %v", status)
	}

	// Check response
	expected := `{"value":0}`
	if rr.Body.String() != expected {
		t.Errorf("Fail with body = %s, expected body = %s", rr.Body.String(), expected)
	}
}
