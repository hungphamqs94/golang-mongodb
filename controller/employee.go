package controller

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"mongodb/models"
	"mongodb/singleton"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := singleton.GetInstance().Collection("employee").Find(ctx, bson.M{})
	if err != nil {
		panic(err.Error())
	}
	var listEmployee []models.Employee
	for cursor.Next(ctx) {
		var employee models.Employee
		cursor.Decode(&employee)
		listEmployee = append(listEmployee, employee)
	}
	json.NewEncoder(w).Encode(listEmployee)
}

func SingleEmployee(w http.ResponseWriter, r *http.Request) {
	id, _ := primitive.ObjectIDFromHex(mux.Vars(r)["id"])
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var employee models.Employee
	errFindOne := singleton.GetInstance().Collection("employee").FindOne(ctx, models.Employee{Id: id}).Decode(&employee)
	if errFindOne != nil {
		panic(errFindOne.Error())
	}

	json.NewEncoder(w).Encode(employee)

}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var employee models.Employee
	json.Unmarshal(reqBody, &employee)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := singleton.GetInstance().Collection("employee").InsertOne(ctx, employee)
	id := result.InsertedID
	employee.Id = id.(primitive.ObjectID)
	json.NewEncoder(w).Encode(employee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var id primitive.ObjectID
	json.Unmarshal(reqBody, &id)
	var employee models.Employee
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := singleton.GetInstance().Collection("employee").UpdateOne(ctx, models.Employee{Id: id}, bson.M{"$set": employee})
	json.NewEncoder(w).Encode(result)
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var id primitive.ObjectID
	json.Unmarshal(reqBody, &id)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, _ := singleton.GetInstance().Collection("employee").DeleteOne(ctx, models.Employee{Id: id})
	json.NewEncoder(w).Encode(result)
}
