package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Note struct {
	Title string `json:"Title"`
	Desc  string `json:"desc"`
	Done  bool   `json:"Done"`
}

var Notes []Note

type PageData struct {
	PageTitle string
	Notes     []Note
}

func handleRequests() {
	homePage := template.Must(template.ParseFiles("./templates/home.html"))
	notes := []Note{
		Note{Title: "TODO1", Desc: "asdf", Done: true},
		Note{Title: "TODO2", Desc: "fdsa", Done: false},
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.Path == "/addTodo" {
			title := r.FormValue("title")
			desc := r.FormValue("desc")
			newNote := Note{Title: title, Desc: desc, Done: false}
			notes = append(notes, newNote)
			fmt.Println("New TODO Created")
		} else if r.Method == "POST" && r.URL.Path == "/updateTodo" {
			title := r.FormValue("title")
			for i, note := range notes {
				if note.Title == title {
					notes[i].Done = true
					fmt.Println("TODO Updated")
					break
				}
			}
		}
		data := PageData{
			PageTitle: "--TODO List--",
			Notes:     notes,
		}
		homePage.Execute(w, data)
	})
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func main() {
	handleRequests()
}
