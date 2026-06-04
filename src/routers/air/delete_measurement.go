package air

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

// @Tags Пробы
// @Summary Удалить пробу воздуха
// @Description Удаляет пробу воздуха из БД по ID
// @ID delete_measure
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param measure_id path int true "ID пробы (не point_id!)"
// @Success 200 "{"message": "measure successfully deleted"}"
// @Failure 400 "{"error": "bad request"}"
// @Failure 401 "{"error": "unauthorized"}"
// @Failure 404 "{"error": "measure not found"}"
// @Failure 500 "{"error": "internal server error"}"
// @Router /api/air/{measure_id} [delete]
func DeleteMeasurement(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	measure_id_str := chi.URLParam(r, "measure_id")
	measure_id, err := strconv.Atoi(measure_id_str)
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

	var delete_measure schemas.AirMeasurements
	res, err := db.NewDelete().Model(&delete_measure).Where("id = ?", measure_id).Exec(ctx)
	rows, _ := res.RowsAffected()
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "measure not found"})
		additional.PrintError(path, fmt.Errorf("Проба с ID %d не найдена", measure_id))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		additional.PrintError(path, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "measure successfully deleted"})
	additional.PrintSuccess(path, fmt.Sprintf("Пользователь %v удалил пробу с id %v", user_check.Login, measure_id))
}
