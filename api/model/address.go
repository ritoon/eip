package model

type Address struct {
	DBField
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
	Lat    string `json:"lat"`
	Lng    string `json:"lng"`
}
