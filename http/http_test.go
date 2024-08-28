package http

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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
	assert.Nil(t, err)
	assert.Equal(t, todo.ID, 1)
}

func TestNilVGet(t *testing.T) {
	err := Get("https://dummyjson.com/todos/1", nil)
	assert.Nil(t, err)
}

func TestGetWithToken(t *testing.T) {
	user := User{}
	body := Login{
		Username: "kminchelle",
		Password: "0lelplR",
	}

	err := Post("https://dummyjson.com/auth/login", &body, &user)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, 15)

	token := user.Token
	user = User{}

	err = GetWithToken("https://dummyjson.com/auth/me", token, &user)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, 15)
}

func TestPost(t *testing.T) {
	todo := Todo{}
	body := TodoAdd{
		Todo:      "Test",
		Completed: true,
		UserID:    5,
	}

	err := Post("https://dummyjson.com/todos/add", &body, &todo)
	assert.Nil(t, err)
	assert.Equal(t, todo.Todo, "Test")
	assert.True(t, todo.Completed)
	assert.Equal(t, todo.UserID, 5)
}
