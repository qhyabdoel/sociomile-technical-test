package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/service"
)

type TicketHandler struct {
	svc *service.TicketService
}

func NewTicketHandler(svc *service.TicketService) *TicketHandler {
	return &TicketHandler{svc: svc}
}

// create: escalate conversation to ticket
func (h *TicketHandler) Create(w http.ResponseWriter, r *http.Request) {
	tenantID := r.Context().Value("tenant_id").(int64)
	agentID := r.Context().Value("agent_id").(int64)

	var req struct {
		ConvID   int64  `json:"conv_id"`
		Title    string `json:"title"`
		Desc     string `json:"desc"`
		Priority string `json:"priority"`
	}

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

func (h *TicketHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")
	ticketID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid ticket ID")
		return
	}
	tenantID := r.Context().Value("tenant_id").(int64)

	var req struct {
		Status string `json:"status"`
	}

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
