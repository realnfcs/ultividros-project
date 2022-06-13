package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestFiber(t *testing.T) {
	res, err := http.Get("http://localhost:3000/tempered-glasses")
	if err != nil {
		t.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if s := res.StatusCode; s != 200 {
		t.Fatalf("want status code 200, got %v", s)
	}

	fmt.Println(string(body))
}
