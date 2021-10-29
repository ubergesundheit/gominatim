package gominatim

import (
	"fmt"
	"net/url"
	"strings"
)

type Config struct {
	OutputFormat string
	UserAgent    string
	Endpoint     string
}

type Coordinate float64

func (c *Coordinate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%.5f", float64(*c))), nil
}

type GeoJSONFeature struct {
	Type       string `json:"type"`
	Properties struct {
		Geocoding struct {
			PlaceID int    `json:"place_id"`
			OsmType string `json:"osm_type"`
			OsmID   int    `json:"osm_id"`
			Type    string `json:"type"`
			Label   string `json:"label"`
			Name    string `json:"name"`
		} `json:"geocoding"`
	} `json:"properties"`
	Geometry struct {
		Type        string       `json:"type"`
		Coordinates []Coordinate `json:"coordinates"`
	} `json:"geometry"`
}

type GeoJSONResult struct {
	Type      string `json:"type"`
	Geocoding struct {
		Version     string `json:"version"`
		Attribution string `json:"attribution"`
		Licence     string `json:"licence"`
		Query       string `json:"query"`
	} `json:"geocoding"`
	Features []GeoJSONFeature `json:"features"`
}

type Gominatim struct {
	config Config
}

type SearchParameters struct {
	Q          string
	City       string
	PostalCode string
	Country    string
}

func (s *SearchParameters) ToQuery() string {

	queryURLParts := []string{}

	if s.Q != "" {
		queryURLParts = append(queryURLParts, fmt.Sprintf("q=%s", url.QueryEscape(s.Q)))
	} else {
		if s.City != "" {
			queryURLParts = append(queryURLParts, fmt.Sprintf("city=%s", url.QueryEscape(s.City)))
		}
		if s.Country != "" {
			queryURLParts = append(queryURLParts, fmt.Sprintf("country=%s", url.QueryEscape(s.Country)))
		}
		if s.PostalCode != "" {
			queryURLParts = append(queryURLParts, fmt.Sprintf("postalcode=%s", url.QueryEscape(s.PostalCode)))
		}
	}

	return strings.Join(queryURLParts, "&")

}
