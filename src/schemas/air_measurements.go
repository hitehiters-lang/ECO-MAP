package schemas

import "time"

type AirMeasurements struct {
	ID            int       `json:"id" bun:"id,pk"`
	SampleNo      int       `json:"sample_no" bun:"sample_no"`
	SampleDate    time.Time `json:"sample_date" bun:"sample_date"`
	PointName     string    `json:"point_name" bun:"point_name"`
	CoordX        float64   `json:"coord_x" bun:"coord_x"`
	CoordY        float64   `json:"coord_y" bun:"coord_y"`
	TimeOfDay     string    `json:"time_of_day" bun:"time_of_day"`
	TemperatureC  float64   `json:"temperature_c" bun:"temperature_c"`
	PressureMmhg  float64   `json:"pressure_mmhg" bun:"pressure_mmhg"`
	HumidityPct   float64   `json:"humidity_pct" bun:"humidity_pct"`
	WindSpeedMs   float64   `json:"wind_speed_ms" bun:"wind_speed_ms"`
	WindDirection string    `json:"wind_direction" bun:"wind_direction"`
	CO            float64   `json:"co" bun:"co"`
	NO            float64   `json:"no" bun:"no"`
	NO2           float64   `json:"no2" bun:"no2"`
	SO2           float64   `json:"so2" bun:"so2"`
	PM25          float64   `json:"pm25" bun:"pm25"`
	PM10          float64   `json:"pm10" bun:"pm10"`
	CXHY          float64   `json:"cxhy" bun:"cxhy"`
	H2S           float64   `json:"h2s" bun:"h2s"`
	VV            float64   `json:"vv" bun:"vv"`
	CH2O          float64   `json:"ch2o" bun:"ch2o"`
	RSh           float64   `json:"r_sh" bun:"r_sh"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at"`
}
