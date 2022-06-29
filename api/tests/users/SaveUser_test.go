package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSaveUser(t *testing.T) {

	reqBody := map[string]any{
		"name":       "test",
		"email":      "test@test.com",
		"password":   "123test",
		"occupation": "user",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/users", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	if s := res.StatusCode; s != 201 {
		t.Fatalf("want status code 201, got %v", s)
	}

	res.Body.Close()
}

func TestSaveUserError(t *testing.T) {

	reqBody := map[string]any{}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/parts", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	if s := res.StatusCode; s != 400 {
		t.Fatalf("want status code 400, got %v", s)
	}

	res.Body.Close()
}
