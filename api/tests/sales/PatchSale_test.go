package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestPatchSale(t *testing.T) {

	reqBody := map[string]any{
		"id":        "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
		"client_id": "b5bc889341ba48549223dcfc325255e5",
		"common_glss_req": []map[string]any{
			{
				"id":             "537896a682ed4eedbbcd02da526d7649",
				"product_id":     "134fccc5de2947168ddc874d3c2b844a",
				"product_name":   "Fumê",
				"product_price":  200,
				"prodt_qty_req":  4,
				"was_cancelled":  false,
				"was_confimed":   true,
				"sale_id":        "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
				"request_width":  1,
				"request_height": 1,
			},
		},
		"parts_req": []map[string]any{
			{
				"id":            "7d7d49f121e840d8bb11e029bab3f3b3",
				"product_id":    "de54315e12a54a5b88876cb50da53677",
				"product_name":  "Bico de tucano",
				"product_price": 20,
				"prodt_qty_req": 2,
				"was_cancelled": true,
				"was_confimed":  false,
				"sale_id":       "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
			},
			{
				"id":            "7f450acba5e84c2f95f4b3ab3cd27a93",
				"product_id":    "04c7d0dff4a049f682c7c9d9c995aebc",
				"product_name":  "Puxador",
				"product_price": 175,
				"prodt_qty_req": 1,
				"was_cancelled": false,
				"was_confimed":  false,
				"sale_id":       "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
			},
		},
		"temp_glss": []map[string]any{
			{
				"id":            "a9966c8a8a9c4a89a5b40cfe6c9efbc7",
				"product_id":    "ba12aedffbd2440f815d7f197c7fd080",
				"product_name":  "Janela",
				"product_price": 500,
				"prodt_qty_req": 4,
				"was_cancelled": true,
				"was_confimed":  false,
				"sale_id":       "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
			},
		},
		"is_active": true,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PATCH", "http://localhost:3000/sales", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 200 {
		t.Fatalf("want status code 200, got %v", s)
	}
}

func TestPatchSaleError(t *testing.T) {

	reqBody := map[string]any{
		"id": "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PATCH", "http://localhost:3000/sales", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 400 {
		t.Fatalf("want status code 400, got %v", s)
	}
}

func TestPatchSaleError2(t *testing.T) {

	reqBody := map[string]any{
		"id":        "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
		"client_id": "b5bc889341ba48549223dcfc325255e5",
		"common_glss_req": []map[string]any{
			{
				"id":             "537896a682ed4eedbbcd02da526d7649",
				"product_id":     "134fccc5de2947168ddc874d3c2b844a",
				"product_name":   "Fumê",
				"product_price":  200,
				"prodt_qty_req":  4,
				"was_cancelled":  true,
				"was_confimed":   true,
				"sale_id":        "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
				"request_width":  1,
				"request_height": 1,
			},
		},
		"parts_req": []map[string]any{
			{
				"id":            "7d7d49f121e840d8bb11e029bab3f3b3",
				"product_id":    "de54315e12a54a5b88876cb50da53677",
				"product_name":  "Bico de tucano",
				"product_price": 20,
				"prodt_qty_req": 2,
				"was_cancelled": true,
				"was_confimed":  true,
				"sale_id":       "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
			},
			{
				"id":            "7f450acba5e84c2f95f4b3ab3cd27a93",
				"product_id":    "04c7d0dff4a049f682c7c9d9c995aebc",
				"product_name":  "Puxador",
				"product_price": 175,
				"prodt_qty_req": 1,
				"was_cancelled": true,
				"was_confimed":  true,
				"sale_id":       "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
			},
		},
		"temp_glss": []map[string]any{
			{
				"id":            "a9966c8a8a9c4a89a5b40cfe6c9efbc7",
				"product_id":    "ba12aedffbd2440f815d7f197c7fd080",
				"product_name":  "Janela",
				"product_price": 500,
				"prodt_qty_req": 4,
				"was_cancelled": true,
				"was_confimed":  true,
				"sale_id":       "7eaf3a4f58fe4efbbfbd7adebfe7efc6",
			},
		},
		"is_active": true,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("PATCH", "http://localhost:3000/sales", bytes.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}

	defer req.Body.Close()

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := new(http.Client)

	res, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 400 {
		t.Fatalf("want status code 400, got %v", s)
	}
}
