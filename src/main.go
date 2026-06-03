package main

import (
	"eco-map-app/additional"
	"eco-map-app/config"
	"eco-map-app/connection"
	_ "eco-map-app/docs"
	"eco-map-app/routers/air"
	"eco-map-app/routers/organizations"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Eco Map API 2
// @version 0.1
// @description Документация к API.
// @BasePath /
func main() {
	additional.PrintMessage("Подключаемся к базе данных")
	db := connection.InitDatabase()
	if db == nil {
		os.Exit(1)
	}

	IPv4 := additional.GetLocalIP()

	r := chi.NewRouter()

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/api/docs/*", httpSwagger.WrapHandler)

	r.Get("/api/air/measurements", air.GetMeasurements)
	r.Get("/api/orgs/{city_id}", organizations.GetOrganizations)

	additional.PrintSuccess("", "Стартуем сервер на "+IPv4)
	additional.PrintSuccess("/api/docs/", "Документация доступна по адресу")

	http.ListenAndServe(config.GetServerPort(), r)
}
