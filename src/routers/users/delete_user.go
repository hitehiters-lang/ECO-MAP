package users

import (
	"eco-map-app/additional"
	"eco-map-app/auth"
	"eco-map-app/connection"
	"eco-map-app/schemas"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// @Tags Пользователи
// @Summary Удалить пользователя (superuser)
// @Description Удаление пользователя из БД. Доступно только superuser. Не может удалять других superuser
// @ID delete_user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user_id path int true "ID пользователя"
// @Success 200 "{"message": "user successfully deleted"}"
// @Failure 400 "{"error": "bad request"}"
// @Failure 403 "{"error": "not superuser"}, {"error": "superuser can't delete superuser"}"
// @Failure 401 "{"error": "unauthorized"}"
// @Failure 404 "{"error": "user not found"}"
// @Failure 500 "{"error": "internal server error"}"
// @Router /api/users/{user_id} [delete]
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	user_id_str := chi.URLParam(r, "user_id")
	user_id, err := strconv.Atoi(user_id_str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad reqest"})
		additional.PrintError(path, err)
		return
	}

	superuser_check, err := auth.GetCurrentUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		additional.PrintError(path, err)
		return
	}
	if !superuser_check.IsSuperuser {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "not superuser"})
		additional.PrintError(path, fmt.Errorf("Попытка %v удалить пользователя, будучи не superuser", superuser_check.Login))
		return
	}

	var check_user schemas.Users
	err = db.NewSelect().Model(&check_user).Where("id = ?", user_id).Scan(ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "user not found"})
		additional.PrintError(path, fmt.Errorf("Пользователь с ID %d не был найден", user_id))
		return
	}
	if check_user.IsSuperuser {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "superuser can't delete superuser"})
		additional.PrintError(path,
			fmt.Errorf("Superuser %v попробовал удалить суперпользователя %v", superuser_check.Login, check_user.Login))
		return
	}

	var delete_user schemas.Users
	_, err = db.NewDelete().Model(&delete_user).Where("id = ?", user_id).Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		additional.PrintError(path, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "user successfully deleted"})
	additional.PrintSuccess(path, fmt.Sprintf("Superuser %v удалил пользователя %v", superuser_check.Login, check_user.Login))
}
