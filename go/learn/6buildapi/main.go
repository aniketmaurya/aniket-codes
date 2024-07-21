package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId   string `json:"courseId"`
	CourseName string `json:"courseName"`
}

var courses []Course = []Course{
	{"1", "Python"},
	{"2", "Golang"},
	{"3", "Machine Learning"},
	{"4", "LLMs"},
}

func (c Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
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
	vars := mux.Vars(r)
	courseId := vars["courseId"]
	println(courseId)
	for _, course := range courses {
		if courseId == course.CourseId {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	fmt.Fprintf(w, "Course not found")
}

func handleRequests() {
	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/getCourseById", getCourseById)
	http.HandleFunc("/getAllCourses", getAllCourses)
	http.ListenAndServe(port, nil)
}

func main() {
	fmt.Println("Hello, World from Golang!")
	handleRequests()
}
