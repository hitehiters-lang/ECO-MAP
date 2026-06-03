package ws

import (
	"context"
	"eco-map-app/additional"
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

				jsonData, err := json.Marshal(sensors)
				if err != nil {
					additional.PrintServerError(err.Error())
					continue
				}

				hub.broadcast <- jsonData
			}
		}
	}()
}
