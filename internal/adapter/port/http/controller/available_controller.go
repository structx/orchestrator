package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/structx/orchestrator/internal/core/domain"
)

// Available
type Available struct {
	service domain.AvailableService
}

// NewAvailable
func NewAvailable(availService domain.AvailableService) *Available {
	return &Available{
		// log:     logger.Sugar().Named("AvailableController"),
		service: availService,
	}
}

// ListRegionsResponse
type ListRegionsResponse struct {
	Regions []string `json:"regions"`
	Elapsed int64    `json:"elapsed"`
}

// RegisterRoutesV1
func (a *Available) RegisterRoutesV1() *chi.Mux {

	r := chi.NewRouter()

	r.Get("/avail/regions", a.ListRegions)

	return r
}

// ListRegions
func (a *Available) ListRegions(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	start := time.Now()

	domainSlice, err := a.service.ListRegions(ctx)
	if err != nil {
		// a.log.Errorf("failed to list regions %v", err)
		http.Error(w, "unable to list regions", http.StatusInternalServerError)
		return
	}

	regionSlice := make([]string, 0, len(domainSlice))
	for _, r := range domainSlice {
		regionSlice = append(regionSlice, string(r))
	}

	response := &ListRegionsResponse{
		Regions: regionSlice,
		Elapsed: time.Since(start).Milliseconds(),
	}

	w.WriteHeader(http.StatusAccepted)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		// a.log.Errorf("failed to encode response %v", err)
		http.Error(w, "unable to encode response", http.StatusInternalServerError)
		return
	}
}
