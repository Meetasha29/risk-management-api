package router

import (
	"net/http"
	"risky_plumber/handler"
)

// SetupRoutes registers all application routes
func SetupRoutes(riskHandler *handler.RiskHandler) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/risks", riskHandler.HandleRiskRequests)      // List & Create Risks
	mux.HandleFunc("/v1/risk/", riskHandler.HandleSingleRiskRequest) // Get Single Risk

	return mux
}
