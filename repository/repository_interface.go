package repository

import (
	"risky_plumber/models"
)

type RiskInterface interface {
	GetRiskById(riskId string) (*models.Risk, error)
	GetRiskList() ([]*models.Risk, error)
	CreateRisk(state, title, desc string) (*models.Risk, error)
}
