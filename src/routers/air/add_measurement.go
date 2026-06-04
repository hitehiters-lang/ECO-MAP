package air

import (
	"eco-map-app/additional"
	"eco-map-app/auth"
	"eco-map-app/connection"
	"eco-map-app/models"
	"eco-map-app/schemas"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// @Tags Пробы
// @Summary Добавить пробу воздуха
// @Description Создание новой пробы воздуха в БД. Возвращает данные о созданной пробе
// @ID add_measure
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body models.AddMeasureBody true "Данные новой пробы"
// @Success 200 {object} models.Measure
// @Failure 400 "{"error": "bad request"}"
// @Failure 401 "{"error": "unauthorized"}"
// @Failure 409 "{"error": "that sample is already in db"}"
// @Failure 500 "{"error": "internal server error"}"
// @Router /api/air/add [post]
func AddMeasurement(w http.ResponseWriter, r *http.Request) {
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

	var input models.AddMeasureBody
	err = json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "bad request"})
		additional.PrintError(path, err)
		return
	}

	var unique_check schemas.AirMeasurements
	err = db.NewSelect().Model(&unique_check).Where("point_id = ? AND sample_datetime = ?", input.PointID, input.SampleDatetime).Scan(ctx)
	if err == nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(map[string]string{"error": "that sample is already in db"})
		additional.PrintError(path, fmt.Errorf("Попытка создать вторую пробу с таким же id и datetime"))
		return
	}

	new_measure := schemas.AirMeasurements{
		PointID:        input.PointID,
		PointName:      input.PointName,
		CoordX:         input.CoordX,
		CoordY:         input.CoordY,
		SampleTime:     input.SampleTime,
		SampleDatetime: input.SampleDatetime,
		TemperatureC:   input.TemperatureC,
		PressureMmhg:   input.PressureMmhg,
		HumidityPct:    input.HumidityPct,
		WindSpeedMs:    input.WindSpeedMs,
		WindDirection:  input.WindDirection,
		CO:             input.CO,
		NO:             input.NO,
		NO2:            input.NO2,
		SO2:            input.SO2,
		PM25:           input.PM25,
		PM10:           input.PM10,
		CXHY:           input.CXHY,
		H2S:            input.H2S,
		VV:             input.VV,
		CH2O:           input.CH2O,
		RSh:            input.RSh,
		CreatedAt:      time.Now(),
	}
	_, err = db.NewInsert().Model(&new_measure).Returning("*").Exec(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": "internal server error"})
		additional.PrintError(path, err)
		return
	}

	measure_response := models.Measure{
		ID:             new_measure.ID,
		PointID:        new_measure.PointID,
		PointName:      new_measure.PointName,
		CoordX:         new_measure.CoordX,
		CoordY:         new_measure.CoordY,
		SampleTime:     new_measure.SampleTime,
		SampleDatetime: new_measure.SampleDatetime,
		TemperatureC:   new_measure.TemperatureC,
		PressureMmhg:   new_measure.PressureMmhg,
		HumidityPct:    new_measure.HumidityPct,
		WindSpeedMs:    new_measure.WindSpeedMs,
		WindDirection:  new_measure.WindDirection,
		CO:             new_measure.CO,
		NO:             new_measure.NO,
		NO2:            new_measure.NO2,
		SO2:            new_measure.SO2,
		PM25:           new_measure.PM25,
		PM10:           new_measure.PM10,
		CXHY:           new_measure.CXHY,
		H2S:            new_measure.H2S,
		VV:             new_measure.VV,
		CH2O:           new_measure.CH2O,
		RSh:            new_measure.RSh,
		CreatedAt:      new_measure.CreatedAt,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(measure_response)
	additional.PrintSuccess(path,
		fmt.Sprintf("Пользователь %v создал пробу id %d, point_name %v",
			user_check.Login, measure_response.ID, measure_response.PointName))
}
