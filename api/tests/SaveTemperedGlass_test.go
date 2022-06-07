package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSaveTemperedGlass(t *testing.T) {
	reqBody := map[string]any{
		"name":        "basculhante",
		"description": "Um peça de vidro temperado muito comum em banheiros",
		"price":       180,
		"quantity":    3,
		"type":        "tempered",
		"color":       "smoked",
		"GlassSheets": 1,
		"milimeter":   8,
		"height":      0.5,
		"width":       1,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/tempered-glasses", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	if res.StatusCode != 200 {
		t.Fatal("some error happened")
	}

	res.Body.Close()
}
