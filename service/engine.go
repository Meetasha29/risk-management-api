package service

import (
	"risky_plumber/repository"
)

type Engine struct {
	RiskRepo repository.RiskInterface
}

func NewEngine(riskRepo repository.RiskInterface) *Engine {
	return &Engine{
		RiskRepo: riskRepo,
	}
}
