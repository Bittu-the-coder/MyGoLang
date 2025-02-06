package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"id,omitempty"`
	CourseName  string  `json:"name"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author,omitempty"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses = []Course{}

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {

	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "1", CourseName: "Python", CoursePrice: 499, Author: &Author{Fullname: "John Doe", Website: "johndoe.com"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Java", CoursePrice: 399, Author: &Author{Fullname: "John Doe", Website: "johndoe.com"}})

	// routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/api/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/api/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/api/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/api/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/api/courses/{id}", deleteOneCourse).Methods("DELETE")

	//listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controller -file

//serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to api in golang</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through our array of courses and find the matching id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with id " + params["id"])
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// What if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what if - {}
	var course Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//generate id , string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through our array of courses and find the matching id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			// decode the json body
			var updatedCourse Course
			_ = json.NewDecoder(r.Body).Decode(&updatedCourse)
			course.CourseId = params["id"]
			courses = append(courses, course)
			// courses[index].CourseName = updatedCourse.CourseName
			json.NewEncoder(w).Encode(courses[index])
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with id " + params["id"])
	return

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through our array of courses and find the matching id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			break
		}
	}
}
