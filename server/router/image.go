package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oribe1115/portfolio/server/model"
)

func PostNewSubImageHandler(c echo.Context) error {
	subImage := model.SubImage{}
	if err := c.Bind(&subImage); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	newSubImage, err := model.NewSubImage(&subImage)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to save")
	}

	return c.JSON(http.StatusOK, newSubImage)
}
