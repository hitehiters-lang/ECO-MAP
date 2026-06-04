package models

import "time"

type AddMeasureBody struct {
	PointID        string    `json:"point_id" example:"AS-1"`
	PointName      string    `json:"point_name" example:"AS-1"`
	CoordX         float64   `json:"coord_x" example:"47.122992"`
	CoordY         float64   `json:"coord_y" example:"51.831142"`
	SampleTime     string    `json:"sample_time" example:"12:00"`
	SampleDatetime time.Time `json:"sample_datetime" example:"2024-02-14T04:00:00+03:00"`
	TemperatureC   float64   `json:"temperature_c" example:"0"`
	PressureMmhg   float64   `json:"pressure_mmhg" example:"764"`
	HumidityPct    float64   `json:"humidity_pct" example:"91"`
	WindSpeedMs    float64   `json:"wind_speed_ms" example:"8"`
	WindDirection  string    `json:"wind_direction" example:"СВ"`
	CO             float64   `json:"co" example:"0"`
	NO             float64   `json:"no" example:"0"`
	NO2            float64   `json:"no2" example:"0"`
	SO2            float64   `json:"so2" example:"0"`
	PM25           float64   `json:"pm25" example:"0"`
	PM10           float64   `json:"pm10" example:"0"`
	CXHY           float64   `json:"cxhy" example:"0"`
	H2S            float64   `json:"h2s" example:"0"`
	VV             float64   `json:"vv" example:"0"`
	CH2O           float64   `json:"ch2o" example:"0"`
	RSh            float64   `json:"r_sh" example:"0"`
}
