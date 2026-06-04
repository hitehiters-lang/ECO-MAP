package authorize

import (
	"eco-map-app/additional"
	"eco-map-app/auth"
	"eco-map-app/models"
	"encoding/json"
	"fmt"
	"net/http"
)

// @Tags Аутентификация
// @Summary Посмотреть данные о себе
// @Description Возвращает данные о своём пользователе из Authorization
// @ID me_user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {object} models.UserPublic
// @Failure 401 "{"error": "unauthorized"}"
// @Router /api/auth/me [get]
func MeUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	w.Header().Set("Content-Type", "application/json")

	self_user, err := auth.GetCurrentUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		additional.PrintError(path, err)
		return
	}

	user_response := models.UserPublic{
		ID:          self_user.ID,
		Login:       self_user.Login,
		CityID:      self_user.CityID,
		IsSuperuser: self_user.IsSuperuser,
		Name:        self_user.Name,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user_response)
	additional.PrintSuccess(path, fmt.Sprintf("Пользователь %v запросил данные о себе", user_response.Login))
}
