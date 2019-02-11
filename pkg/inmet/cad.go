package inmet

import (
	"context"
	"fmt"
)

// CADService TODO
type CADService service

type CADFinal struct {
	Value *float64 `json:"cadFinal"`
}

// GET TODO
func (s *CADService) GetFinal(ctx context.Context, soilTypeCode int64, cultureCycleCode string) (*float64, *Response, error) {
	u := fmt.Sprintf("solo/%v/%v/getCadFinal.json", soilTypeCode, cultureCycleCode)

	req, err := s.client.NewRequest("GET", u, nil)
	if err != nil {
		return nil, nil, err
	}

	var cadFinal *CADFinal
	resp, err := s.client.Do(ctx, req, &cadFinal)
	if err != nil {
		return nil, resp, err
	}

	return cadFinal.Value, resp, nil
}
