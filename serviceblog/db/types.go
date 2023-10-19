package db

import (
	"time"
)

type User struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Hashed_password string `json:"hashed_password"`
	Resume          string `json:"resume"`
}

type CreateUserRequest struct {
	Nombre          string `json:"nombre"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Hashed_password string `json:"hashed_password"`
	Resume          string `json:resume"`
}

func NewUser(nombre, username, email, password, hashed_password, resume string) (*User, error) {
	return &User{
		Nombre:          nombre,
		Username:        username,
		Email:           email,
		Password:        password,
		Hashed_password: hashed_password,
		Resume:          resume,
	}, nil
}

type Category struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
}

type CreateCategoryRequest struct {
	Nombre string `json:"nombre"`
}

func NewCategory(nombre string) (*Category, error) {
	return &Category{
		Nombre: nombre,
	}, nil
}

type Posts struct {
	ID        int       `json:"id"`
	Category  string    `json:"category"`
	Title     string    `json:"title"`
	Resume    string    `json:"resume"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreatePostParams struct {
	Category int    `json:"category"`
	Title    string `json:"title"`
	Resume   string `json:"resume"`
	Content  string `json:"content"`
	Author   int    `json:"author"`
}

type CreatePostResponse struct {
	Category int    `json:"category"`
	Title    string `json:"title"`
	Resume   string `json:"resume"`
	Content  string `json:"content"`
	Author   int    `json:"author"`
}
