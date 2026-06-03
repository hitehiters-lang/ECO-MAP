package cities

import (
	"eco-map-app/additional"
	"eco-map-app/connection"
	"eco-map-app/models"
	"eco-map-app/schemas"
	"encoding/json"
	"fmt"
	"net/http"
)

// @Tags Города
// @Summary Получить все города
// @Description Возвращает список всех городов
// @ID get_cities
// @Accept json
// @Produce json
// @Success 200 {array} models.City
// @Failure 404 "{"error": "cities not found"}"
// @Router /api/cities/list [get]
func GetCities(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	var cities []schemas.Cities
	count, err := db.NewSelect().
		Model(&cities).
		ScanAndCount(ctx)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "cities not found"}`, http.StatusNotFound)
		return
	}
	if count == 0 || len(cities) == 0 {
		additional.PrintError(path, fmt.Errorf("Города не найдены"))
		http.Error(w, `{"error": "cities not found"}`, http.StatusNotFound)
		return
	}

	var response_cities []models.City
	for _, c := range cities {
		city := models.City{
			ID:       c.ID,
			CityName: c.CityName,
		}
		response_cities = append(response_cities, city)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response_cities)

	additional.PrintSuccess(path, fmt.Sprintf("Получены данные о городах (%d)", count))
}
