package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetTemperedGlass(t *testing.T) {

	res, err := http.Get("http://localhost:3000/tempered-glasses/id=09a49643554f462da01cc9ec471af7d6")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if s := res.StatusCode; s != 200 {
		t.Fatalf("wanted status code 200, got %v", s)
	}

	fmt.Println(string(body))
}
