package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type todo struct {
	UserId    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	http.HandleFunc("/", appHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/**/
func appHandler(w http.ResponseWriter, r *http.Request) {
	const baseUrl = "https://jsonplaceholder.typicode.com"

	res, err := http.Get(baseUrl + "/todos/" + r.URL.Path[1:])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer res.Body.Close()
	var todoI todo

	if err = json.NewDecoder(res.Body).Decode(&todoI); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		create html template
			- create the todo html structure  #1
			- inject it to the template  #2
			- parsing, then executing  #3
	*/

	// #1
	var todoStructure = `
		<body style="text-align: center; font-family: Dank mono, Monospace">
			<h1>Todo number #{{.ID}}</h1>
			<h3>{{.Title}}</h3>	
			<hr />
			<div>{{printf "User ID: %d" .UserId}}</div>
			<h4>{{printf "Completed? %t" .Completed}}</h4>
		</body>
	`

	// #2
	tmpl := template.New("todo")

	// #3
	tmpl.Parse(todoStructure)
	tmpl.Execute(w, todoI)
}
