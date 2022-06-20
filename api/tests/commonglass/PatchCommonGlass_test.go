package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestPatchCommonGlass(t *testing.T) {

	reqBody := map[string]any{
		"id":    "f478d5670d344feead98ad4f492ae906",
		"price": 270,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PATCH", "http://localhost:3000/common-glasses", bytes.NewReader(body))
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

func TestPatchCommonGlassError(t *testing.T) {

	reqBody := map[string]any{
		"price": 200,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PATCH", "http://localhost:3000/common-glasses", bytes.NewReader(body))
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
