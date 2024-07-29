package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseID   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	Name    string `json:"fullname"`
	Website string `json:"site"`
}

// db
var courses []Course

// middleware
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" || c.CourseName == ""
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/:id", fetchOneCourse).Methods("GET")
	router.HandleFunc("/", serverHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// controllers - file
// home route

func serverHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>welcome</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi in all")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func fetchOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hi from one")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	fmt.Println(params)
}
