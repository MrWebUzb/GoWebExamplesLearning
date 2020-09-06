package main

import (
	"net/http"
	"text/template"
)

// Todo type
type Todo struct {
	Title       string
	IsCompleted bool
}

// TodoPage for todo page template
type TodoPage struct {
	PageTitle string
	TodoList  []Todo
}

func main() {
	tpl := template.Must(template.ParseFiles("main.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		todoPage := &TodoPage{
			PageTitle: "TodoList program in Golang",
			TodoList: []Todo{
				{
					Title:       "Transfer.edu.uz uchun pdf generatsiya qilish",
					IsCompleted: true,
				},
				{
					Title:       "Golangda todolist dasturini tuzish",
					IsCompleted: true,
				},
				{
					Title:       "Universitetda berilgan vazifalarni bajarish",
					IsCompleted: false,
				},
			},
		}

		tpl.Execute(w, todoPage)
	})

	http.ListenAndServe(":8080", nil)
}
