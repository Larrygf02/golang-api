package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json: ID`
	Name    string `json: Name`
	Content string `json: Content`
}

type resp struct {
	status int    `json: status`
	msg    string `json: msg`
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	// get body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte una tarea valida")
	}
	json.Unmarshal(reqBody, &newTask)
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newTask)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get parameters of url
	parameters := mux.Vars(r)
	taskID, err := strconv.Atoi(parameters["id"])
	if err != nil {
		fmt.Fprintf(w, "Id invalido")
		return
	}
	for _, task := range tasks {
		if task.ID == taskID {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	response := resp{
		status: 200,
		msg:    "No hay tarea con este id",
	}
	json.NewEncoder(w).Encode(response)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get parameters of url
	parameters := mux.Vars(r)
	taskID, err := strconv.Atoi(parameters["id"])
	if err != nil {
		fmt.Fprintf(w, "Id invalido")
		return
	}
	for i, task := range tasks {
		if task.ID == taskID {
			tasks = append(tasks[:i], tasks[i+i:]...)
			fmt.Fprintf(w, "La tarea ha sido eliminado correctamente")
			return
		}
	}

}

func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// get parameters of url
	parameters := mux.Vars(r)
	taskID, err := strconv.Atoi(parameters["id"])

	var updateTask task
	if err != nil {
		fmt.Fprintf(w, "Id invalido")
		return
	}

	// leendo el body
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Body invalido")
	}
	json.Unmarshal(reqBody, &updateTask)
	for _, task := range tasks {
		if task.ID == taskID {
			task.Name = updateTask.Name
			task.Content = updateTask.Content
		}
	}
}

func main() {
	// Configuración inicial de la ruta
	router := mux.NewRouter().StrictSlash(true)

	// Definiendo una ruta
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks)
	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/task/{id}", getTask).Methods("GET")
	router.HandleFunc("/task/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/task/{id}", updateTask).Methods("PUT")
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"content-type"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
	)
	router.Use(cors)
	// Iniciando un servicdor
	log.Fatal(http.ListenAndServe(":5000", router))
}
