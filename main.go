package main

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoInstance struct {
	Client
	Db
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb+srv://lumi:<password>@cluster0.yl6aswd.mongodb.net/?retryWrites=true&w=majority" + dbName

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    float64
}

func Connect() error {
	client, err := mongo.NewClient()
	context.WithTimeout(context.Background(), 30*time.Second)
}
