package models

import "time"

type Measure struct {
	ID             int       `json:"id"`
	PointID        string    `json:"point_id"`
	PointName      string    `json:"point_name"`
	CoordX         float64   `json:"coord_x"`
	CoordY         float64   `json:"coord_y"`
	SampleTime     string    `json:"sample_time"`
	SampleDatetime time.Time `json:"sample_datetime"`
	TemperatureC   float64   `json:"temperature_c"`
	PressureMmhg   float64   `json:"pressure_mmhg"`
	HumidityPct    float64   `json:"humidity_pct"`
	WindSpeedMs    float64   `json:"wind_speed_ms"`
	WindDirection  string    `json:"wind_direction"`
	CO             float64   `json:"co"`
	NO             float64   `json:"no"`
	NO2            float64   `json:"no2"`
	SO2            float64   `json:"so2"`
	PM25           float64   `json:"pm25"`
	PM10           float64   `json:"pm10"`
	CXHY           float64   `json:"cxhy"`
	H2S            float64   `json:"h2s"`
	VV             float64   `json:"vv"`
	CH2O           float64   `json:"ch2o"`
	RSh            float64   `json:"r_sh"`
	CreatedAt      time.Time `json:"created_at"`
}
