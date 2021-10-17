package models

type GeoResponse struct {
	Result []Result `json:"results"`
}
type Result struct {
	AddressComponent []AddressComponent `json:"address_components"` 
}
type AddressComponent struct {
	LongName string `json:"long_name"`
	ShortName string `json:"short_name"`
	Types []string `json:"types"`
}