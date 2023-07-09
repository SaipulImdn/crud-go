package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name  string
	Email string
}

func main() {
	// Konfigurasi koneksi MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Mendapatkan referensi ke database dan koleksi
	collection := client.Database("mydb").Collection("persons")

	// Contoh operasi Create
	person := Person{Name: "Budi", Email: "budi@gmail.com"}
	insertResult, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("ID yang baru disisipkan:", insertResult.InsertedID)

	// Contoh operasi Read
	filter := bson.M{"name": "Budi"}
	var result Person
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Data yang ditemukan:", result)

	// Contoh operasi Update
	update := bson.M{"$set": bson.M{"email": "newemail@example.com"}}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Jumlah dokumen yang diupdate:", updateResult.ModifiedCount)

	// Contoh operasi Delete
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Jumlah dokumen yang dihapus:", deleteResult.DeletedCount)
}
