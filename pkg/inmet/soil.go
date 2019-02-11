package inmet

import (
	"context"
)

// SoilService TODO
type SoilService service

type SoilType struct {
	Code        int64  `json:"codigo"`
	Description string `json:"descricao"`
}
type soilTypeResposne struct {
	SoilsType []*SoilType `json:"solos"`
}

// TypeList TODO
func (s *SoilService) TypeList(ctx context.Context) ([]*SoilType, *Response, error) {
	u := "solos/list/select.json"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var soilsType soilTypeResposne
	resp, err := s.client.Do(ctx, req, &soilsType)
	if err != nil {
		return nil, resp, err
	}

	return soilsType.SoilsType, resp, nil
}
