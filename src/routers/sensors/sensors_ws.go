package sensors

import (
	"eco-map-app/additional"
	"encoding/json"
	"fmt"
	"net/http"
)

// @Tags Датчики
// @Summary Получение резултатов датчика (Websocket)
// @Description Пример вебсокета. Сам роут не работает
// @ID sensors_ws
// @Accept json
// @Produce json
// @Success 200 {array} models.Sensor
// @Failure 400 "Bad Request"
// @Router /ws/sensors [get]
func SensorsWebsocketRUST(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(map[string]string{"error": "bad request"})
	additional.PrintError("/ws", fmt.Errorf("Попытка подключения к вебсокету из API"))
}
