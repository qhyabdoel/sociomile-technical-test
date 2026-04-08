package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/middleware"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/service"
)

type ConversationHandler struct {
	svc *service.ConversationService
}

func NewConversationHandler(svc *service.ConversationService) *ConversationHandler {
	return &ConversationHandler{svc: svc}
}

func (h *ConversationHandler) HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var req model.WebhookRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	// service will get old conversation or create new conversation
	err := h.svc.ProcessIncomingMessage(r.Context(), req.TenantID, req.ExternalID, req.Message)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusOK, "Message processed successfully")
}

func (h *ConversationHandler) List(w http.ResponseWriter, r *http.Request) {
	tenantID, _ := r.Context().Value(middleware.TenantIDKey).(string)

	conversations, err := h.svc.GetConversationsByTenant(r.Context(), tenantID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to get conversations")
		return
	}

	WriteJSON(w, http.StatusOK, conversations)
}

// reply used by agent to reply message
func (h *ConversationHandler) Reply(w http.ResponseWriter, r *http.Request) {
	convID := chi.URLParam(r, "id")
	tenantID, _ := r.Context().Value(middleware.TenantIDKey).(string)

	var req struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid message")
		return
	}

	err := h.svc.AddAgentReply(r.Context(), tenantID, convID, req.Message)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusCreated, "Reply sent")
}
