package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateUser(*User) error
	GetUsers() ([]*User, error)
	CreateCategory(*Category) error
	GetCategories() ([]*Category, error)
	CreatePost(*Posts) error
}

type PostgreStore struct {
	db *sql.DB
}

func NewPostgreStore() (*PostgreStore, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgreStore{
		db: db,
	}, nil
}

// Create User
func (s *PostgreStore) CreateUser(user *User) error {
	query := `INSERT INTO users
	(nombre, username, email, password, hashed_password, resume)
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := s.db.Query(
		query,
		user.Nombre,
		user.Username,
		user.Email,
		user.Password,
		user.Hashed_password,
		user.Resume,
	)
	if err != nil {
		return err
	}
	return nil
}

// Get all users
func (s *PostgreStore) GetUsers() ([]*User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	users := []*User{}

	for rows.Next() {
		user := new(User)
		err := rows.Scan(
			&user.ID,
			&user.Nombre,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Hashed_password,
			&user.Resume,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Create category
func (s *PostgreStore) CreateCategory(category *Category) error {
	query := `INSERT INTO category
	(nombre)
	VALUES($1)`

	_, err := s.db.Query(
		query,
		category.Nombre,
	)
	if err != nil {
		return err
	}
	return nil
}

// Get all category
func (s *PostgreStore) GetCategories() ([]*Category, error) {
	rows, err := s.db.Query("SELECT * FROM category")
	if err != nil {
		return nil, err
	}
	categories := []*Category{}

	for rows.Next() {
		category := new(Category)
		err := rows.Scan(
			&category.ID,
			&category.Nombre)

		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// Create post
func (s *PostgreStore) CreatePost(posts *Posts) error {
	query := `INSERT INTO posts
	(category, title, resume, content, author, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)`
	resp, err := s.db.Query(
		query,
		posts.Category,
		posts.Title,
		posts.Resume,
		posts.Content,
		posts.Author,
		posts.CreatedAt,
	)
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", resp)
	return nil
}
