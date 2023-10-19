package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/billdev1958/billdev2.0.git/db"
)

// Handle func users
func (s *Server) handleUsers(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handleGetUsers(w, r)
	}
	if r.Method == "POST" {
		return s.handleCreateUser(w, r)
	}
	return fmt.Errorf("method not allowed %s", r.Method)

}

type createUserRequest struct {
	Nombre          string `json:"nombre"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Hashed_password string `json:"hashed_password"`
	Resume          string `json:resume"`
}

type userResponse struct {
	Nombre   string `json:"nombre"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserResponse(user db.User) userResponse {
	return userResponse{
		Username: user.Username,
		Nombre:   user.Nombre,
		Email:    user.Email,
	}
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

// handler get users
func (s *Server) handleGetUsers(w http.ResponseWriter, r *http.Request) error {
	users, err := s.store.GetUsers()
	if err != nil {
		return err
	}
	return WriteJson(w, http.StatusOK, users)
}
