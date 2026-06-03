package schemas

import "time"

type Sensors struct {
	ID               int64     `bun:"id,pk,autoincrement"`
	SensorID         int       `bun:"sensor_id"`
	MeasuredAt       time.Time `bun:"measured_at"`
	TemperatureC     float64   `bun:"temperature_c"`
	PressureMmhg     float64   `bun:"pressure_mmhg"`
	HumidityPct      float64   `bun:"humidity_pct"`
	WindSpeedMs      float64   `bun:"wind_speed_ms"`
	WindDirectionDeg float64   `bun:"wind_direction_deg"`
	CO               float64   `bun:"co"`
	NO               float64   `bun:"no"`
	NO2              float64   `bun:"no2"`
	SO2              float64   `bun:"so2"`
	PM2_5            float64   `bun:"pm2_5"`
	PM10             float64   `bun:"pm10"`
	CXHY             float64   `bun:"cxhy"`
	H2S              float64   `bun:"h2s"`
	VV               float64   `bun:"vv"`
	CH2O             float64   `bun:"ch2o"`
	R_SH             float64   `bun:"r_sh"`
	Raw              string    `bun:"raw,type:jsonb"`
	CreatedAt        time.Time `bun:"created_at"`
}
