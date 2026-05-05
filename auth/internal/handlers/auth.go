package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/calebs-company/auth/internal/token"
)

type loginRequest struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type verifyResponse struct {
	Valid  bool   `json:"valid"`
	UserID string `json:"user_id,omitempty"`
	Email  string `json:"email,omitempty"`
	Role   string `json:"role,omitempty"`
}

func IssueToken(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var req loginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	ttl := 30 * time.Minute
	t, err := token.Issue(req.UserID, req.Email, req.Role, ttl)
	if err != nil {
		http.Error(w, "could not issue token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenResponse{
		AccessToken: t,
		TokenType:   "bearer",
		ExpiresIn:   int(ttl.Seconds()),
	})
}

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	t := r.URL.Query().Get("token")
	claims, err := token.Verify(t)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(verifyResponse{Valid: false})
		return
	}
	json.NewEncoder(w).Encode(verifyResponse{
		Valid:  true,
		UserID: claims.UserID,
		Email:  claims.Email,
		Role:   claims.Role,
	})
}
