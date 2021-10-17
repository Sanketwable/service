package routes

import (
	"service/controllers"
)

var covid_cases = []Route {
	{
		Uri: "/covidcases",
		Handler: controllers.GetCovidCases,
	},
}
