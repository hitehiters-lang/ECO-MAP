package categories

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

// @Tags Категории
// @Summary Добавить категорию
// @Description Создание новой категории в БД. Возвращает данные о созданной категории
// @ID add_category
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body models.AddCategoryBody true "Данные новой категории"
// @Success 200 {object} models.Category
// @Failure 400 "{"error": "bad request"}"
// @Failure 401 "{"error": "unauthorized"}"
// @Failure 500 "{"error": "internal server error"}"
// @Router /api/categories/add [post]
func AddCategory(w http.ResponseWriter, r *http.Request) {
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

	var input models.AddCategoryBody
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request"})
		additional.PrintError(path, err)
		return
	}

	new_category := schemas.Categories{
		Name:            input.Name,
		IsContainerYard: input.IsContainerYard,
	}
	_, err = db.NewInsert().Model(&new_category).Returning("*").Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		additional.PrintError(path, err)
		return
	}

	category_response := models.Category{
		ID:              new_category.ID,
		Name:            new_category.Name,
		IsContainerYard: new_category.IsContainerYard,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category_response)
	additional.PrintSuccess(path, fmt.Sprintf("Пользователь %v создал категорию %v", user_check.Login, category_response.Name))
}
