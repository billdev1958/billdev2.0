package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/billdev1958/billdev2.0.git/db"
)

func (s *Server) handlePosts(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateCategory(w, r)
	}
	if r.Method == "GET" {
		return s.handleGetCategories(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

// Handler post user
func (s *Server) handleCreateUser(w http.ResponseWriter, r *http.Request) error {
	req := new(db.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	user, err := db.NewUser(
		req.Nombre,
		req.Username,
		req.Email,
		req.Password,
		req.Hashed_password,
		req.Resume,
	)

	if err != nil {
		return err
	}
	if err := s.store.CreateUser(user); err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, user)
}

// handler post category
func (s *Server) handleCreateCategory(w http.ResponseWriter, r *http.Request) error {
	req := new(db.CreateCategoryRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}
	category, err := db.NewCategory(
		req.Nombre,
	)

	if err != nil {
		return err
	}
	if err := s.store.CreateCategory(category); err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, category)
}

// handler get categories
func (s *Server) handleGetCategories(w http.ResponseWriter, r *http.Request) error {
	categories, err := s.store.GetCategories()
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, categories)
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
