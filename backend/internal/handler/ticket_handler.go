package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/middleware"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/service"
)

type TicketHandler struct {
	svc *service.TicketService
}

func NewTicketHandler(svc *service.TicketService) *TicketHandler {
	return &TicketHandler{svc: svc}
}

type CreateTicketRequest struct {
	ConvID   int64  `json:"conv_id"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Priority string `json:"priority"`
}

// @Summary Create ticket
// @Description Create ticket
// @Tags Ticket
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param body body CreateTicketRequest true "Ticket request"
// @Success 201 {string} string "Ticket created successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Failed to create ticket"
// @Router /tickets [post]
func (h *TicketHandler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID, _ := r.Context().Value(middleware.TenantIDKey).(int64)
	agentID, _ := r.Context().Value(middleware.UserIDKey).(int64)

	var req CreateTicketRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	ticket, err := h.svc.EscalateToTicket(r.Context(),
		tenantID,
		agentID,
		req.ConvID,
		req.Title,
		req.Desc,
		req.Priority,
	)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusCreated, ticket)
}

type UpdateTicketStatusRequest struct {
	Status string `json:"status"`
}

// @Summary Update ticket status
// @Description Update ticket status
// @Tags Ticket
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Ticket ID"
// @Param body body UpdateTicketStatusRequest true "Ticket status"
// @Success 200 {string} string "Ticket status updated successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Failed to update ticket status"
// @Router /tickets/{id}/status [patch]
func (h *TicketHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")
	ticketID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid ticket ID")
		return
	}
	tenantID, _ := r.Context().Value(middleware.TenantIDKey).(int64)

	var req UpdateTicketStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := h.svc.ChangeTicketStatus(r.Context(), tenantID, ticketID, req.Status); err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, "Ticket status updated successfully")
}
