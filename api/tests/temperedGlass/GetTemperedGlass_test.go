package tests

import (
	"net/http"
	"testing"
)

func TestGetTemperedGlass(t *testing.T) {

	res, err := http.Get("http://localhost:3000/tempered-glasses/id=cac5882dca71479fb94ce120f631f36e")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	// Caso queira ver o body da response //
	/*
		body, err := io.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
	*/

	if s := res.StatusCode; s != 200 {
		t.Fatalf("wanted status code 200, got %v", s)
	}
}

func TestGetTemperedGlassError(t *testing.T) {

	// Invalid ID in URL
	res, err := http.Get("http://localhost:3000/tempered-glasses/id=123")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 404 {
		t.Fatalf("wanted status code 404, got %v", s)
	}
}
