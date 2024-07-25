package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	Id         string `json:"id"`
	CourseName string `json:"courseName"`
}

var courses []Course = []Course{
	{"1", "Python"},
	{"2", "Golang"},
	{"3", "Machine Learning"},
	{"4", "LLMs"},
}

func (c Course) IsEmpty() bool {
	return c.Id == "" && c.CourseName == ""
}

const port = ":8500"

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World Golang!")
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
	return
}

func getCourseById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getCourseById")
	params := mux.Vars(r)
	id := params["id"]
	println(id)
	for _, course := range courses {
		if id == course.Id {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	fmt.Fprintf(w, "Course not found")
}

func handleRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", helloWorld).Methods("GET")
	r.HandleFunc("/getCourseById/{id}", getCourseById).Methods("GET")
	r.HandleFunc("/getAllCourses", getAllCourses).Methods("GET")
	http.ListenAndServe(port, r)
}

func main() {
	fmt.Println("Hello, World from Golang!")
	handleRequests()
}
