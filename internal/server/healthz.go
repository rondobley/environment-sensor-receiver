package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *httpServer) healthz(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte("Healthy"))
}
