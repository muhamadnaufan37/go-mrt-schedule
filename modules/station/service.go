package station

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"

	"github.com/muhamadnaufan37/go-mrt-schedule.git/modules/common/client"
)

type Service interface {
	GetAllStation() (response []StationListResponse, err error)
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
}

func (s *service) GetAllStation() (response []StationListResponse, err error) {
	url := "https://jsonplaceholder.typicode.com/posts"

	byteResponse, err := client.DoRequest(s.client, url)
	if err != nil {
		return
	}

	var station []Station
	err = json.Unmarshal(byteResponse, &station)
	if err != nil {
		return
	}

	for _, item := range station {
		response = append(response, StationListResponse{
			Stations: []Station{
				item,
			},
		})
	}

	return
}
