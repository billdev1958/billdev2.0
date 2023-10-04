package api

import (
	"log"
	"net/http"

	"github.com/billdev1958/billdev2.0.git/db"
	"github.com/gorilla/mux"
)

type Server struct {
	store      db.Storage
	listenAddr string
}

func NewServer(listenAddr string, store db.Storage) *Server {
	return &Server{
		listenAddr: listenAddr,
		store:      store,
	}

}

func (s Server) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/category", makeHTTPHandleFunc(s.handlePosts))

	log.Println("JSON API server runing on port:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}
