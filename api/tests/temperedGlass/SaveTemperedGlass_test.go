package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSaveTemperedGlass(t *testing.T) {

	reqBody := map[string]any{
		"name":         "Janela",
		"description":  "Janela 4 folhas 8mm fumê",
		"price":        500,
		"quantity":     1,
		"type":         "tempered",
		"color":        "smoked",
		"glass_sheets": 4,
		"milimeter":    8,
		"height":       1.20,
		"width":        1,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/tempered-glasses", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	if s := res.StatusCode; s != 201 {
		t.Fatalf("want status code 201, got %v", s)
	}

	res.Body.Close()
}

func TestSaveTemperedGlassError(t *testing.T) {
	t.SkipNow()
	reqBody := map[string]any{}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/tempered-glasses", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	if s := res.StatusCode; s != 400 {
		t.Fatalf("want status code 400, got %v", s)
	}

	res.Body.Close()
}
