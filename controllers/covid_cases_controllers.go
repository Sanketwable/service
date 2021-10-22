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

type fetchResponse struct{
	Message string `json:"message"`
	Error string `json:"error"`

}

func GetCovidCases(c echo.Context) error {
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return err
	}

	location := models.Location{}
	json.Unmarshal(body, &location)

	cache := cache.NewRedisCache(config.REDIS_HOST, config.DB_INDEX, (time.Duration(config.EXP) * time.Second))

	fR := fetchResponse{}
	fR.Error = "Enter Valid coordinates within region"
	fR.Message = "failed to provide covid info"

	var result *models.Response = cache.Get(&location)
	if result.StateName == "" {
		district, _, _, _ := GetStateDistrict(location.Latitude, location.Longitude)
		if district == "" {
			
			return c.JSON(http.StatusBadRequest, fR)
		}
		*result, err = GetData(district)
		if err != nil {
			return c.JSON(http.StatusBadRequest, fR)
		}
		cache.Set(&location, result)
	}
	fmt.Println(*result)
	return c.JSON(http.StatusAccepted, *result)
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

func FetchCases(c echo.Context) error {
	client := database.Connect()
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)
	
	data := GetCovidCasesFromAPI()
	if data == nil {
		fR := fetchResponse{}
		fR.Message = "unable to fetch data from given api"
		fR.Error = "No Data found"
		return c.JSON(http.StatusCreated, fR)
	}

	storeRes := []interface{}{}
	for _, ele := range data {
		for _, e := range ele.Data {
			m := models.Response{}
			m.StateName = ele.State
			m.ActiveNo = int32(e.Active)
			m.ConfirmedNo = int32(e.Confirmed)
			m.DeceasedNo = int32(e.Deceased)
			m.District = e.Name
			m.RecoveredNo = int32(e.Recovered)
			storeRes = append(storeRes, m)
		}

	}
	collection := client.Database("Service")
	if err := collection.Drop(ctx); err != nil {
		fmt.Println("error droping db")
		fmt.Println(err)
	}

	CovidCollection := collection.Collection("CovidData")
	CovidResult, err := CovidCollection.InsertMany(ctx, storeRes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(CovidResult)


	// for _, ele := range data {
	// 	for _, element := range ele.Data {
	// 		CovidResult, err := CovidCollection.InsertOne(ctx, bson.M{
	// 			"state", Value: ele.State,
	// 			"district", Value: element.Name,
	// 			"active", Value: element.Active,
	// 			"confirmed", Value: element.Confirmed,
	// 			"deceased", Value: element.Deceased,
	// 			"recovered", Value: element.Recovered,
	// 		})
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		fmt.Println(CovidResult)
	// 	}
	// }
	fR := fetchResponse{}
	fR.Message = "data successfully fetched from api"
	fR.Error = "nil"
	return c.JSON(http.StatusCreated, fR)
}
