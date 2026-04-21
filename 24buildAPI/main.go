package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for the Course -File

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Creating a fake database
var courses []Course

// middleware , helper- File
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""

}

func main() {

}

// controllers - File

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>WELCOME TO API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourese(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Courese")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request(user)
	params := mux.Vars(r)
	//loop through courses and match the idand return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course Found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Courese")
	w.Header().Set("Content-Type", "application/json")
	//what if:body is Empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send Some Data")
	}
	//what if data={}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data indide the Json")
		return
	}

	//generate new id
	//append new course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func uppdateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Onne Course")
	w.Header().Set("Content-Type", "application/json")

	//first grab id from request
	params := mux.Vars(r)

	// loop, id,remove,add with my id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
