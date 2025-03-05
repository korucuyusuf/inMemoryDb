package handler

import (
	"encoding/json"
	"inMemoryDb/service"
	"net/http"
)

type Handler struct {
	service *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]string

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil || requestData["name"] == "" {
		http.Error(w, "Geçersiz veri", http.StatusBadRequest)
		return
	}

	newData, err := h.service.CreateData(requestData["name"])
	if err != nil {
		http.Error(w, "Veri oluşturulamadı", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newData)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	data, err := h.service.GetAllData()
	if err != nil {
		http.Error(w, "Veriler alınamadı", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
