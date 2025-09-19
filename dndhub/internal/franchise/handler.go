package franchise

import (
	"encoding/json"
	"net/http"

	"github.com/rafrey/dndhub/internal/user"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/create", h.handleCreateFranchise)
	r.Get("/", h.handleListFranchises)
	r.Get("/{id}", h.handleGetFranchise)
	r.Post("/{id}/workers", h.handleCreateWorker)
	r.Put("/{id}", h.handleUpdateFranchise)
	r.Put("/{id}/workers/{workerId}", h.handleUpdateWorker)
	r.Delete("/{id}/workers/{workerId}", h.handleDeleteWorker)

	return r
}

func (h *Handler) handleCreateFranchise(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := user.GetUserID(r)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var body struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uid, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	fr, err := h.service.CreateFranchise(r.Context(), uid, body.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fr)
}

func (h *Handler) handleListFranchises(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := user.GetUserID(r)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	uid, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	franchises, err := h.service.ListFranchises(r.Context(), uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(franchises)
}

func (h *Handler) handleGetFranchise(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := user.GetUserID(r)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	uid, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	franchiseIDStr := chi.URLParam(r, "id")
	fid, err := uuid.Parse(franchiseIDStr)
	if err != nil {
		http.Error(w, "invalid franchise id", http.StatusBadRequest)
		return
	}

	fr, err := h.service.GetFranchiseWithWorkers(r.Context(), uid, fid)
	if err != nil {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fr)
}

func (h *Handler) handleUpdateFranchise(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := user.GetUserID(r)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	uid, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	franchiseIDStr := chi.URLParam(r, "id")
	fid, err := uuid.Parse(franchiseIDStr)
	if err != nil {
		http.Error(w, "invalid franchise id", http.StatusBadRequest)
		return
	}

	var body Franchise
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.service.UpdateFranchise(r.Context(), uid, fid, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) handleCreateWorker(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := user.GetUserID(r)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	uid, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	franchiseIDStr := chi.URLParam(r, "id")
	fid, err := uuid.Parse(franchiseIDStr)
	if err != nil {
		http.Error(w, "invalid franchise id", http.StatusBadRequest)
		return
	}

	// confirm franchise belongs to this user
	_, err = h.service.GetFranchise(r.Context(), uid, fid)
	if err != nil {
		http.Error(w, "franchise not found or not yours", http.StatusForbidden)
		return
	}

	var body struct {
		Name             string  `json:"name"`
		MonthlyCostCents int     `json:"monthly_cost_cents"`
		Creativity       int     `json:"creativity"`
		Discipline       int     `json:"discipline"`
		Charisma         int     `json:"charisma"`
		Efficiency       int     `json:"efficiency"`
		Exploration      int     `json:"exploration"`
		Notes            *string `json:"notes"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newWorker := UniqueWorker{
		FranchiseID:      fid,
		Name:             body.Name,
		MonthlyCostCents: body.MonthlyCostCents,
		Creativity:       body.Creativity,
		Discipline:       body.Discipline,
		Charisma:         body.Charisma,
		Efficiency:       body.Efficiency,
		Exploration:      body.Exploration,
		Notes:            body.Notes,
	}

	worker, err := h.service.CreateWorker(r.Context(), fid, newWorker)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(worker)
}

func (h *Handler) handleUpdateWorker(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := user.GetUserID(r)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	uid, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	franchiseIDStr := chi.URLParam(r, "id")
	fid, err := uuid.Parse(franchiseIDStr)
	if err != nil {
		http.Error(w, "invalid franchise id", http.StatusBadRequest)
		return
	}

	workerIDStr := chi.URLParam(r, "workerId")
	wid, err := uuid.Parse(workerIDStr)
	if err != nil {
		http.Error(w, "invalid worker id", http.StatusBadRequest)
		return
	}

	// ensure franchise belongs to this user
	_, err = h.service.GetFranchise(r.Context(), uid, fid)
	if err != nil {
		http.Error(w, "franchise not found or not yours", http.StatusForbidden)
		return
	}

	var body UniqueWorker
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updated, err := h.service.UpdateWorker(r.Context(), fid, wid, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updated)
}

func (h *Handler) handleDeleteWorker(w http.ResponseWriter, r *http.Request) {
	userIDStr, ok := user.GetUserID(r)
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	uid, err := uuid.Parse(userIDStr)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusUnauthorized)
		return
	}

	franchiseIDStr := chi.URLParam(r, "id")
	fid, err := uuid.Parse(franchiseIDStr)
	if err != nil {
		http.Error(w, "invalid franchise id", http.StatusBadRequest)
		return
	}

	workerIDStr := chi.URLParam(r, "workerId")
	wid, err := uuid.Parse(workerIDStr)
	if err != nil {
		http.Error(w, "invalid worker id", http.StatusBadRequest)
		return
	}

	// confirm franchise belongs to this user
	_, err = h.service.GetFranchise(r.Context(), uid, fid)
	if err != nil {
		http.Error(w, "franchise not found or not yours", http.StatusForbidden)
		return
	}

	if err := h.service.DeleteWorker(r.Context(), fid, wid); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
