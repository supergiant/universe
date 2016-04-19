package routectl

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"golang.org/x/net/context"

	"github.com/coreos/etcd/client"
	"github.com/julienschmidt/httprouter"
)

func etcd() (client.KeysAPI, error) {
	cfg := client.Config{
		Endpoints: []string{"http://localhost:2379"},
		Transport: client.DefaultTransport,
		// set timeout per request to fail fast when the target endpoint is unavailable
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	kapi := client.NewKeysAPI(c)

	return kapi, nil
}

func attemptToAdd(location string, db client.KeysAPI) *client.Response {
	sl := strings.Split(location, "/")
	name := sl[len(sl)-1]
	readme := "" + location + "/blob/master/README.md"

	new := component{
		Name:      name,
		Location:  location,
		Readme:    readme,
		Downloads: 0,
	}

	js, _ := json.MarshalIndent(new, "", "  ")

	resp, err := db.Set(context.Background(), "/"+name+"", string(js), nil)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// Search is the main entrypoint all for the whole API. If a searched component
// exists in the db, it will return. If not, it will add it if possible.
func Search(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	db, err := etcd()
	if err != nil {
		fmt.Fprintf(w, "Error:, %s!\n", err.Error())
	}

	resp, err := db.Get(context.Background(), "/"+ps.ByName("component")+"", nil)
	if err != nil {
		log.Println("Component:", ps.ByName("component"), "Not Found, we will attempt to add it.")
		resp := attemptToAdd(ps.ByName("component"), db)
		fmt.Fprintf(w, "%s", resp.Node.Value)
	} else {
		fmt.Fprintf(w, "%s", resp.Node.Value)
	}

}
