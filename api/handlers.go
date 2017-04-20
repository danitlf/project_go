package api

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/danitlf/project_go/db"
	"github.com/danitlf/project_go/music"
	"github.com/dimfeld/httptreemux"
)

type MyHandler struct{
	Repository *db.MusicRepository
}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := httptreemux.ContextParams(r.Context())
	my_music := &music.Music{}
	my_music, err := h.Repository.FindById(params["id"])
	if err != nil {
		fmt.Fprintf(w, "nao tem essa musica cadastrada %s", params["id"])
	}
	fmt.Fprintf(w, "O nome da musica é: %s!", my_music.Name)
	
}

type PutHandler struct{
	  Repository *db.MusicRepository
}

func (h *PutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	music := &music.Music{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(music)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	

	h.Repository.Create(music)
	

}

