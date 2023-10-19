package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/billdev1958/billdev2.0.git/db"
)

// Handle func categories
func (s *Server) handleCategory(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreateCategory(w, r)
	}
	if r.Method == "GET" {
		return s.handleGetCategories(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
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

// Handle func articles
func (s *Server) handlePosts(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "POST" {
		return s.handleCreatePost(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)
}

// handler post articles
func (s *Server) handleCreatePost(w http.ResponseWriter, r *http.Request) error {
	var newPost db.Posts

	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		return err
	}

	if err := s.store.CreatePost(&newPost); err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, newPost)
}
