package categories

import (
	"eco-map-app/additional"
	"eco-map-app/connection"
	"eco-map-app/models"
	"eco-map-app/schemas"
	"encoding/json"
	"fmt"
	"net/http"
)

// @Tags Категории
// @Summary Получить все категории
// @Description Возвращает все категории организаций и контейнерных площадок
// @ID get_categories
// @Accept json
// @Produce json
// @Success 200 {array} models.Category
// @Failure 404 "{"error": "categories not found"}"
// @Router /api/categories/list [get]
func GetCategories(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	var categories []schemas.Categories
	count, err := db.NewSelect().
		Model(&categories).
		ScanAndCount(ctx)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "organizations not found"}`, http.StatusNotFound)
		return
	}
	if count == 0 || len(categories) == 0 {
		additional.PrintError(path, fmt.Errorf("Категории не найдены"))
		http.Error(w, `{"error": "categories not found"}`, http.StatusNotFound)
		return
	}

	var response_categories []models.Category
	for _, c := range categories {
		category := models.Category{
			ID:              c.ID,
			Name:            c.Name,
			IsContainerYard: c.IsContainerYard,
		}
		response_categories = append(response_categories, category)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response_categories)

	additional.PrintSuccess(path, fmt.Sprintf("Получены данные о катеориях (%d)", count))
}
