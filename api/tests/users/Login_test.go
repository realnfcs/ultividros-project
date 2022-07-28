package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestLogin(t *testing.T) {

	reqBody := map[string]any{
		"email":    "testbot@test.com",
		"password": "123testbot",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/users/login", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 200 {
		t.Fatalf("expect status code 200, got %v", s)
	}
}

func TestLoginError(t *testing.T) {

	reqBody := map[string]any{
		"email":    "testbot@test.com",
		"password": "12345testbot",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/users/login", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 401 {
		t.Fatalf("expect status code 401, got %v", s)
	}
}
