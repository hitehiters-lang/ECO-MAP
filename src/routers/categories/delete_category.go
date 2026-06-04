package categories

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

// @Tags Категории
// @Summary Удалить категорию
// @Description Удаляет категорию из БД по ID. Также удаляет каскадом все организации, принадлежвашие этой категории
// @ID delete_category
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param category_id path int true "ID категории"
// @Success 200 "{"message": "category successfully deleted"}"
// @Failure 400 "{"error": "bad request"}"
// @Failure 401 "{"error": "unauthorized"}"
// @Failure 404 "{"error": "category not found"}"
// @Failure 500 "{"error": "internal server error"}"
// @Router /api/categories/{category_id} [delete]
func DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	category_id_str := chi.URLParam(r, "category_id")
	category_id, err := strconv.Atoi(category_id_str)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request"})
		additional.PrintError(path, err)
		return
	}

	user_check, err := auth.GetCurrentUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		additional.PrintError(path, err)
		return
	}

	var delete_category schemas.Categories
	res, err := db.NewDelete().Model(&delete_category).Returning("*").Where("id = ?", category_id).Exec(ctx)
	rows, _ := res.RowsAffected()
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "category not found"})
		additional.PrintError(path, fmt.Errorf("Категория с ID %d не найдена", category_id))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		additional.PrintError(path, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "category successfully deleted"})
	additional.PrintSuccess(path, fmt.Sprintf("Пользователь %v удалил организацию %v", user_check.Login, delete_category.Name))
}
