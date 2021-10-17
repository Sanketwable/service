package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"service/models"
)

func GetCovidCasesFromAPI() []models.CovidData {
	url := "https://data.covid19india.org/state_district_wise.json"
	var data map[string]interface{}
	r, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal(body, &data)
	covidDatas := []models.CovidData{}

	for k := range data {
		m := models.CovidData{}
		m.State = k
		covidDatas = append(covidDatas, m)
	}
	for i, ele := range covidDatas {
		by, _ := json.Marshal(data[ele.State])
		var data1 map[string]interface{}
		json.Unmarshal(by, &data1)
		var districts map[string]interface{}
		by, _ = json.Marshal(data1["districtData"])
		json.Unmarshal(by, &districts)
		dis := []models.DistrictData{}
		for k := range districts {
			lp := models.DistrictData{}
			lp.Name = k
			
			var disData map[string]int
			by, _ = json.Marshal(districts[k])
			json.Unmarshal(by, &disData)
			lp.Active = disData["active"]
			lp.Confirmed = disData["confirmed"]
			lp.Deceased = disData["deceased"]
			lp.Recovered = disData["recovered"]
			dis = append(dis, lp)
		}
		covidDatas[i].Data = dis
	}
	return covidDatas
}
