package tests

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestGetTemperedGlass(t *testing.T) {
	res, err := http.Get("http://localhost:3000/tempered-glasses/id=15f0002cf27744efa213caf18e7ae54c")
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
