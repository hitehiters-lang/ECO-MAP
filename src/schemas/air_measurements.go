package schemas

import "time"

type AirMeasurements struct {
	ID             int       `bun:"id,pk,autoincrement"`
	PointID        string    `bun:"point_id"`
	PointName      string    `bun:"point_name"`
	CoordX         float64   `bun:"coord_x"`
	CoordY         float64   `bun:"coord_y"`
	SampleTime     string    `bun:"sample_time"`
	SampleDatetime time.Time `bun:"sample_datetime"`
	TemperatureC   float64   `bun:"temperature_c"`
	PressureMmhg   float64   `bun:"pressure_mmhg"`
	HumidityPct    float64   `bun:"humidity_pct"`
	WindSpeedMs    float64   `bun:"wind_speed_ms"`
	WindDirection  string    `bun:"wind_direction"`
	CO             float64   `bun:"co"`
	NO             float64   `bun:"no"`
	NO2            float64   `bun:"no2"`
	SO2            float64   `bun:"so2"`
	PM25           float64   `bun:"pm25"`
	PM10           float64   `bun:"pm10"`
	CXHY           float64   `bun:"cxhy"`
	H2S            float64   `bun:"h2s"`
	VV             float64   `bun:"vv"`
	CH2O           float64   `bun:"ch2o"`
	RSh            float64   `bun:"r_sh"`
	CreatedAt      time.Time `bun:"created_at"`
}
