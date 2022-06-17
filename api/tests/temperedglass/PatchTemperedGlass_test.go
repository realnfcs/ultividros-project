package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestPatchTemperedGlasses(t *testing.T) {

	reqBody := map[string]any{
		"id":    "cac5882dca71479fb94ce120f631f36e",
		"price": 1000,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PATCH", "http://localhost:3000/tempered-glasses", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if s := res.StatusCode; s != 200 {
		t.Fatalf("want status code 200, got %v", s)
	}

	defer req.Body.Close()
	defer res.Body.Close()

	// Caso queria ver o body da response
	/*
		bodyBytes, _ := ioutil.ReadAll(res.Body)

		// Conversão do request body em bytes para string
		bodyString := string(bodyBytes)

		fmt.Println(bodyString)
	*/
}

func TestPatchTemperedGlassesError(t *testing.T) {
	reqBody := map[string]any{
		"price": 900,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PATCH", "http://localhost:3000/tempered-glasses", bytes.NewReader(body))
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
