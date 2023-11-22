package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/health", healthcheck)
	mux.HandleFunc("/v1/tasks", tasks)
	mux.HandleFunc("/v1/tasks/", task)
	http.ListenAndServe(":3000", mux)
}

type Task struct {
	Id      int64     `json:"id"`
	Name    string    `json:"name"`
	Created time.Time `json:"created_time"`
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	status := map[string]interface{}{
		"status":      "Active",
		"environment": "Dev",
		"Version":     "1.0.0",
	}

	js, err := json.Marshal(status)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application-json")
	w.Write(js)
}

func tasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		all_tasks(w, r)
	case http.MethodPost:
		create_task(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func all_tasks(w http.ResponseWriter, r *http.Request) {
	task := Task{
		Id:      1,
		Name:    "Test",
		Created: time.Now(),
	}

	js, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "StatusInternalServerError", http.StatusInternalServerError)
	}

	js = append(js, '\t')

	w.Header().Set("Content-Type", "application-json")
	w.Write(js)
}

func create_task(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name         string    `json:"name"`
		Created time.Time `json:"created_time"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Fprintln(w, "not decoded")
	}

	js, err := json.Marshal(input)
	if err != nil {
		fmt.Fprintln(w, "error 2")
	}

	w.Header().Set("Content-Type", "application-json")
	w.Write(js)
}

func task(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		get_task(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func get_task(w http.ResponseWriter, r *http.Request){
    id := r.URL.Path[len("/v1/tasks/"):]
    idInt, err := strconv.ParseInt(id,10,64)
    if err != nil || idInt < 1{
        fmt.Fprintln(w, "invalid data id")
        return
    }

    data := Task{
        Id: idInt,
        Name: "Simple task",
        Created: time.Now(),
    }

    js, err := json.Marshal(data)
    if err != nil{
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    }

    js = append(js, '\n')

    w.Header().Set("Content-Type", "application-json")
    w.Write(js)
}
