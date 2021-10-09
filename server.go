package main

import (
	"fmt"
	c "go-rest-api/api/config"
	"go-rest-api/api/controllers"
	controller "go-rest-api/api/controllers"
	"log"
	"net/http"
	"strings"
)

// var users []User
// var posts []Post

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage is hit")
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		controller.GetUser(w, r)
	case "POST":
		controller.CreateUser(w, r)
	default:
		fmt.Println("error in default")
	}

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch r.Method {
	case "GET":
		parsed := strings.Split(r.URL.Path, "/")
		fmt.Println(parsed, len(parsed), parsed[0], parsed[1])
		if len(parsed) == 4 && parsed[2] == "users" {
			fmt.Println("IN GET OF POST WITH USERID")
			controllers.GetPostsByUserId(w, r, parsed[3])
		} else {
			controller.GetPostById(w, r, parsed[2])
		}
	case "POST":
		controllers.CreatePost(w, r)
	default:

	}
}
func main() {
	c.ConnectDB()
	fmt.Println("Successfully connected to database")
	http.HandleFunc("/users/", UserHandler)
	http.HandleFunc("/posts/", PostHandler)

	//http.HandleFunc("/", Homepage)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
