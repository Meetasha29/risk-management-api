package service

import (
	"context"
	"risky_plumber/dto"
)

func (e *Engine) CreateRisk(ctx context.Context, req *dto.CreateRiskRequest) (*dto.Risk, error) {
	risk, err := e.RiskRepo.CreateRisk(req.State, req.Title, req.Desc)
	if err != nil {
		return nil, err
	}

	return &dto.Risk{
		RiskId: risk.RiskId,
		State:  risk.State,
		Title:  risk.Title,
		Desc:   risk.Desc,
	}, nil
}

func (e *Engine) GetRiskByID(ctx context.Context, req *dto.GetRiskByIdRequest) (*dto.Risk, error) {
	risk, err := e.RiskRepo.GetRiskById(req.RiskId)
	if err != nil {
		return nil, err
	}

	return &dto.Risk{
		RiskId: risk.RiskId,
		State:  risk.State,
		Title:  risk.Title,
		Desc:   risk.Desc,
	}, nil
}

func (e *Engine) GetRiskList(ctx context.Context, req *dto.EmptyRequest) (*dto.GetRisksResponse, error) {
	risks, err := e.RiskRepo.GetRiskList()
	if err != nil {
		return nil, err
	}

	var riskList = make([]*dto.Risk, 0)
	for _, risk := range risks {
		riskList = append(riskList, &dto.Risk{
			RiskId: risk.RiskId,
			State:  risk.State,
			Title:  risk.Title,
			Desc:   risk.Desc,
		})
	}

	return &dto.GetRisksResponse{
		Risks: riskList,
	}, nil
}
