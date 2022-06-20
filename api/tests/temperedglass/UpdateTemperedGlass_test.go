package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestUpdateTemperedGlasses(t *testing.T) {

	reqBody := map[string]any{
		"id":           "cac5882dca71479fb94ce120f631f36e",
		"name":         "Test",
		"description":  "Descrição test",
		"price":        100,
		"quantity":     1,
		"type":         "tempered",
		"color":        "transparent",
		"glass_sheets": 1,
		"milimeter":    10,
		"height":       1,
		"width":        1,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "http://localhost:3000/tempered-glasses", bytes.NewReader(body))
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

	// Caso queira ver o body da response //
	/*
		bodyBytes, _ := ioutil.ReadAll(res.Body)

		// Conversão do request body em bytes para string
		bodyString := string(bodyBytes)

		fmt.Println(bodyString)
	*/
}

func TestUpdateTemperedGlassesError(t *testing.T) {
	reqBody := map[string]any{
		"id": "09a49643554f462da01cc9ec471af7d6",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PUT", "http://localhost:3000/tempered-glasses", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if s := res.StatusCode; s != 400 {
		t.Fatalf("want status code 400, got %v", s)
	}

	defer req.Body.Close()
	defer res.Body.Close()
}
