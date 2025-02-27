package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"risky_plumber/dto"
	"risky_plumber/service"
	"strings"
)

type RiskHandler struct {
	engine service.Engine
}

func NewRiskHandler(engine *service.Engine) *RiskHandler {
	return &RiskHandler{engine: *engine}
}

func (h *RiskHandler) HandleRiskRequests(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	switch r.Method {
	case http.MethodGet:
		h.ListRisks(ctx, w, r)
	case http.MethodPost:
		h.CreateRisk(ctx, w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *RiskHandler) HandleSingleRiskRequest(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	riskID := parts[3]

	if r.Method == http.MethodGet {
		h.GetRisk(ctx, w, r, riskID)
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func (h *RiskHandler) ListRisks(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	resp, err := h.engine.GetRiskList(ctx, &dto.EmptyRequest{})
	if err != nil {
		log.Println(fmt.Sprintf("ListRisk failed %v", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *RiskHandler) CreateRisk(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var req *dto.CreateRiskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println(fmt.Sprintf("ListRisk failed %v", err.Error()))
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.engine.CreateRisk(ctx, req)
	if err != nil {
		log.Println(fmt.Sprintf("CreateRisk failed %v", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

func (h *RiskHandler) GetRisk(ctx context.Context, w http.ResponseWriter, r *http.Request, id string) {
	resp, err := h.engine.GetRiskByID(ctx, &dto.GetRiskByIdRequest{RiskId: id})
	if err != nil {
		log.Println(fmt.Sprintf("GetRiskByID failed %v", err.Error()))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
