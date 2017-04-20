package main

	// "github.com/danitlf/project_go/db"
	// "github.com/danitlf/project_go/music"

import (
	"github.com/danitlf/project_go/api"
	"github.com/danitlf/project_go/db"
	"log"
	"net/http"
	"gopkg.in/mgo.v2"
	"github.com/dimfeld/httptreemux"
)

// type UpsertCarHandler struct{}

// func (h *UpsertCarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	params := httptreemux.ContextParams(r.Context())
// 	fmt.Fprintf(w, "Eu deveria criar um carro chamado: %s!", params["id"])
// 	fmt.Fprintln(w, "Não crio por que sou mal!")
// }


func main() {
	session, err := mgo.Dial("localhost:27017/go-course")
	Repository := db.NewMusicRepository(session)
	if err != nil {
		log.Fatal(err)
	}

	addr := "127.0.0.1:8081"
	router := httptreemux.NewContextMux()
	router.Handler(http.MethodGet, "/music/:id", &api.MyHandler{Repository})
	router.Handler(http.MethodPost, "/music/", &api.PutHandler{Repository})
	// router.Handler(http.MethodGet, "/daniel/:num", &DanielHandler{})

	log.Printf("Running web server on: http://%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))

	// execute
	// curl http://localhost:8081/cars/gol
	// curl -XPUT http://localhost:8081/cars/fusca -d'{"name": 1}'
}
