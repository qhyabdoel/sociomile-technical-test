package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

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

// @Summary Handle incoming webhook
// @Description Handle incoming webhook from channel
// @Tags Conversation
// @Accept json
// @Produce json
// @Param body body model.WebhookRequest true "Webhook request"
// @Success 200 {string} string "Message processed successfully"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Failed to process message"
// @Router /channel/webhook [post]
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

// @Summary Get conversations by tenant
// @Description Get conversations by tenant
// @Tags Conversation
// @Security BearerAuth
// @Produce json
// @Success 200 {array} model.Conversation
// @Failure 500 {string} string "Failed to get conversations"
// @Router /conversations [get]
func (h *ConversationHandler) List(w http.ResponseWriter, r *http.Request) {
	tenantID, _ := r.Context().Value(middleware.TenantIDKey).(int64)
	conversations, err := h.svc.GetConversationsByTenant(r.Context(), tenantID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to get conversations")
		return
	}

	WriteJSON(w, http.StatusOK, conversations)
}

type ReplyRequest struct {
	Message string `json:"message"`
}

// @Summary Reply to conversation
// @Description Reply to conversation
// @Tags Conversation
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int64 true "Conversation ID"
// @Param body body handler.ReplyRequest true "Message"
// @Success 200 {string} string "Reply sent"
// @Failure 400 {string} string "Invalid conversation ID"
// @Failure 500 {string} string "Failed to reply"
// @Router /conversations/{id}/messages [post]
func (h *ConversationHandler) Reply(w http.ResponseWriter, r *http.Request) {
	paramID := chi.URLParam(r, "id")
	convID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid conversation ID")
		return
	}
	tenantID, _ := r.Context().Value(middleware.TenantIDKey).(int64)

	var req ReplyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid message")
		return
	}

	err = h.svc.AddAgentReply(r.Context(), tenantID, convID, req.Message)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	WriteJSON(w, http.StatusCreated, "Reply sent")
}

// @Summary Get conversation by ID
// @Description Get conversation by ID
// @Tags Conversation
// @Security BearerAuth
// @Produce json
// @Param id path int64 true "Conversation ID"
// @Success 200 {object} model.Conversation
// @Failure 400 {string} string "Invalid conversation ID"
// @Failure 500 {string} string "Failed to get conversation"
// @Router /conversations/{id} [get]
func (h *ConversationHandler) GetDetail(w http.ResponseWriter, r *http.Request) {
	// parse id from url
	paramID := chi.URLParam(r, "id")
	convID, err := strconv.ParseInt(paramID, 10, 64)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid conversation ID")
		return
	}

	// get tenant id from context
	tenantId, _ := r.Context().Value(middleware.TenantIDKey).(int64)

	// get conversation by id
	detail, err := h.svc.GetConversationByID(r.Context(), tenantId, convID)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to get conversation")
		return
	}

	WriteJSON(w, http.StatusOK, detail)
}
