package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"service/models"
)

func GetStateDistrict(lat float64, lon float64) (string, string, string, string) {
	key := "AIzaSyCd9bdKJ1FOZrydzwm9sX-KSyeBdFGQR-4"
	url := "https://maps.googleapis.com/maps/api/geocode/json?latlng="+ fmt.Sprintf("%f,%f&key=", lat, lon) + key
	georesponse := models.GeoResponse{}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(body, &georesponse)
	if err != nil {
		fmt.Println(err)
	}
	if len(georesponse.Result) == 0 {
		return "", "", "", ""
	}
	var (
		district  = ""
		state     = ""
		stateAncr = ""
		country   = ""
	)
	for _, element := range georesponse.Result[0].AddressComponent {
		if element.Types[0] == "administrative_area_level_2" {
			district = element.LongName
		}
		if element.Types[0] == "administrative_area_level_1" {
			state = element.LongName
			stateAncr = element.ShortName
		}
		if element.Types[0] == "country" {
			country = element.LongName
		}

	}

	return district, state, stateAncr, country

}
