package models

type CovidData struct {
	State string         `bson:"state" json:"state"`
	Data  []DistrictData `bson:"district" json:"district"`
}
type DistrictData struct {
	Name      string `bson:"district_name" json:"district_name"`
	Active    int    `bson:"active" json:"active"`
	Confirmed int    `bson:"confirmed" json:"confirmed"`
	Deceased  int    `bson:"deceased" json:"deceased"`
	Recovered int    `bson:"recovered" json:"recovered"`
}

type Response struct {
	StateName   string `bson:"state"`
	District    string `bson:"district"`
	ActiveNo    int32  `bson:"active"`
	ConfirmedNo int32  `bson:"confirmed"`
	DeceasedNo  int32  `bson:"deceased"`
	RecoveredNo int32  `bson:"recovered"`
}
