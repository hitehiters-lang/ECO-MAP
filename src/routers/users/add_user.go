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
// @Summary Добавить пользователя (superuser)
// @Description Создание нового пользователя в БД. Доступно только superuser. Возвращает данные о созданном пользователе
// @ID add_user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body models.AddUserBody true "Данные нового пользователя"
// @Success 200 {object} models.UserPublic
// @Failure 400 "{"error": "bad request"}, {"error": "login is not unique"}"
// @Failure 403 "{"error": "not superuser"}"
// @Failure 401 "{"error": "unauthorized"}"
// @Failure 500 "{"error": "internal server error"}"
// @Router /api/users/add [post]
func AddUser(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	superuser_check, err := auth.GetCurrentUser(r)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "unauthorized"})
		additional.PrintError(path, err)
		return
	}
	if !superuser_check.IsSuperuser {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(map[string]string{"error": "not superuser"})
		additional.PrintError(path, fmt.Errorf("Попытка %v создать пользователя, будучи не superuser", superuser_check.Login))
		return
	}

	var input models.AddUserBody
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
		IsSuperuser: false,
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
	additional.PrintSuccess(path, fmt.Sprintf("Superuser %v создал пользователя %v", superuser_check.Login, user_response.Login))
}
