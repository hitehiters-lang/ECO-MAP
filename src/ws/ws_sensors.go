package ws

import (
	"context"
	"eco-map-app/additional"
	"eco-map-app/models"
	"eco-map-app/schemas"
	"encoding/json"
	"time"

	"github.com/uptrace/bun"
)

func StartSensorsWebsocket(db *bun.DB, hub *Hub, ctx context.Context) {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				additional.PrintSuccess("/ws", "Остановка стримера данных")
				return
			case <-ticker.C:

				var sensors []schemas.Sensors
				err := db.NewSelect().
					Model(&sensors).
					DistinctOn("sensor_id").
					OrderExpr("sensor_id ASC, measured_at DESC").
					Scan(ctx)
				if err != nil {
					additional.PrintServerError(err.Error())
					continue
				}
				var response_sensors []models.Sensor
				for _, s := range sensors {
					sensor := models.Sensor{
						ID:               s.ID,
						SensorID:         s.SensorID,
						MeasuredAt:       s.MeasuredAt,
						TemperatureC:     s.TemperatureC,
						PressureMmhg:     s.PressureMmhg,
						HumidityPct:      s.HumidityPct,
						WindSpeedMs:      s.WindSpeedMs,
						WindDirectionDeg: s.WindDirectionDeg,
						CO:               s.CO,
						NO:               s.NO,
						NO2:              s.NO2,
						SO2:              s.SO2,
						PM2_5:            s.PM2_5,
						PM10:             s.PM10,
						CXHY:             s.CXHY,
						H2S:              s.H2S,
						VV:               s.VV,
						CH2O:             s.CH2O,
						R_SH:             s.R_SH,
						Raw:              s.Raw,
						CreatedAt:        s.CreatedAt,
						CoordX:           s.CoordX,
						CoordY:           s.CoordY,
					}
					response_sensors = append(response_sensors, sensor)
				}

				jsonData, err := json.Marshal(response_sensors)
				if err != nil {
					additional.PrintServerError(err.Error())
					continue
				}

				hub.broadcast <- jsonData
			}
		}
	}()
}
