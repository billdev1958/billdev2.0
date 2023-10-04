package db

type User struct {
	ID              int    `json:"id"`
	Nombre          string `json:"nombre"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Hashed_password string `json:"hashed_password"`
	Resume          string `	json:"resume"`
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
	ID       int
	Category int
	Title    string
	Resume   string
	Content  string
	Author   int
}

type CreatePostRequest struct {
	Category int
	Title    string
	Resume   string
	Content  string
	Author   int
}

/*func NewPost(title, resume, content, category, author string) (*Posts, error) {
	return &Posts{
		Category:  category,
		Title:     title,
		Resume:    resume,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now().UTC(),
	}, nil
}*/
