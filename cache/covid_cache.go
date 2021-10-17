package cache

import "service/models"

type CovidCache interface {
	Set(key *models.Location, value *models.Response)
	Get(key *models.Location) *models.Response
}