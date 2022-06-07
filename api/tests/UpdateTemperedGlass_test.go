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
		"id":           "184751d5-fa3f-4837-b3de-4831164aac1b",
		"name":         "Janela",
		"description":  "Uma janela temperada de 4 folhas",
		"price":        430,
		"quantity":     5,
		"type":         "Tempered",
		"color":        "smoked",
		"glass_sheets": 4,
		"milimeter":    8,
		"height":       1.20,
		"width":        1.0,
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
	fmt.Print(res)
	if res.StatusCode != 200 {
		t.Fatal("some error happened")
	}

	defer req.Body.Close()
	defer res.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(res.Body)

	// Conversão do request body em bytes para string
	bodyString := string(bodyBytes)

	fmt.Println(bodyString)
}
