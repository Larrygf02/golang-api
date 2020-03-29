package main

import (
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
	fmt.Fprintf(w, "Bienvenido API en go")
}

func main() {
	// Configuración inicial de la ruta
	router := mux.NewRouter().StrictSlash(true)

	// Definiendo una ruta
	router.HandleFunc("/", indexRoute)

	// Iniciando un servicdor
	// logs
	log.Fatal(http.ListenAndServe(":3000", router))
}
