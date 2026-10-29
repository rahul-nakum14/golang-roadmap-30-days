package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

		"strconv"
	"github.com/gorilla/mux"
)

type Course struct{
	CourseId string `json:"courseId"`
	CourseName string `json:"courseName"`
	CoursePrice string `json:"coursePrice"`
	Author *Author	`json:"author"`
}

type Author struct{
	Fullname string `json:"fullname"`
}


var courses []Course


func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/health", HealthHandler).Methods("GET")

	log.Println("Server starting on port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func (c *Course)IsEmpty()bool{
	return c.CourseName == "" 
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home endpoint called")
	w.Write([]byte("Hello World"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is running"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-type","application/json")

		params := mux.Vars(r)
	
		for _,course :=range courses {
			if course.CourseId == params["id"]{
				json.NewEncoder(w).Encode(course)
				return
			}
		}

		json.NewEncoder(w).Encode("No course found")
				return

}

func addCourse(w http.ResponseWriter, r *http.Request){
			w.Header().Set("Content-type","application/json")

			if r.Body == nil{
				json.NewEncoder(w).Encode("please send some data")
			}


			var course Course

			_ = json.NewDecoder(r.Body).Decode(&course)

			if course.IsEmpty(){
								json.NewEncoder(w).Encode("please send some data")
					return
			}


			course.CourseId = strconv.Itoa(rand.Intn(100))

			courses = append(courses, course)

			json.NewEncoder(w).Encode(course)

			return 

		}