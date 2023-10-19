package api

import (
	"encoding/json"
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

	router.HandleFunc("/category", makeHTTPHandleFunc(s.handleCategory))
	router.HandleFunc("/user", makeHTTPHandleFunc(s.handleUsers))
	router.HandleFunc("/article", makeHTTPHandleFunc(s.handlePosts))
	log.Println("JSON API server runing on port:", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

// structure handlerfunc
type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string
}

// Make handler function for api
func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJson(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

// encode json function
func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "aplication/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
