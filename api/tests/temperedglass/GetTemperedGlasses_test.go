package tests

import (
	"net/http"
	"testing"
)

func TestGetTemperedGlasses(t *testing.T) {
	res, err := http.Get("http://localhost:3000/tempered-glasses")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	// Caso queira ver o body da response //
	/*
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
	*/

	if s := res.StatusCode; s != 200 {
		t.Fatalf("want status code 200, got %v", s)
	}
}

func TestGetTemperedGlassesError(t *testing.T) {
	t.SkipNow()
	res, err := http.Get("http://localhost:3000/tempered-glasses")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	if s := res.StatusCode; s != 404 {
		t.Fatalf("want status code 404, got %v", s)
	}
}
