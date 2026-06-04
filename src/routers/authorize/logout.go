package authorize

import (
	"eco-map-app/additional"
	"eco-map-app/auth"
	"encoding/json"
	"net/http"
)

// @Tags Аутентификация
// @Summary Выход из аккаунта
// @Description Кидает токен в чёрный список, делая невозможными действия с ним
// @ID log_out
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 "{"message": "success logout"}"
// @Failure 429 "{"error": "token has been revoked"}, {"error": "authorization header required"}, {"error": "invalid token"}"
// @Router /api/auth/logout [post]
func Logout(w http.ResponseWriter, r *http.Request) {

	route := "/api/auth/logout"
	w.Header().Set("Content-Type", "application/json")

	logout_user, _ := auth.GetCurrentUser(r)

	token := r.Header.Get("Authorization")
	claims, _ := auth.ValidateToken(token)
	auth.TokenBlacklist[token] = claims.ExpiresAt.Time
	go auth.CleanupBlacklist()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "success logout"})
	additional.PrintSuccess(route, "Пользователь "+logout_user.Login+" вышел")
}
