package tests

import (
	"net/http"
	"testing"
)

func TestGetCommonGlass(t *testing.T) {

	res, err := http.Get("http://localhost:3000/common-glasses/id=134fccc5de2947168ddc874d3c2b844a")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 200 {
		t.Fatalf("wanted status code 200, got %v", s)
	}
}

func TestGetCommonGlassError(t *testing.T) {

	res, err := http.Get("http://localhost:3000/common-glasses/id=123")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 404 {
		t.Fatalf("wanted status code 404, got %v", s)
	}
}
