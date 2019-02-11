package inmet

import (
	"context"
	"fmt"
)

// BHCService TODO
type BHCService service

type BHCCultivationLossProductivity struct {
	TempMax           float64 `json:"tmax"`
	Deficit           float64 `json:"deficit"`
	Excess            float64 `json:"excesso"`
	Penalty           float64 `json:"penalty"`
	ETC               float64 `json:"etc"`
	Date              string  `json:"date"`
	Temp              float64 `json:"temperature"`
	Preciptation      float64 `json:"precipitacao"`
	ETP               float64 `json:"etp"`
	Alteration        float64 `json:"alteracao"`
	ETRXTPC           float64 `json:"etrcetpc"`
	AcumulatedPenalty float64 `json:"penalidadeAcumulada"`
	CAD               float64 `json:"cad"`
	ETR               float64 `json:"etr"`
	PlantingDays      float64 `json:"diasPosPlantio"`
	CADFinal          float64 `json:"cadFinal"`
	KC                float64 `json:"kc"`
	ARMPercentual     float64 `json:"armPercentual"`
	Notice            bool    `json:"aviso"`
	ARM               float64 `json:"arm"`
	TempMin           float64 `json:"tmin"`
	Productivity      float64 `json:"produtividade"`
	PETC              float64 `json:"p_etc"`
}

type bhcCultivationLossProdResponse struct {
	sucess      string                            `json:"success"`
	BHCLossProd []*BHCCultivationLossProductivity `json:"bhc"`
}

type BHCLossProdRequest struct {
	DatePlanting string `json:"dataPlantio"`
	CultureCode  string `json:"culturaId"`
	StationCode  string `json:"estacaoId"`
	SoilCode     string `json:"soloId"`
	CADFinal     string `json:"cad"`
}

// BHCLossProdGet TODO
func (s *BHCService) BHCLossProdGet(ctx context.Context, bhcRequest *BHCLossProdRequest) ([]*BHCCultivationLossProductivity, *Response, error) {
	u := fmt.Sprintf("monitoramento/bhc.json")

	req, err := s.client.NewRequest("POST", u, bhcRequest)
	if err != nil {
		return nil, nil, err
	}

	var bhcLossProd bhcCultivationLossProdResponse
	resp, err := s.client.Do(ctx, req, &bhcLossProd)
	if err != nil {
		return nil, resp, err
	}

	return bhcLossProd.BHCLossProd, resp, nil
}
