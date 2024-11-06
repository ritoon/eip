package model

type Game struct {
	DBField
	Name     string `json:"name"`
	URIImage string `json:"uri_image"`
}
