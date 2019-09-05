package router

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oribe1115/portfolio/server/model"
)

func PostNewSubImageHandler(c echo.Context) error {
	filePath, err := uploadSubImage(c)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to save sub image")
	}

	pathParam := c.Param("contentID")
	contentID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}

	subImage := model.SubImage{}

	subImage.URL = filePath
	subImage.ContentID = contentID

	newSubImage, err := model.NewSubImage(&subImage)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to save")
	}

	return c.JSON(http.StatusOK, newSubImage)
}

func uploadSubImage(c echo.Context) (string, error) {
	file, err := c.FormFile("subImage")
	if err != nil {
		fmt.Println("faild to get")
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		fmt.Println("faild to file open")
		return "", err
	}
	defer src.Close()

	slice := strings.Split(file.Filename, ".")
	fileName := uuid.New().String() + "." + slice[1]

	dst, err := os.Create("/portfolio/images/subImages/" + fileName)
	if err != nil {
		fmt.Println("faild to create")
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return "/portfolio/images/subImages/" + fileName, nil
}
