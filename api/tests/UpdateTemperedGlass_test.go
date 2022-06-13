package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestUpdateTemperedGlasses(t *testing.T) {
	reqBody := map[string]any{
		"id":           "09a49643554f462da01cc9ec471af7d6",
		"name":         "Porta",
		"description":  "Uma porta transparente 1 folhas 10mm",
		"price":        430,
		"quantity":     3,
		"type":         "tempered",
		"color":        "transparent",
		"glass_sheets": 1,
		"milimeter":    10,
		"height":       2.5,
		"width":        1.4,
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

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	// Conversão do request body em bytes para string
	bodyString := string(bodyBytes)

	fmt.Println(bodyString)
}
