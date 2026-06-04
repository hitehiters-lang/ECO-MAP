package authorize

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

// @Tags Аутентификация
// @Summary Войти
// @Description Вход по логину и паролю. Возвращает JWT
// @ID log_in
// @Accept json
// @Produce json
// @Param input body models.LoginBody true "Данные пользователя"
// @Success 200 "{"token": "token", "type": "bearer"}"
// @Failure 400 "{"error": "bad request"}"
// @Failure 404 "{"error": "invalid credentials"}"
// @Failure 500 "{"error": "internal server error"}"
// @Router /api/auth/login [post]
func Login(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	var input models.LoginBody
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request"})
		additional.PrintError(path, err)
		return
	}

	var user schemas.Users
	err = db.NewSelect().
		Model(&user).
		Where("login = ?", input.Login).
		Scan(ctx)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid credentials"})
		additional.PrintError(path, fmt.Errorf("Попытка логина пользователя '%s'", input.Login))
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{"error": "invalid credentials"})
		additional.PrintError(path, fmt.Errorf("Попытка логина пользователя '%s'", input.Login))
		return
	}

	token, err := auth.GenerateJWT(user.ID)
	if err != nil {
		additional.PrintServerError(err.Error())
		http.Error(w, `{"error": "internal server error"}`, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
		"type":  "bearer",
	})
	additional.PrintSuccess(path, "Пользователь "+user.Login+" вошёл в аккаунт")

}
