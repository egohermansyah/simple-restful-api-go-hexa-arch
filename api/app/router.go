package app

import (
	"github.com/labstack/echo/v4"
	"simple-restful-api-go-hexa-arch/api/app/v1/role"
	"net/http"
)

func RegisterPath(e *echo.Echo, roleController *role.Controller) {
	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<strong>Hidup GarasiMan</strong>")
	})
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
	roleV1 := e.Group("api/v1/role")
	roleV1.POST("", roleController.Insert)
	roleV1.GET("/:id", roleController.FindById)
	roleV1.PUT("/:id", roleController.UpdateById)
	roleV1.DELETE("/:id", roleController.DeleteById)
	roleV1.GET("", roleController.List)
}
