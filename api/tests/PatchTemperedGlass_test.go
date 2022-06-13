package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestPatchTemperedGlasses(t *testing.T) {
	reqBody := map[string]any{
		"id":    "09a49643554f462da01cc9ec471af7d6",
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
