package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestSaveSale(t *testing.T) {

	reqBody := map[string]any{
		"client_id": "ecc24f56966342e4a335edb28da0d08b",
		"parts_req": []map[string]any{
			{
				"product_id":    "de54315e12a54a5b88876cb50da53677",
				"product_name":  "Bico de tucano",
				"prod_qty_req":  2,
				"product_price": 20,
			},
		},
		"temp_glss_req": []map[string]any{
			{
				"product_id":    "ba12aedffbd2440f815d7f197c7fd080",
				"product_name":  "Janela",
				"prod_qty_req":  1,
				"product_price": 500,
			},
		},
		"common_glss_req": []map[string]any{
			{
				"product_id":     "134fccc5de2947168ddc874d3c2b844a",
				"product_name":   "Fumê",
				"prod_qty_req":   1,
				"product_price":  200,
				"request_width":  1,
				"request_height": 1,
			},
		},
		"is_active": true,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/sales", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 201 {
		t.Fatalf("want status code 201, got %v", s)
	}
}

func TestSaveSaleError(t *testing.T) {

	reqBody := map[string]any{
		"client_id": "",
		"common_glss_req": []map[string]any{
			{
				"product_price":  200,
				"request_height": 1,
			},
		},
		"is_active": true,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	res, err := http.Post("http://localhost:3000/sales", "application/json", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 400 {
		t.Fatalf("want status code 400, got %v", s)
	}
}
