package app

import (
	"net/http"
	"runtime"
	"time"

	"github.com/labstack/echo"
)

type statusHandler struct {
	healthService HealthCheckService
	settings      *WechySettings
}

// Status creates a new Status HTTP handler
func Status(healthService HealthCheckService, settings *WechySettings) echo.HandlerFunc {
	return statusHandler{healthService, settings}.get()
}

func (h statusHandler) get() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"build":   h.settings.BuildTime,
			"vresion": h.settings.Version,
			"healthy": echo.Map{
				"database": h.healthService.IsDatabaseOnline(),
			},
			"version": runtime.Version(),
			"now":     time.Now().Format("2006.01.02.150405"),
		})
	}
}
