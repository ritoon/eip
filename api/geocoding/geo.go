package geocoding

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/goccy/go-json"
)

type GeoResponse struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

func New(ctx context.Context, query string) (lat float64, lng float64, err error) {
	r, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/search?q="+query, nil)
	if err != nil {
		return 0, 0, err
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Println("geocoding: http.DefaultClient.Do error:", err)
		if ctx.Err() != nil {
			log.Println("geocoding: ctx.Err() error:", err)
			return 0, 0, newErrorTimeout("context", ctx.Err())
		}
		return 0, 0, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	var payload GeoResponse
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return 0, 0, err
	}

	return payload.Lat, payload.Lng, nil
}
