package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestDeletePart(t *testing.T) {

	reqBody := map[string]any{
		"id": "3c024f267e5b444c945a97186ec68682",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(body)
	}

	req, err := http.NewRequest("DELETE", "http://localhost:3000/parts", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer req.Body.Close()
	defer res.Body.Close()

	if s := res.StatusCode; s != 200 {
		t.Fatalf("want status code 200, got %v", s)
	}
}

func TestDeletePartError(t *testing.T) {

	reqBody := map[string]any{}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(body)
	}

	req, err := http.NewRequest("DELETE", "http://localhost:3000/parts", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer req.Body.Close()
	defer res.Body.Close()

	if s := res.StatusCode; s != 404 {
		t.Fatalf("want status code 404, got %v", s)
	}
}
