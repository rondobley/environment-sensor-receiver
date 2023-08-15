package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
	"net/http"
)

func (s *httpServer) processMessage(c *gin.Context) {
	log.Info().Msg("Processing message . . .")

	var message Message

	jsonData, _ := c.GetRawData()
	err := json.Unmarshal(jsonData, &message)
	if err != nil {
		log.Info().Err(err).Msgf("Error unmarshalling message: %s", jsonData)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad Request"})
		return
	}

	insertQuery := `INSERT INTO temperature_data (timestamp, sensor_id, dew_point_c, humidity_rh, pressure_pa, temperature_c)
		VALUES (@timestamp, @sensor, @dew_point, @humidity, @pressure, @temperature);`

	_, err = s.Db.Exec(s.Ctx, insertQuery, pgx.NamedArgs{
		"timestamp":   message.Time,
		"sensor":      message.SensorId,
		"dew_point":   message.DewPointC,
		"humidity":    message.HumidityRh,
		"pressure":    message.PressurePa,
		"temperature": message.TemperatureC,
	})
	if err != nil {
		log.Info().Err(err).Msgf("Error inserting data: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
	} else {
		c.Status(http.StatusNoContent)
	}
}
