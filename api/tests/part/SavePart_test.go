package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSavePart(t *testing.T) {

	reqBody := map[string]any{
		"name":        "Trinco",
		"description": "Trinco para portas e vidros temperados",
		"price":       45.0,
		"quantity":    10,
		"for_type":    "tempered",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/parts", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	if s := res.StatusCode; s != 201 {
		t.Fatalf("want status code 201, got %v", s)
	}

	res.Body.Close()
}

func TestSavePartError(t *testing.T) {

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
