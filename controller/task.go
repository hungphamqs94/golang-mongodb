package controller

import (
	"context"
	"encoding/json"
	"mongodb/models"
	"mongodb/singleton"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := singleton.GetInstance().Collection("task").Find(ctx, bson.M{})

	if err != nil {
		panic(err.Error())
	}

	var listTask []models.Task

	for cursor.Next(ctx) {
		var task models.Task
		cursor.Decode(&task)
		listTask = append(listTask, task)
	}

	json.NewEncoder(w).Encode(listTask)

}

func SingleTask(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	var task models.Task
	err := singleton.GetInstance().Collection("task").FindOne(ctx, models.Task{Id: id}).Decode(task)
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(task)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var task models.Task
	cursor, err := singleton.GetInstance().Collection("task").InsertOne(ctx, task)
	if err != nil {
		panic(err.Error())
	}
	id := cursor.InsertedID
	task.Id = id.(primitive.ObjectID)
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	var task models.Task
	updateTask, err := singleton.GetInstance().Collection("task").UpdateOne(ctx, models.Task{Id: id}, bson.M{"$set": task})
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(updateTask)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	deleteTask, err := singleton.GetInstance().Collection("task").DeleteOne(ctx, models.Task{Id: id})
	if err != nil {
		panic(err.Error())
	}
	json.NewEncoder(w).Encode(deleteTask)
}
