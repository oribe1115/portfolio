package router

import (
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oribe1115/portfolio/server/model"
)

func PostNewSubImageHandler(c echo.Context) error {
	fileName, err := uploadImage(c)
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

	if !model.IsExistContentID(contentID) {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid contentID")
	}

	subImage := model.SubImage{}

	subImage.Name = fileName
	subImage.URL = c.Scheme() + "://" + c.Request().Host + "/images/" + fileName
	subImage.ContentID = contentID

	newSubImage, err := model.NewSubImage(&subImage)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to save")
	}

	return c.JSON(http.StatusOK, newSubImage)
}

func DeleteSubImageHandler(c echo.Context) error {
	pathParam := c.Param("subImageID")
	subImageID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsExistSubImageID(subImageID) {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid subImageID")
	}

	subImage, err := model.GetSubImage(subImageID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get")
	}

	err = deleteImage(subImage.Name)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to delete file")
	}

	err = model.DeleteSubImage(subImage)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to delete on db")
	}

	return c.NoContent(http.StatusOK)
}

func uploadImage(c echo.Context) (string, error) {
	file, err := c.FormFile("image")
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	fileName := uuid.New().String() + ext

	dst, err := os.Create("/portfolio/images/" + fileName)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return "", err
	}

	return fileName, nil
}

func deleteImage(fileName string) error {
	err := os.Remove("/portfolio/images/" + fileName)
	if err != nil {
		return err
	}
	return nil
}
