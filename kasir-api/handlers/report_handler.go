package handlers

import (
	"encoding/json"
	"kasir-api/services"
	"net/http"
	"time"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service:service}
}

func (h *ReportHandler) GetSummary(w http.ResponseWriter, r *http.Request) {
    startDate := r.URL.Query().Get("start_date")
    endDate := r.URL.Query().Get("end_date")

    // Default jika parameter kosong: Hari ini
    if startDate == "" {
        startDate = time.Now().Format("2006-01-02") + " 00:00:00"
    }
    if endDate == "" {
        endDate = time.Now().Format("2006-01-02") + " 23:59:59"
    }

    report, err := h.service.GetSalesSummary(startDate, endDate)
    // ... handling error dan kirim JSON ...

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}