package tests

import (
	"net/http"
	"testing"
)

func TestGetSale(t *testing.T) {

	res, err := http.Get("http://localhost:3000/sales/id=fb856a24000a4cc889561b6bd0064d54")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 200 {
		t.Fatalf("wanted status code 200, got %v", s)
	}

}

func TestGetSaleError(t *testing.T) {

	res, err := http.Get("http://localhost:3000/sales/id=123")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 404 {
		t.Fatalf("wanted status code 404, got %v", s)
	}

}
