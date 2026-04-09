package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/model"
	"github.com/qhyabdoel/sociomile-technical-test/backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthHandler(repo repository.UserRepository, jwtSecret string) *AuthHandler {
	return &AuthHandler{userRepo: repo, jwtSecret: jwtSecret}
}

// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body model.LoginRequest true "Login request"
// @Success 200 {string} string "Login successful"
// @Failure 400 {string} string "Invalid request body"
// @Failure 500 {string} string "Failed to login"
// @Router /login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	user, err := h.userRepo.FindByEmail(r.Context(), req.Email)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if user == nil {
		WriteError(w, http.StatusUnauthorized, "User not found")
		return
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		WriteError(w, http.StatusUnauthorized, "Wrong password")
		return
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.ID,
		"tenant_id": user.TenantID,
		"role":      user.Role,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(h.jwtSecret))

	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	WriteJSON(w, http.StatusOK, map[string]string{
		"token": tokenString,
	})
}
