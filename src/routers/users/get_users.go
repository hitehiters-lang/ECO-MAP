package users

import (
	"eco-map-app/additional"
	"eco-map-app/connection"
	"eco-map-app/models"
	"eco-map-app/schemas"
	"encoding/json"
	"fmt"
	"net/http"
)

// @Tags Пользователи
// @Summary Получить всех пользователей
// @Description Возвращает список всех пользователей
// @ID get_users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Success 200 {array} models.UserPublic
// @Failure 404 "{"error": "users not found"}"
// @Failure 403 "{"error": "not superuser"}"
// @Failure 401 "{"error": "unauthorized"}"
// @Router /api/users/list [get]
func GetUsers(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	var users []schemas.Users
	count, err := db.NewSelect().
		Model(&users).
		ScanAndCount(ctx)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "users not found"}`, http.StatusNotFound)
		return
	}
	if count == 0 || len(users) == 0 {
		additional.PrintError(path, fmt.Errorf("Пользователи не найдены"))
		http.Error(w, `{"error": "users not found"}`, http.StatusNotFound)
		return
	}

	var response_users []models.UserPublic
	for _, u := range users {
		user := models.UserPublic{
			ID:          u.ID,
			Login:       u.Login,
			CityID:      u.CityID,
			IsSuperuser: u.IsSuperuser,
		}
		response_users = append(response_users, user)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response_users)

	additional.PrintSuccess(path, fmt.Sprintf("Получены данные о пользователях (%d)", count))
}
