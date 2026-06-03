package auth

import (
	"context"
	"eco-map-app/connection"
	"eco-map-app/schemas"
	"net/http"
)

func GetCurrentUser(r *http.Request) (schemas.Users, error) {
	var db = connection.GetDatabase()
	var user schemas.Users
	var ctx = context.Background()

	if id, ok := r.Context().Value(UserIDKey).(int); ok {
		err := db.NewSelect().Model(&user).Where("user_id = ?", id).Scan(ctx)
		if err != nil {
			return schemas.Users{}, err
		}
	}
	return user, nil
}
