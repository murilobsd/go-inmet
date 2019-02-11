package inmet

import (
	"context"
)

// StationService TODO
type StationService service

type Station struct {
	Code string  `json:"codigoStr"`
	Lat  float64 `json:"latitude"`
	Lon  float64 `json:"longitude"`
	Name string  `json:"nome"`
}
type stationResposne struct {
	Stations []*Station `json:"estacoes"`
}

func (s *StationService) List(ctx context.Context) ([]*Station, *Response, error) {
	u := "estacoes/list.json"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var stations stationResposne
	resp, err := s.client.Do(ctx, req, &stations)
	if err != nil {
		return nil, resp, err
	}

	return stations.Stations, resp, nil
}
