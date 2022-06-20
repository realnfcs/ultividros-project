package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestDeleteCommonGlass(t *testing.T) {

	reqBody := map[string]any{
		"id": "f478d5670d344feead98ad4f492ae906",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(body)
	}

	req, err := http.NewRequest("DELETE", "http://localhost:3000/common-glasses", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(res)
	}

	defer req.Body.Close()
	defer res.Body.Close()

	if s := res.StatusCode; s != 200 {
		t.Fatalf("want status code 200, got %v", s)
	}
}

func TestDeleteCommonGlassError(t *testing.T) {

	reqBody := map[string]any{}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "http://localhost:3000/common-glasses", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(res)
	}

	defer req.Body.Close()
	defer res.Body.Close()

	if s := res.StatusCode; s != 404 {
		t.Fatalf("want status code 404, got %v", s)
	}
}
