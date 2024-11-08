package scrap

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"io"
	"log"
	"net/http"

	"github.com/ritoon/eip/api/model"
)

type Item struct {
	Name      string `xml:"name"`
	Thumbnail string `xml:"thumbnail"`
}

type Items struct {
	XMLName xml.Name `xml:"items"`
	Items   []Item   `xml:"item"`
}

func New() error {
	resp, err := http.Get("https://boardgamegeek.com/xmlapi/collection/wolfgarou")
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var items Items
	err = xml.Unmarshal(body, &items)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items.Items {
		log.Println(item.Name)
		var game model.Game
		game.Name = item.Name
		game.URIImage = item.Thumbnail
		data, err := json.Marshal(game)
		if err != nil {
			log.Println(err)
		}
		body := bytes.NewReader(data)
		r, err := http.NewRequest("POST", "http://localhost:8888/api/v1/games", body)
		if err != nil {
			log.Println(err)
		}
		r.Header.Set("Content-Type", "application/json")
		r.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkX3VzZXIiOiJ1c3ItMm9aUXZPeWVkbXNLNzl3SHFZdWRyOVEyVmdJIiwiYWNjZXNzX2xldmVsIjoiIiwiZW1haWwiOiJ0b3RvIiwiZXhwIjoxNzMxMDc4NDcxfQ.oJzWfeidd5jE6sCUlEWJlVthCbwnxwHcdJTkLAVPGjI")

		resp, err := http.DefaultClient.Do(r)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			log.Println("error")
		}

	}
	return nil
}
