package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	return c.CourseName == ""
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
	muxRoute.HandleFunc("/course/{id}", updateCourse).Methods("put")
	muxRoute.HandleFunc("/course/{id}", deleteACourse).Methods("delete")
	muxRoute.HandleFunc("/courses/all", deleteAllCourse).Methods("delete")
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
	w.Write([]byte("{course_not_found:" + params["id"] + "}"))
	// json.NewEncoder(w).Encode(interface{"ky":"va"})
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("requested create course")

	var userCourse Course
	json.NewDecoder(r.Body).Decode(&userCourse)
	fmt.Println(userCourse)

	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		fmt.Println("body is empty")
		w.Write([]byte("Error, Invalid Data"))
		return
	} else if userCourse.IsEmpty() {
		fmt.Println("data is invalid")
		w.WriteHeader(400)
		w.Write([]byte("Error, Invalid Data"))
		return
	} else {
		for _, courseData := range courses {
			if userCourse.CourseName == courseData.CourseName {
				fmt.Println("Duplicate Course")
				w.Write([]byte("Course Exist"))
				return
			}
		}
	}
	userCourse.CourseId = strconv.Itoa(len(courses) + 1)
	courses = append(courses, userCourse)
	w.Write([]byte("success"))
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request update course")
	params := mux.Vars(r)
	fmt.Println(params)
	fmt.Println(params["id"])

	var userCourseToEdit Course
	decodedJsonBody := json.NewDecoder(r.Body)
	err := decodedJsonBody.Decode(&userCourseToEdit)
	if err != nil {
		fmt.Println("user json data validation error")
		w.Write([]byte("request error"))
		return
	}

	for sliceIndex, courseData := range courses {
		if courseData.CourseId == params["id"] {
			fmt.Println("Course Id Matched")
			courses = append(courses[:sliceIndex], courses[sliceIndex+1:]...)
			userCourseToEdit.CourseId = params["id"]
			courses = append(courses, userCourseToEdit)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(userCourseToEdit)
			return
		}
	}
	fmt.Println("can't find element for update")
	w.Write([]byte("request error"))
}

func deleteACourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request delete course")
	parms := mux.Vars(r)
	for sliceIndex, courseData := range courses {
		fmt.Println(sliceIndex)
		fmt.Println(courseData)
		if parms["id"] == courseData.CourseId {
			fmt.Println("Course Id Matched")
			courses = append(courses[:sliceIndex], courses[sliceIndex+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(courseData)
			return
		}
	}
	fmt.Println("can't find element for delete")
	w.Write([]byte("request error"))
}

func deleteAllCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request Delete All Course")
	if len(courses) > 0 {
		courses = []Course{}
		json.NewEncoder(w).Encode("success")
		return
	}
	json.NewEncoder(w).Encode("no data to delete")
}

func fourNotFour(w http.ResponseWriter, r *http.Request) {
	fmt.Println("route not found")
	w.WriteHeader(404)
	w.Write([]byte("Not Found"))
}
