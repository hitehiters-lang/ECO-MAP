package models

import "time"

type Sensor struct {
	ID               int64     `json:"id"`
	SensorID         int       `json:"sensor_id"`
	MeasuredAt       time.Time `json:"measured_at"`
	TemperatureC     float64   `json:"temperature_c"`
	PressureMmhg     float64   `json:"pressure_mmhg"`
	HumidityPct      float64   `json:"humidity_pct"`
	WindSpeedMs      float64   `json:"wind_speed_ms"`
	WindDirectionDeg float64   `json:"wind_direction_deg"`
	CO               float64   `json:"co"`
	NO               float64   `json:"no"`
	NO2              float64   `json:"no2"`
	SO2              float64   `json:"so2"`
	PM2_5            float64   `json:"pm2_5"`
	PM10             float64   `json:"pm10"`
	CXHY             float64   `json:"cxhy"`
	H2S              float64   `json:"h2s"`
	VV               float64   `json:"vv"`
	CH2O             float64   `json:"ch2o"`
	R_SH             float64   `json:"r_sh"`
	Raw              string    `json:"raw"`
	CreatedAt        time.Time `json:"created_at"`
	CoordX           float64   `json:"coord_x"`
	CoordY           float64   `json:"coord_y"`
}
