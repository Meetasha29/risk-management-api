package util

import "errors"

var (
	ErrInvalidState  = errors.New("Invalid state")
	ErrRiskNotFound  = errors.New("Risk not found")
	ErrInvalidRiskId = errors.New("Invalid Risk ID")
)
