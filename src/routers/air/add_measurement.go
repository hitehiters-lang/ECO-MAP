package air

import (
	"eco-map-app/additional"
	"eco-map-app/connection"
	"eco-map-app/models"
	"eco-map-app/schemas"
	"encoding/json"
	"net/http"
	"time"
)

// @Tags Пробы
// @Summary Создать новую пробу
// @Description Возвращает ОК если проба успешно добавлена в БД
// @ID add_measurements
// @Accept json
// @Produce json
// @Param comment body models.AirMeasurements true "Данные пробы"
// @Success 201 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /api/air/measurements [post]
func AddMeasurements(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	ctx := r.Context()
	db := connection.GetDatabase()
	w.Header().Set("Content-Type", "application/json")

	var input models.AirMeasurements
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "bad request"}`, http.StatusBadRequest)
		return
	}
	parsedDate, err := time.Parse("2006-01-02", input.SampleDate)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "bad date request"}`, http.StatusBadRequest)
		return
	}

	full_measure := schemas.AirMeasurements{
		SampleNo:      input.SampleNo,
		SampleDate:    parsedDate,
		PointName:     input.PointName,
		CoordX:        input.CoordX,
		CoordY:        input.CoordY,
		TimeOfDay:     input.TimeOfDay,
		TemperatureC:  input.TemperatureC,
		PressureMmhg:  input.PressureMmhg,
		HumidityPct:   input.HumidityPct,
		WindSpeedMs:   input.WindSpeedMs,
		WindDirection: input.WindDirection,
		CO:            input.CO,
		NO:            input.NO,
		NO2:           input.NO2,
		SO2:           input.SO2,
		PM25:          input.PM25,
		PM10:          input.PM10,
		CXHY:          input.CXHY,
		H2S:           input.H2S,
		VV:            input.VV,
		CH2O:          input.CH2O,
		RSh:           input.RSh,
		CreatedAt:     time.Now(),
	}

	_, err = db.NewInsert().Model(&full_measure).Exec(ctx)
	if err != nil {
		additional.PrintError(path, err)
		http.Error(w, `{"error": "failed to create measure"}`, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "measurement successfully added"})

	additional.PrintSuccess(path, "Добавлена проба "+full_measure.PointName)
}
