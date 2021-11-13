package gominatim

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func DefaultConfig() Config {
	return Config{
		UserAgent: "gominatim",
		Endpoint:  "https://nominatim.openstreetmap.org",
	}
}

func NewGominatim(config Config) (*Gominatim, error) {
	if config.Endpoint == "" {
		return nil, fmt.Errorf("Endpoint must not be empty")
	}
	if config.UserAgent == "" {
		return nil, fmt.Errorf("UserAgent must not be empty")
	}

	g := Gominatim{
		config: config,
	}

	return &g, nil
}

func (g *Gominatim) request(parameters SearchParameters) ([]byte, error) {
	requestURL := fmt.Sprintf("%s/search?%s&format=geocodejson&limit=1",
		g.config.Endpoint,
		parameters.ToQuery(),
	)

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
	var geoJSONResp GeoJSONResult

	respBytes, err := g.request(parameters)
	if err != nil {
		return geoJSONResp, err
	}

	err = json.Unmarshal(respBytes, &geoJSONResp)
	if err != nil {
		return geoJSONResp, err
	}

	return geoJSONResp, nil
}
