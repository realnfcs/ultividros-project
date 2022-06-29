package tests

import (
	"net/http"
	"testing"
)

func TestGetUser(t *testing.T) {

	res, err := http.Get("http://localhost:3000/users/id=b5bc889341ba48549223dcfc325255e5")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 200 {
		t.Fatalf("wanted status code 200, got %v", s)
	}
}

func TestGetUserError(t *testing.T) {

	res, err := http.Get("http://localhost:3000/users/id=123")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 404 {
		t.Fatalf("wanted status code 404, got %v", s)
	}
}
