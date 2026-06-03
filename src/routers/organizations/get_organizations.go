package organizations

import (
	"eco-map-app/additional"
	"eco-map-app/connection"
	"eco-map-app/models"
	"eco-map-app/schemas"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// @Tags Организации
// @Summary Получить все организации в определённом городе
// @Description Возвращает все организации в городе
// @ID get_organizations
// @Accept json
// @Produce json
// @Param city_id path int true "ID города" example(1) default(1)
// @Success 200 {array} models.Organization
// @Failure 400 "{"error": "bad city id request"}"
// @Failure 404 "{"error": "organizations not found"}"
// @Router /api/orgs/{city_id} [get]
func GetOrganizations(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	city_id_str := chi.URLParam(r, "city_id")

	city_id, err := strconv.Atoi(city_id_str)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "bad city id request"}`, http.StatusBadRequest)
		return
	}

	var orgs []schemas.Organizations
	count, err := db.NewSelect().
		Model(&orgs).
		Relation("Category").
		Where("city_id = ?", city_id).
		ScanAndCount(ctx)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "organizations not found"}`, http.StatusNotFound)
		return
	}
	if count == 0 || len(orgs) == 0 {
		additional.PrintError(path, fmt.Errorf("Организации в городе не найдены"))
		http.Error(w, `{"error": "organizations not found"}`, http.StatusNotFound)
		return
	}

	var response_orgs []models.Organization
	response_orgs = make([]models.Organization, 0, len(orgs))
	for _, o := range orgs {
		if !o.Category.IsContainerYard {
			org := models.Organization{
				ID:         o.ID,
				Name:       o.Name,
				CategoryID: o.CategoryID,
				CityID:     o.CityID,
			}
			response_orgs = append(response_orgs, org)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response_orgs)

	additional.PrintSuccess(path, fmt.Sprintf("Получены данные об организациях (%d)", count))
}
