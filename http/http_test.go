package http

import "testing"

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Gender    string `json:"gender"`
	Image     string `json:"image"`
	Token     string `json:"token"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Todo struct {
	ID        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"userId"`
}

type TodoAdd struct {
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	UserID    int    `json:"userId"`
}

func TestGet(t *testing.T) {
	todo := Todo{}

	err := Get("https://dummyjson.com/todos/1", &todo)
	if err != nil {
		t.Error(err)
	}

	if todo.ID != 1 {
		t.Error("todo id not equal 1")
	}
}

func TestNilVGet(t *testing.T) {
	err := Get("https://dummyjson.com/todos/1", nil)
	if err != nil {
		t.Error(err)
	}
}

func TestGetWithToken(t *testing.T) {
	user := User{}
	body := Login{
		Username: "kminchelle",
		Password: "0lelplR",
	}

	err := Post("https://dummyjson.com/auth/login", &body, &user)
	if err != nil {
		t.Error(err)
	}

	if user.ID != 15 {
		t.Error("user id not equal 15")
	}

	token := user.Token
	user = User{}

	err = GetWithToken("https://dummyjson.com/auth/me", token, &user)
	if err != nil {
		t.Error(err)
	}

	if user.ID != 15 {
		t.Error("user id not equal 15")
	}
}

func TestPost(t *testing.T) {
	todo := Todo{}
	body := TodoAdd{
		Todo:      "Test",
		Completed: true,
		UserID:    5,
	}

	err := Post("https://dummyjson.com/todos/add", &body, &todo)
	if err != nil {
		t.Error(err)
	}

	if todo.Todo != "Test" {
		t.Error("todo name not equal Test")
	}

	if !todo.Completed {
		t.Error("todo completed not equal true")
	}

	if todo.UserID != 5 {
		t.Error("todo user id not equal 5")
	}
}
