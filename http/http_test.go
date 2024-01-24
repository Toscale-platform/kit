package http

import "testing"

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

func TestPost(t *testing.T) {
	todo := Todo{}
	body := TodoAdd{
		Todo:      "Toscale",
		Completed: false,
		UserID:    5,
	}

	err := Post("https://dummyjson.com/todos/add", &body, &todo)
	if err != nil {
		t.Error(err)
	}

	if todo.Todo != "Toscale" {
		t.Error("todo name not equal Toscale")
	}

	if todo.Completed != false {
		t.Error("todo completed not equal false")
	}

	if todo.UserID != 5 {
		t.Error("todo user id not equal 5")
	}
}
