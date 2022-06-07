package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetTemperedGlass(t *testing.T) {
	res, err := http.Get("http://localhost:3000/tempered-glasses/id=b220daae-f3e4-4aa4-9cd5-be00dcb1325f")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(string(body))
}
