package gominatim

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewGominatim() (*Gominatim, error) {
	g := Gominatim{
		config: Config{
			UserAgent: "gominatim",
			Endpoint:  "https://nominatim.openstreetmap.org",
		},
	}

	return &g, nil
}

func (g *Gominatim) request(requestURL string) ([]byte, error) {
	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", g.config.UserAgent)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return respBytes, nil
}

func (g *Gominatim) Search(parameters SearchParameters) (GeoJSONResult, error) {
	queryURL := fmt.Sprintf("%s/search?%s&format=geocodejson&limit=1",
		g.config.Endpoint,
		parameters.ToQuery(),
	)

	var geoJSONResp GeoJSONResult

	respBytes, err := g.request(queryURL)
	if err != nil {
		return geoJSONResp, err
	}

	err = json.Unmarshal(respBytes, &geoJSONResp)
	if err != nil {
		return geoJSONResp, err
	}

	return geoJSONResp, nil
}
