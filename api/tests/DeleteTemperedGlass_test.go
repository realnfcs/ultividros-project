package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDeleteTemperedGlass(t *testing.T) {
	reqBody := map[string]any{
		"id":           "5e78e83e52ff4f4fab13678a17d81a74",
		"name":         "Janela",
		"description":  "Janela 4 folhas 8mm fumê",
		"price":        500,
		"quantity":     1,
		"type":         "tempered",
		"color":        "smoked",
		"glass_sheets": 0,
		"milimeter":    8,
		"height":       1.2,
		"width":        1,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("DELETE", "http://localhost:3000/tempered-glasses", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Print(res)
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
