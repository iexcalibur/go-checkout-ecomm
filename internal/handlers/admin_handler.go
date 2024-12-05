package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/iexcalibur/backend/internal/core/services"
)

type AdminHandler struct {
	adminService *services.AdminService
}

func NewAdminHandler(adminService *services.AdminService) *AdminHandler {
	return &AdminHandler{
		adminService: adminService,
	}
}

func (h *AdminHandler) GetStats(w http.ResponseWriter, r *http.Request) {
	stats := h.adminService.GetStats()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

func (h *AdminHandler) GenerateDiscountCode(w http.ResponseWriter, r *http.Request) {
	discount, err := h.adminService.GenerateDiscountCode()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(discount)
}

func (h *AdminHandler) GetActiveDiscountCode(w http.ResponseWriter, r *http.Request) {
	discount, err := h.adminService.GetActiveDiscountCode()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	response := struct {
		Code    string    `json:"code"`
		Percent float64   `json:"percent"`
		ValidAt time.Time `json:"validAt"`
	}{
		Code:    discount.Code,
		Percent: discount.Percent,
		ValidAt: discount.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
