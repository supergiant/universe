package routectl

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type component struct {
	Name      string `json:"name"`
	Location  string `json:"location"`
	Readme    string `json:"readme"`
	Downloads int    `json:"downloads"`
}

// Index returns a simple top 10 list of most popular components.
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Example provides a simple example in json of a object response.
func Example(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	example := component{
		Name:      "example",
		Location:  "github.com/example/example_component",
		Readme:    "github.com/example/exaple_component/blob/master/README.md",
		Downloads: 34,
	}

	js, err := json.MarshalIndent(example, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
