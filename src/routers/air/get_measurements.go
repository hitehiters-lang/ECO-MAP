package air

import (
	"eco-map-app/additional"
	"eco-map-app/connection"
	"eco-map-app/schemas"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// @Tags Пробы
// @Summary Получить все пробы воздуха за период времени
// @Description Возвращает все пробы воздуха за дату ДО и ПОСЛЕ
// @ID air_measurements
// @Accept json
// @Produce json
// @Param date_from query string false "Дата с (YYYY-MM-DD)" example(2024-02-01) default(2024-02-01)
// @Param date_to query string false "Дата до (YYYY-MM-DD)" example(2024-02-28) default(2024-02-28)
// @Success 200 {array} schemas.AirMeasurements
// @Failure 400 {array} map[string]string
// @Failure 404 {array} map[string]string
// @Router /api/air/measurements [get]
func GetMeasurements(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	date_from_str := r.URL.Query().Get("date_from")
	date_to_str := r.URL.Query().Get("date_to")

	date_from, err := time.Parse("2006-01-02", date_from_str)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "bad date request"}`, http.StatusBadRequest)
		return
	}

	date_to, err := time.Parse("2006-01-02", date_to_str)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "bad date request"}`, http.StatusBadRequest)
		return
	}

	var air_samples []schemas.AirMeasurements
	err = db.NewSelect().Model(&air_samples).Where("sample_date >= ? AND sample_date <= ?", date_from, date_to).Scan(ctx)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "air samples not found"}`, http.StatusNotFound)
		return
	}
	if len(air_samples) == 0 {
		additional.PrintError(path, fmt.Errorf("Пробы не найдены по фильтрам"))
		http.Error(w, `{"error": "air samples not found"}`, http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(air_samples)

	additional.PrintSuccess(path, "Получены данные проб воздуха")
}
