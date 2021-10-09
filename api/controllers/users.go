package controllers

import (
	"context"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	c "go-rest-api/api/config"
	m "go-rest-api/api/models"
	u "go-rest-api/api/utils"
	"io"
	"net/http"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("IN POST")
	var user m.User
	h := sha512.New()

	json.NewDecoder(r.Body).Decode(&user)
	io.WriteString(h, user.Password)
	user.Password = string(h.Sum(nil))

	collection := c.Client.Database("golangrestapi").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	user.Id = primitive.NewObjectID()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		u.SendError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	s := strings.Split(r.URL.Path, "/")[2]

	fmt.Println("IN GET USERS")
	var user m.User
	collection := c.Client.Database("golangrestapi").Collection("users")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	userId, er := primitive.ObjectIDFromHex(s)
	if er != nil {
		u.SendError(w, er, http.StatusInternalServerError)
		return
	}
	errr := collection.FindOne(ctx, bson.M{"_id": userId}).Decode(&user)
	if errr != nil {
		u.SendError(w, errr, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}
