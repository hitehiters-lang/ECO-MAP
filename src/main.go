package main

import (
	"context"
	"eco-map-app/additional"
	"eco-map-app/config"
	"eco-map-app/connection"
	_ "eco-map-app/docs"
	"eco-map-app/routers/air"
	"eco-map-app/routers/categories"
	"eco-map-app/routers/organizations"
	"eco-map-app/ws"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Eco Map API 2
// @version 0.13
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

	hub := ws.NewHub()
	go hub.Run()

	go ws.StartSensorsWebsocket(db, hub, context.Background())

	// CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080", "http://localhost:1234", "http://192.168.6.177:5173"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token", "Upgrade",
			"Sec-WebSocket-Key", "Sec-WebSocket-Version", "Sec-WebSocket-Extensions", "Connection"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/api/docs/*", httpSwagger.WrapHandler)

	r.Get("/api/air/measurements", air.GetMeasurements)

	r.Get("/api/orgs/{city_id}", organizations.GetOrganizations)

	r.Get("/api/categories/list", categories.GetCategories)

	r.Get("/ws/sensors", ws.WsHandler(hub))

	additional.PrintSuccess("", "Стартуем сервер на "+IPv4)
	additional.PrintSuccess("/api/docs/", "Документация доступна по адресу")

	http.ListenAndServe(config.GetServerPort(), r)
}
