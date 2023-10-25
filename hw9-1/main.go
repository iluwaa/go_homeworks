package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	zlog "github.com/rs/zerolog/log"
)

var TaskList = make(Tasks)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", root)
	r.HandleFunc("/tasks", createTask).Methods("POST")
	r.HandleFunc("/tasks", listTasks).Methods("GET")
	r.HandleFunc("/tasks/{date}", getTasks).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	zlog.Err(srv.ListenAndServe())
}

func root(w http.ResponseWriter, r *http.Request) {
	zlog.Print("Handled " + r.URL.Path)

	w.Write([]byte("Hello it is a task server!\n"))
	w.Write([]byte("To create task send POST request to /tasks with body: TODO body\n"))
	w.Write([]byte("To list all task for any day send GET to /tasks/{YYYY-MM-DD}\n"))
}

func createTask(w http.ResponseWriter, r *http.Request) {
	zlog.Print("Handled " + r.URL.Path)
	resp := make(map[string]string)
	body, err := io.ReadAll(r.Body)

	if err != nil {
		zlog.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var newTask Task
	err = json.Unmarshal(body, &newTask)
	if err != nil {
		zlog.Print(err)
		http.Error(w, "Failed to unmarshal json", http.StatusInternalServerError)
		return
	}

	layout := "2006-01-02"
	_, err = time.Parse(layout, newTask.Date)
	if err != nil {
		zlog.Print("Invalid date format:", err)
		resp["error"] = fmt.Sprintf("Invalid date format: %s", err.Error())
		jsonResp, _ := json.Marshal(resp)

		http.Error(w, string(jsonResp), http.StatusBadRequest)
		return
	}

	if _, ok := TaskList[newTask.Date]; !ok {
		TaskList[newTask.Date] = []Task{}
	}

	newTask.Id = len(TaskList[newTask.Date])
	TaskList[newTask.Date] = append(TaskList[newTask.Date], newTask)

	resp["created"] = "true"
	jsonResp, _ := json.Marshal(resp)
	http.Error(w, string(jsonResp), http.StatusOK)

}

func listTasks(w http.ResponseWriter, r *http.Request) {
	zlog.Print("Handled " + r.URL.Path)
	jsonResp, _ := json.Marshal(&TaskList)
	http.Error(w, string(jsonResp), http.StatusOK)
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	zlog.Print("Handled " + r.URL.Path)
	vars := mux.Vars(r)

	jsonResp, _ := json.Marshal(TaskList[vars["date"]])
	http.Error(w, string(jsonResp), http.StatusOK)
}

type Tasks map[string][]Task

type Task struct {
	Id          int
	Name        string `json:"name"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
