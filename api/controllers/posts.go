package controllers

import (
	"context"
	"encoding/json"
	c "go-rest-api/api/config"
	m "go-rest-api/api/models"
	u "go-rest-api/api/utils"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	var post m.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		u.SendError(w, err, http.StatusInternalServerError)
	}
	post.PostedTimestamp = time.Now().UTC().String() // gets the current timestamp of utc time zone
	collection := c.Client.Database("golangrestapi").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	post.Id = primitive.NewObjectID()
	result, err := collection.InsertOne(ctx, post)
	if err != nil {
		u.SendError(w, err, http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetPostsByUserId(w http.ResponseWriter, r *http.Request, userId string) {
	collection := c.Client.Database("golangrestapi").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	userObjectId, e := primitive.ObjectIDFromHex(userId)
	if e != nil {
		u.SendError(w, e, http.StatusInternalServerError)
		return
	}
	cursor, errr := collection.Find(ctx, bson.M{"createdBy": userObjectId})
	if errr != nil {
		u.SendError(w, errr, http.StatusInternalServerError)
		return
	}

	// usercollection := c.Client.Database("golangrestapi").Collection("users")
	// var user m.User

	// err := usercollection.FindOne(ctx, bson.M{"_id": userObjectId}).Decode(&user)
	// if err != nil {
	// 	u.SendError(w, err, http.StatusInternalServerError)
	// 	defer cursor.Close(ctx)
	// 	return
	// }
	var posts []m.Post
	for cursor.Next(ctx) {
		var post m.Post
		err := cursor.Decode(&post)
		if err != nil {
			u.SendError(w, err, http.StatusInternalServerError)
			defer cursor.Close(ctx)
			return
		}
		posts = append(posts, post)
	}

	// result := make(map[string]string)
	// result["creatorEmail"] = user.Email
	// result["creatorName"] = user.Name
	// result["creatorId"] = user.Id.String()
	// type Data struct {
	// 	creator map[string]string
	// 	posts   []m.Post
	// // }
	// data, err := json.Marshal(Data{result, posts})
	// if err != nil {
	// 	u.SendError(w, err, http.StatusInternalServerError)
	// 	defer cursor.Close(ctx)
	// 	return
	// }
	// w.Write(data)
	// fmt.Println(data)
	defer cursor.Close(ctx)
	json.NewEncoder(w).Encode(posts)
}

func GetPostById(w http.ResponseWriter, r *http.Request, postIdString string) {
	var post m.Post
	collection := c.Client.Database("golangrestapi").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	postId, er := primitive.ObjectIDFromHex(postIdString)
	if er != nil {
		u.SendError(w, er, http.StatusInternalServerError)
		return
	}
	errr := collection.FindOne(ctx, bson.M{"_id": postId}).Decode(&post)
	if errr != nil {
		u.SendError(w, errr, http.StatusInternalServerError)
		return
	}
	usercollection := c.Client.Database("golangrestapi").Collection("users")
	var user m.User
	err := usercollection.FindOne(ctx, bson.M{"_id": post.CreatedBy}).Decode(&user)
	if err != nil {
		u.SendError(w, errr, http.StatusInternalServerError)
		return
	}
	result := make(map[string]string)
	result["caption"] = post.Caption
	result["_id"] = post.Id.String()
	result["imageUrl"] = post.ImageUrl
	result["postedTimestamp"] = post.PostedTimestamp
	result["creatorEmail"] = user.Email
	result["creatorName"] = user.Name
	result["creatorId"] = user.Id.String()
	json.NewEncoder(w).Encode(result)
}
