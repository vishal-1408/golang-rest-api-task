package main

import (
	"fmt"
	c "go-rest-api/api/config"
	"go-rest-api/api/controllers"
	controller "go-rest-api/api/controllers"
	"log"
	"net/http"
	"os"
	"strings"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch r.Method {
	case "GET":
		controller.GetUser(w, r)
	case "POST":
		controller.CreateUser(w, r)
	default:
		fmt.Println("invalid http method")
		os.Exit(1)
	}

}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	switch r.Method {
	case "GET":
		parsed := strings.Split(r.URL.Path, "/")
		if len(parsed) == 4 && parsed[2] == "users" {
			controllers.GetPostsByUserId(w, r, parsed[3])
		} else {
			controller.GetPostById(w, r, parsed[2])
		}
	case "POST":
		controllers.CreatePost(w, r)
	default:
		fmt.Println("invalid http method")
		os.Exit(1)
	}
}
func main() {
	c.ConnectDB()
	fmt.Println("Successfully connected to database")
	http.HandleFunc("/users/", UserHandler)
	http.HandleFunc("/posts/", PostHandler)

	log.Fatal(http.ListenAndServe(":3000", nil))
}
