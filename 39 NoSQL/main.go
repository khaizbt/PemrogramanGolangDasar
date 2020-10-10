package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"

	//"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var ctx = context.Background()

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func main() {
	//insert()
	find()
	fmt.Println("--===============--")
	update()
	fmt.Println("--===============--")
	delete()
}

func connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)

	if err != nil {
		return nil, err
	}

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database("belajar"), nil
}

func insert() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Khaiz", 2})
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(ctx, student{"Tammam", 1})
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Insert Sukses")
}

func find() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Find(ctx, bson.M{"name": "Tammam"})
	if err != nil {
		log.Fatal(err.Error())
	}

	defer csr.Close(ctx)

	result := make([]student, 0)

	for csr.Next(ctx) {
		var row student
		err := csr.Decode(&row)

		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)

	}

	if len(result) > 0 {
		fmt.Println("Name : ", result[0].Name)
		fmt.Println("Grade", result[0].Grade)
	}
}

func update() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())

	}

	var selector = bson.M{"name": "Tammam"}
	var changes = student{"Khaiz Badaru Tammam", 23}
	_, err = db.Collection("student").UpdateOne(ctx, selector, bson.M{"$set": changes})

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Update Sukses")
}

func delete() {
	db, err := connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": "Khaiz"}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("Delete Data Sukses")
}
