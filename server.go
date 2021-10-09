package main

import (
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Id       int    `json:","`
	Name     string `json:", omitempty"`
	Email    string `json:", omitempty"`
	Password string `json:", omitempty"` // to ignore the field "-"
}

type Post struct {
	Id              int    `json:", omitempty"`
	Caption         string `json:",omitempty"`
	ImageUrl        string `json:",omitempty"`
	PostedTimestamp string `json:", omitempty"`
	CreatedBy       int    `json:", omitempty`
}

var users []User
var posts []Post

func Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage is hit")
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	fmt.Println(r.Method)
	switch r.Method {
	case "GET":
		fmt.Println(r.URL.Path)
		s := strings.Split(r.URL.Path, "/")[2]
		userId, err := strconv.Atoi(s)
		if err == nil {
			for i, v := range users {
				fmt.Println(v.Id, userId)
				if v.Id == userId {
					fmt.Println("sending it *******************")
					json.NewEncoder(w).Encode(users[i])
					return
				}
			}
			// no user with that id
		}
		// what happens if an error occurs

	case "POST":
		fmt.Println("IN POST")
		var user User
		h := sha512.New()

		var err = json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Fprintf(w, "error")
			return
		}
		user.Id = len(users)
		io.WriteString(h, user.Password)
		a := h.Sum(nil)
		user.Password = string(a)
		users = append(users, user)
		// send a response stating it is successfull

		fmt.Println("done")
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
			userId, err := strconv.Atoi(parsed[3])
			var userPosts []Post
			if err == nil {
				// valid value for postId
				for _, v := range posts {
					if v.CreatedBy == userId {
						userPosts = append(userPosts, v)
					}
				}
				json.NewEncoder(w).Encode(userPosts)
				return
				// no post with id!!
			}
		} else {
			fmt.Println("IN GET OF POST WITH POSTID")
			postId, err := strconv.Atoi(parsed[2])
			if err == nil {
				// valid value for postId
				for _, v := range posts {
					if v.Id == postId {
						json.NewEncoder(w).Encode(v)
						return
					}
				}
				// no post with id!!
			}
		}

	case "POST":
		fmt.Println("IN POST OF POST")
		var post Post
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			fmt.Println("error")
			return
		}
		// timestamp addition
		post.Id = len(posts)
		post.PostedTimestamp = time.Now().UTC().String() // gets the current timestamp of utc time zone
		posts = append(posts, post)
	default:

	}
}
func main() {
	http.HandleFunc("/users/", CreateUser)
	http.HandleFunc("/posts/", PostHandler)

	//http.HandleFunc("/", Homepage)
	log.Fatal(http.ListenAndServe(":3000", nil))
}
