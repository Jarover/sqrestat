package handlers

import (
	"net/http"
	"time"

	"github.com/Jarover/sqrestat/internal/config"
	"github.com/gin-gonic/gin"
)

// StartPage
func startPage(c *gin.Context) {
	t := time.Now()

	c.JSON(http.StatusOK, gin.H{
		"version":     config.Version.VersionStr(),
		"data":        config.Version.BuildTime,
		"currentTime": t.Format("2006-01-02 15:04:05")})
}
