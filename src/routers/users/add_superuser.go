package users

import (
	"eco-map-app/additional"
	"eco-map-app/auth"
	"eco-map-app/connection"
	"eco-map-app/models"
	"eco-map-app/schemas"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// @Tags Пользователи
// @Summary Добавить пользователя (root)
// @Description Создание нового пользователя в БД. Доступно только root. Возвращает данные о созданном пользователе
// @ID add_superuser
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body models.AddSuperuserBody true "Данные нового пользователя"
// @Success 200 {object} models.UserPublic
// @Failure 400 "{"error": "bad request"}, {"error": "login is not unique"}"
// @Failure 403 "{"error": "not root"}"
// @Failure 401 "{"error": "unauthorized"}"
// @Failure 500 "{"error": "internal server error"}"
// @Router /api/users/root/add [post]
func AddSuperuser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	root_check, err := auth.GetCurrentUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		additional.PrintError(path, err)
		return
	}
	if root_check.Login != "root" {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "not root"})
		additional.PrintError(path, fmt.Errorf("Попытка %v создать пользователя, будучи не root", root_check.Login))
		return
	}

	var input models.AddSuperuserBody
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request"})
		additional.PrintError(path, err)
		return
	}

	var unique_check schemas.Users
	err = db.NewSelect().Model(&unique_check).Where("login = ?", input.Login).Scan(ctx)
	if err == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "login is not unique"})
		additional.PrintError(path, fmt.Errorf("Попытка создать второго пользователя %v", input.Login))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		additional.PrintError(path, err)
		return
	}

	var new_user schemas.Users
	new_user = schemas.Users{
		Login:       input.Login,
		Password:    string(hashedPassword),
		CityID:      input.CityID,
		IsSuperuser: input.IsSuperuser,
		Name:        input.Name,
	}
	_, err = db.NewInsert().Model(&new_user).Returning("*").Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		additional.PrintError(path, err)
		return
	}

	user_response := models.UserPublic{
		ID:          new_user.ID,
		Login:       new_user.Login,
		CityID:      new_user.CityID,
		IsSuperuser: new_user.IsSuperuser,
		Name:        new_user.Name,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user_response)
	additional.PrintSuccess(path, fmt.Sprintf("Root создал пользователя %v", user_response.Login))
}
