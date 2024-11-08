package geocoding

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"

	"github.com/goccy/go-json"
)

// GeoResponse response from geocoding service
type GeoResponse struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

var (
	ErrNilContext = errors.New("geocode: need context not nil")
	ErrEmptyQuery = errors.New("geocode: need query not empty")
)

// New geocoding service to get lat and lng from query
// example: New(ctx, "jakarta")
func New(ctx context.Context, query string) (lat float64, lng float64, err error) {
	if ctx == nil {
		return 0, 0, ErrNilContext
	}
	if query == "" {
		return 0, 0, ErrEmptyQuery
	}
	urlQuery := url.PathEscape(query)
	// create new request with context
	r, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/search?q="+urlQuery, nil)
	if err != nil {
		return 0, 0, err
	}
	// use default http client to send request
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		if ctx.Err() != nil {
			return 0, 0, newErrorTimeout("context", ctx.Err())
		}
		return 0, 0, err
	}
	// close response body after function ends
	defer resp.Body.Close()

	// read response body to data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, err
	}

	// unmarshal data to GeoResponse
	var payload GeoResponse
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return 0, 0, err
	}
	// return lat and lng from payload
	return payload.Lat, payload.Lng, nil
}
