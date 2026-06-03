package air

import (
	"eco-map-app/additional"
	"eco-map-app/connection"
	"eco-map-app/models"
	"eco-map-app/schemas"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// @Tags Пробы
// @Summary Получить все пробы воздуха за период времени
// @Description Возвращает все пробы воздуха за дату ДО и ПОСЛЕ
// @ID air_measurements
// @Accept json
// @Produce json
// @Param date_from query string true "Дата с (YYYY-MM-DD)" example(2024-02-01) default(2024-02-01)
// @Param date_to query string true "Дата до (YYYY-MM-DD)" example(2024-02-28) default(2024-02-28)
// @Param time query string true "Время пробы" Enums(01:00, 07:00, 12:00, 18:00)
// @Success 200 {array} schemas.AirMeasurements
// @Failure 400 "{"error": "bad datetime request"}"
// @Failure 404 "{"error": "air samples not found"}"
// @Router /api/air/measurements [get]
func GetMeasurements(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	date_from_str := r.URL.Query().Get("date_from")
	date_to_str := r.URL.Query().Get("date_to")
	time_str := r.URL.Query().Get("time")

	date_from, err := time.Parse("2006-01-02", date_from_str)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "bad datetime request"}`, http.StatusBadRequest)
		return
	}

	date_to, err := time.Parse("2006-01-02", date_to_str)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "bad datetime request"}`, http.StatusBadRequest)
		return
	}

	var air_samples []schemas.AirMeasurements
	err = db.NewSelect().
		Model(&air_samples).
		Where("sample_datetime >= ? AND sample_datetime <= ? AND sample_time = ?", date_from, date_to, time_str).
		Scan(ctx)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "air samples not found"}`, http.StatusNotFound)
		return
	}
	if len(air_samples) == 0 {
		additional.PrintError(path, fmt.Errorf("Пробы не найдены по фильтрам"))
		http.Error(w, `{"error": "air samples not found"}`, http.StatusNotFound)
		return
	}

	var response_samples []models.Measure
	response_samples = make([]models.Measure, 0, len(air_samples))
	for _, s := range air_samples {
		measure := models.Measure{
			ID:             s.ID,
			PointID:        s.PointID,
			PointName:      s.PointName,
			CoordX:         s.CoordX,
			CoordY:         s.CoordY,
			SampleTime:     s.SampleTime,
			SampleDatetime: s.SampleDatetime,
			TemperatureC:   s.TemperatureC,
			PressureMmhg:   s.PressureMmhg,
			HumidityPct:    s.HumidityPct,
			WindSpeedMs:    s.WindSpeedMs,
			WindDirection:  s.WindDirection,
			CO:             s.CO,
			NO:             s.NO,
			NO2:            s.NO2,
			SO2:            s.SO2,
			PM25:           s.PM25,
			PM10:           s.PM10,
			CXHY:           s.CXHY,
			H2S:            s.H2S,
			VV:             s.VV,
			CH2O:           s.CH2O,
			RSh:            s.RSh,
			CreatedAt:      s.CreatedAt,
		}
		response_samples = append(response_samples, measure)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response_samples)

	additional.PrintSuccess(path, "Получены данные проб воздуха")
}
