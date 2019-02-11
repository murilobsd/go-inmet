package inmet

import (
	"context"
)

// CultureService TODO
type CultureService service

type CultureCycle struct {
	Code        string `json:"codigo"`
	Description string `json:"descricao"`
}
type cultureCycleResposne struct {
	CulturesCycle []*CultureCycle `json:"culturasciclo"`
}

// CycleList TODO
func (s *CultureService) CycleList(ctx context.Context) ([]*CultureCycle, *Response, error) {
	u := "culturasciclo.json"

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var culturesCycle cultureCycleResposne
	resp, err := s.client.Do(ctx, req, &culturesCycle)
	if err != nil {
		return nil, resp, err
	}

	return culturesCycle.CulturesCycle, resp, nil
}
