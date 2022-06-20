package tests

import (
	"net/http"
	"testing"
)

func TestGetCommonGlasses(t *testing.T) {
	res, err := http.Get("http://localhost:3000/common-glasses")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 200 {
		t.Fatalf("want status code 200, got %v", s)
	}
}
