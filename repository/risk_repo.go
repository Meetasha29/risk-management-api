package repository

import (
	"risky_plumber/models"
	"risky_plumber/util"
	"slices"

	"github.com/google/uuid"
)

type InMemoryRiskRepo struct {
	Risks map[string]*models.Risk
}

func NewInMemoryRiskRepo() *InMemoryRiskRepo {
	return &InMemoryRiskRepo{
		Risks: make(map[string]*models.Risk),
	}
}

func (r *InMemoryRiskRepo) CreateRisk(state, title, desc string) (*models.Risk, error) {
	if !slices.Contains(util.ValidStates, state) {
		return nil, util.ErrInvalidState
	}

	riskUUID := uuid.New().String()
	newRisk := &models.Risk{
		RiskId: riskUUID,
		State:  state,
		Title:  title,
		Desc:   desc,
	}

	r.Risks[riskUUID] = newRisk

	return newRisk, nil
}

func (r *InMemoryRiskRepo) GetRiskById(riskId string) (*models.Risk, error) {

	if riskId == "" {
		return nil, util.ErrInvalidRiskId
	}

	risk, ok := r.Risks[riskId]
	if !ok {
		return nil, util.ErrRiskNotFound
	}

	return risk, nil
}

func (r *InMemoryRiskRepo) GetRiskList() ([]*models.Risk, error) {

	var riskList = make([]*models.Risk, 0)
	for _, risk := range r.Risks {
		riskList = append(riskList, risk)
	}

	return riskList, nil
}
