package routes

import (
	"service/controllers"
)

var fetch_cases = []Route{
	{
		Uri:     "/fetchcases",
		Handler: controllers.FetchCases,
	},
}
