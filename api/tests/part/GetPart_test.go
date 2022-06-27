package tests

import (
	"net/http"
	"testing"
)

func TestGetPart(t *testing.T) {

	res, err := http.Get("http://localhost:3000/parts/id=04c7d0dff4a049f682c7c9d9c995aebc")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 200 {
		t.Fatalf("wanted status code 200, got %v", s)
	}
}

func TestGetPartError(t *testing.T) {

	res, err := http.Get("http://localhost:3000/parts/id=123")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 404 {
		t.Fatalf("wanted status code 200, got %v", s)
	}
}
