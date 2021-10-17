package repository

import (
	"context"
	"fmt"
	"service/controllers"
	"service/database"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func AddData() {
	client := database.Connect()
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)
	data := controllers.GetCovidCasesFromAPI()
	collection := client.Database("Service")
	CovidCollection := collection.Collection("CovidData")
	for _, ele := range data {
		for _, element := range ele.Data {
			CovidResult, err := CovidCollection.InsertOne(ctx, bson.D{
				{Key: "state", Value: ele.State},
				{Key: "district", Value: element.Name},
				{Key: "active", Value: element.Active},
				{Key: "confirmed", Value: element.Confirmed},
				{Key: "deceased", Value: element.Deceased},
				{Key: "recovered", Value: element.Recovered},
			})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(CovidResult)
		}
	}
}
