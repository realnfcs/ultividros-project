package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSaveCommonGlass(t *testing.T) {

	reqBody := map[string]any{
		"name":             "Martelado",
		"description":      "Vidro martelado transparente de 6mm",
		"price":            250,
		"type":             "hammered",
		"color":            "colorless",
		"milimeter":        6,
		"height_available": 2,
		"width_available":  3.5,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/common-glasses", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	if s := res.StatusCode; s != 201 {
		t.Fatalf("want status code 201, got %v", s)
	}

	res.Body.Close()
}

func TestSaveCommonGlassError(t *testing.T) {

	reqBody := map[string]any{}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/common-glasses", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	if s := res.StatusCode; s != 400 {
		t.Fatalf("want status code 400, got %v", s)
	}

	res.Body.Close()
}
