package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"soap-server/internal/entities"
)

type Services struct {
	apiURLs map[string]string
	apiKey  string
	client  *http.Client
}

func InitServices(apiKey string) Services {
	apiUrls := make(map[string]string)
	apiUrls["kzhkProjects"] = "https://data.egov.kz/api/v4/kurylys_salushylardyn_sany7/v2"
	apiUrls["astanaBuildings"] = "https://data.egov.kz/api/v4/nur-sultan_kalasyndagy_turgyn_/v5"
	return Services{
		apiURLs: apiUrls,
		apiKey:  apiKey,
		client:  http.DefaultClient,
	}
}

func (s Services) GiveAstanaBuildings() ([]entities.AstanaBuilding, error) {
	var res []entities.AstanaBuilding
	// TODO: check DB
	resp, err := s.getDataFromApi(s.apiURLs["astanaBuildings"])
	if err != nil {
		return nil, err
	}
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	fmt.Println(res)

	return res, nil
}

func (s Services) GiveKZHKProjects() ([]entities.KHCAllowedProject, error) {
	var res []entities.KHCAllowedProject
	// TODO: check DB
	resp, err := s.getDataFromApi(s.apiURLs["kzhkProjects"])
	if err != nil {
		return nil, err
	}
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	// fmt.Println(res)
	return res, nil
}

func (s Services) GiveAbout() entities.About {
	abt := entities.About{
		Description:   "This service provides following data with the URLs:",
		AvailableData: make([]string, 0, len(s.apiURLs)),
	}
	for k, v := range s.apiURLs {
		abt.AvailableData = append(abt.AvailableData, "service: "+k+" URL: "+v)
	}
	return abt
}

func (s Services) getDataFromApi(apiUrl string) (*http.Response, error) {
	querries := url.Values{}
	querries.Set("apiKey", s.apiKey)
	return s.client.Get(apiUrl + "?" + querries.Encode())
}
