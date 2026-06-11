package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

// LoginRequest contains login credentials
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
}

// LoginResponse contains user info after successful login
type LoginResponse struct {
	Success   bool      `json:"success"`
	Token     string    `json:"token"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	UserID    string    `json:"user_id"`
	LoginTime time.Time `json:"login_time"`
	Message   string    `json:"message"`
}

// LogoutRequest for logout
type LogoutRequest struct {
	UserID string `json:"user_id"`
}

// LogoutResponse for logout confirmation
type LogoutResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	LogoutAt time.Time `json:"logout_at"`
}

// Login handles user login and sets session
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Validate credentials (implement your actual auth logic)
	// This is a placeholder
	if req.Email == "" && req.Username == "" {
		http.Error(w, "Email or Username required", http.StatusBadRequest)
		return
	}

	// Create response with username
	username := req.Username
	if username == "" {
		username = "User"
	}

	response := LoginResponse{
		Success:   true,
		Username:  username,
		Email:     req.Email,
		UserID:    generateUserID(), // Implement your ID generation
		LoginTime: time.Now(),
		Token:     generateToken(),  // Implement your token generation
		Message:   "Login successful",
	}

	// Set secure cookie with username
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    username,
		Path:     "/",
		HttpOnly: false, // Allow JS access for watermark
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400 * 30, // 30 days
	})

	// Set auth token cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    response.Token,
		Path:     "/",
		HttpOnly: true, // Don't allow JS access
		Secure:   false, // Set to true in production with HTTPS
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400, // 24 hours
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Logout handles user logout
func Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Clear cookies
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: false,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "authToken",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	})

	response := LogoutResponse{
		Success:  true,
		Message:  "Logout successful",
		LogoutAt: time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetCurrentUser returns current logged-in user info
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	// Get username from cookie
	cookie, err := r.Cookie("username")
	var username string
	if err == nil {
		username = cookie.Value
	} else {
		username = "Guest"
	}

	// Check if auth token exists
	_, tokenErr := r.Cookie("authToken")
	isLoggedIn := tokenErr == nil

	response := map[string]interface{}{
		"username":    username,
		"isLoggedIn":  isLoggedIn,
		"currentTime": time.Now(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Helper functions
func generateUserID() string {
	return "user_" + time.Now().Format("20060102150405")
}

func generateToken() string {
	// Implement proper JWT generation
	return "token_" + time.Now().Format("20060102150405")
}
