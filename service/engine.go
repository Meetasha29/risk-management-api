package service

import (
	"risky_plumber/repository"
)

// Engine is the core service layer that interacts with the risk repository.
type Engine struct {
	RiskRepo repository.RiskInterface
}

func NewEngine(riskRepo repository.RiskInterface) *Engine {
	return &Engine{
		RiskRepo: riskRepo,
	}
}
