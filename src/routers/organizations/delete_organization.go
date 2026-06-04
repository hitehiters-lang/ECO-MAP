package organizations

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

// @Tags Организации
// @Summary Удалить организацию
// @Description Удаляет организацию из БД по ID
// @ID delete_organization
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param org_id path int true "ID организации"
// @Success 200 "{"message": "organization successfully deleted"}"
// @Failure 400 "{"error": "bad request"}"
// @Failure 401 "{"error": "unauthorized"}"
// @Failure 404 "{"error": "organization not found"}"
// @Failure 500 "{"error": "internal server error"}"
// @Router /api/orgs/{org_id} [delete]
func DeleteOrganization(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	org_id_str := chi.URLParam(r, "org_id")
	org_id, err := strconv.Atoi(org_id_str)
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

	var delete_measure schemas.Organizations
	res, err := db.NewDelete().Model(&delete_measure).Returning("*").Where("id = ?", org_id).Exec(ctx)
	rows, _ := res.RowsAffected()
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "organization not found"})
		additional.PrintError(path, fmt.Errorf("Организация с ID %d не найдена", org_id))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		additional.PrintError(path, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "organization successfully deleted"})
	additional.PrintSuccess(path, fmt.Sprintf("Пользователь %v удалил организацию %v", user_check.Login, delete_measure.Name))
}
