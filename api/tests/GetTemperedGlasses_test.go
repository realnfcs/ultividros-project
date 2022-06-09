package tests

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/realnfcs/ultividros-project/api/domain/usecases/temperedglasses/gettemperedglasses"
	"github.com/realnfcs/ultividros-project/api/infra/repository"
)

func TestExecute(t *testing.T) {
	memoryRepo := new(repository.TemperedGlassRepositoryMemory).Init()
	g := gettemperedglasses.GetTemperedGlasses{TemperedGlassRepository: memoryRepo}

	res := g.Execute()
	if res == nil {
		t.Fatalf("want a slice of Tempered Glasses, got %v", res)
	}
}

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

	fmt.Println(string(body))
}
