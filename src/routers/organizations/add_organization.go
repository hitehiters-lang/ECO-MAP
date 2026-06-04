package organizations

import (
	"eco-map-app/additional"
	"eco-map-app/auth"
	"eco-map-app/connection"
	"eco-map-app/models"
	"eco-map-app/schemas"
	"encoding/json"
	"fmt"
	"net/http"
)

// @Tags Организации
// @Summary Добавить организацию
// @Description Создание новой организации в БД. Возвращает данные о созданной организации
// @ID add_organization
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body models.AddOrganizationBody true "Данные новой организации"
// @Success 200 {object} models.Organization
// @Failure 400 "{"error": "bad request"}"
// @Failure 401 "{"error": "unauthorized"}"
// @Failure 404 "{"error": "category not found"}"
// @Failure 500 "{"error": "internal server error"}"
// @Router /api/orgs/add [post]
func AddOrganization(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	user_check, err := auth.GetCurrentUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		additional.PrintError(path, err)
		return
	}

	var input models.AddOrganizationBody
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request"})
		additional.PrintError(path, err)
		return
	}

	var category_check schemas.Categories
	err = db.NewSelect().Model(&category_check).Where("id = ?", input.CategoryID).Scan(ctx)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "category not found"})
		additional.PrintError(path, err)
		return
	}

	new_organization := schemas.Organizations{
		Name:        input.Name,
		CategoryID:  input.CategoryID,
		CityID:      input.CityID,
		Description: input.Description,
		CoordX:      input.CoordX,
		CoordY:      input.CoordY,
	}
	_, err = db.NewInsert().Model(&new_organization).Returning("*").Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		additional.PrintError(path, err)
		return
	}

	org_response := models.Organization{
		ID:          new_organization.ID,
		Name:        new_organization.Name,
		CategoryID:  new_organization.CategoryID,
		CityID:      new_organization.CityID,
		Description: new_organization.Description,
		CoordX:      new_organization.CoordX,
		CoordY:      new_organization.CoordY,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(org_response)
	additional.PrintSuccess(path, fmt.Sprintf("Пользователь %v создал организацию %v на карте", user_check.Login, org_response.Name))
}
