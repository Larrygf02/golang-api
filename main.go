package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json: ID`
	Name    string `json: Name`
	Content string `json: Content`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Task one",
		Content: "Some content",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido API en golang")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func main() {
	// Configuración inicial de la ruta
	router := mux.NewRouter().StrictSlash(true)

	// Definiendo una ruta
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks)

	// Iniciando un servicdor
	log.Fatal(http.ListenAndServe(":3000", router))
}
