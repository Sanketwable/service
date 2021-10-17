package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"service/cache"
	"service/config"
	"service/database"
	"service/models"
	"time"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func GetCovidCases(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	location := models.Location{}
	json.Unmarshal(body, &location)

	cache := cache.NewRedisCache(config.REDIS_HOST, config.DB_INDEX, (time.Duration(config.EXP)*time.Second))

	var result *models.Response =  cache.Get(&location)
	if result.StateName == "" {
		district, _, _, _ := GetStateDistrict(location.Latitude, location.Longitude)
		if district == "" {
			return c.JSON(http.StatusAccepted, "Enter Valid coordinates")
		}	
		*result, err = GetData(district)
		if err != nil {
			return c.JSON(http.StatusBadRequest, "Enter Valid coordinates")
		}
		cache.Set(&location, result)
	}
	fmt.Println(*result)
	return c.JSON(http.StatusBadRequest, *result)
}
func GetData(city string) (models.Response, error) {
	result := models.Response{}
	fmt.Println("City = ", city)
	filter := bson.D{{Key: "district", Value: city}}
	client := database.Connect()
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)
	collection := client.Database("Service")
	CovidCollection := collection.Collection("CovidData")
	err := CovidCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}
	return result, nil
}
