package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCloseSale(t *testing.T) {

	reqBody := map[string]any{
		"id":        "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
		"client_id": "b5bc889341ba48549223dcfc325255e5",
		"is_active": true,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PATCH", "http://localhost:3000/sales/close", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 200 {
		t.Fatalf("want status code 200, got %v", s)
	}
}

func TestCloseSaleError(t *testing.T) {

	reqBody := map[string]any{
		"id": "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PATCH", "http://localhost:3000/sales/close", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 400 {
		t.Fatalf("want status code 400, got %v", s)
	}
}
