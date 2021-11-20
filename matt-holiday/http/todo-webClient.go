package main

import (
	"encoding/json"
	"html/template"
	"io"
	"log"
	"net/http"
)

type todoBluePrint struct {
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

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}(res.Body)
	var todoI todoBluePrint

	if err = json.NewDecoder(res.Body).Decode(&todoI); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		create html template
			- create the _todo html structure  #1
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
	_, err = tmpl.Parse(todoStructure)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = tmpl.Execute(w, todoI)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
