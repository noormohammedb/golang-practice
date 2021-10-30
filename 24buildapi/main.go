package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice float64 `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// middleware,helper - file

func (c *Course) IsEmpty() bool {
	return c.CourseId == ""
}

// DataBase
var courses []Course

func main() {
	fmt.Println("Golang API Server")
	muxRoute := mux.NewRouter()
	muxRoute.HandleFunc("/", serveHome).Methods("get")
	muxRoute.HandleFunc("/allcourse", getAllCourses).Methods("get")
	muxRoute.HandleFunc("/course/{id}", getOneCourse).Methods("get")
	muxRoute.HandleFunc("/course", createCourse).Methods("post")
	muxRoute.HandleFunc("/course", updateCourse).Methods("put")
	muxRoute.HandleFunc("/course", deleteACourse).Methods("delete")
	muxRoute.HandleFunc("/{*}", fourNotFour)

	// seeding data
	courses = append(courses, Course{CourseId: "1", CourseName: "courseName-1", CoursePrice: 333, Author: &Author{FullName: "author-1", Website: "www.author1.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "courseName-2", CoursePrice: 444, Author: &Author{FullName: "author-2", Website: "www.author2.com"}})

	log.Fatal(http.ListenAndServe(":8080", muxRoute))
}

// controllers - file

func serveHome(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested /")
	w.Write([]byte("<h1>home</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested ")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Requested One Course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Println(params["id"])
	for _, value := range courses {
		if value.CourseId == params["id"] {
			json.NewEncoder(w).Encode(value)
			return
		}
	}
	w.Write([]byte("{course_not_found: true}"))
	// json.NewEncoder(w).Encode(interface{"ky":"va"})
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("requested create course")
	// check if the body is empty
	w.Header().Set("Content-Type", "application/json")

	var userCourse Course
	json.NewDecoder(r.Body).Decode(&userCourse)
	fmt.Println(userCourse)

	if userCourse.IsEmpty() {
		fmt.Println("data is invalid")
		w.WriteHeader(400)
		w.Write([]byte("Error, Invalid Data"))
		return
	}
	courses = append(courses, userCourse)
	w.Write([]byte("success"))
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request update course need to write logic")
	w.Write([]byte("contact administrator"))
}

func deleteACourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request delete course need to wire logic")
	w.Write([]byte("contact administrator"))
}

func fourNotFour(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route not found")
	w.WriteHeader(404)
	w.Write([]byte("Not Found"))
}
