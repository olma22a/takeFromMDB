package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type ExchangeRate struct {
	Koken string             `json:"static_koken"`
	Time  time.Time          `json:"time"`
	Base  string             `json:"base_code"`
	Rates map[string]float64 `json:"conversion_rates"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
	}()

	collection := client.Database("exchange").Collection("rub")
	timeLayout := "2006-01-02T15:04:05.999-07:00"
	timeStr := "2024-03-19T22:36:27.628+00:00"
	parsedTime, err := time.Parse(timeLayout, timeStr)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{
		"time1": parsedTime,
	}
	var result ExchangeRate
	if err := collection.FindOne(context.Background(), filter).Decode(&result); err != nil {
		log.Fatal(err)
	}

	staticKoken := "randomToken"
	result.Koken = "randomToken"
	if result.Koken != staticKoken {
		fmt.Println("Введите верный токен")
		log.Fatal(err)
	}

	fmt.Println(result)
}
