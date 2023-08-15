package server

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"io"
	"temperature-sensor-receiver/internal/checkerror"
	"temperature-sensor-receiver/internal/db"
)

func NewHTTPServer() (*gin.Engine, *httpServer) {
	httpServer := newHTTPServer()

	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(corsMiddleware())
	r.Use(LoggerMiddleware())

	r.POST("/", httpServer.processMessage)
	r.GET("/healthz", httpServer.healthz)

	return r, httpServer
}

func newHTTPServer() *httpServer {
	ctx, pool, err := db.NewDb()
	checkerror.CheckError(err)
	return &httpServer{
		Db:  pool,
		Ctx: ctx,
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log the incoming request details
		log.Info().Msgf("Request: %s %s", c.Request.Method, c.Request.URL.Path)

		// Read the request body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println("Error reading request body:", err)
			log.Info().Err(err).Msgf("%+v", err)
		} else {
			// Log the request body
			log.Info().Msgf("Request Body: %s", body)
			// Reset the request body, so it can be read again in the actual route handler
			c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		}

		c.Next()
	}
}
