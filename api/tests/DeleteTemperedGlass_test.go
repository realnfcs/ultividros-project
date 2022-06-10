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
		"id":           "718f4814b432453fb28526b514ede25b",
		"name":         "basculhante",
		"description":  "Um peça de vidro temperado muito comum em banheiros",
		"price":        180,
		"quantity":     3,
		"type":         "tempered",
		"color":        "smoked",
		"glass_sheets": 0,
		"milimeter":    8,
		"height":       0.5,
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
