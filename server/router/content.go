package router

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oribe1115/portfolio/server/model"
)

type ContentDetail struct {
	ID          string    `json:"id"`
	CategoryID  string    `json:"category_id"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	// タグ
	// サブイメージ
}

func PostNewContentHandler(c echo.Context) error {
	contentDetail := ContentDetail{}
	if err := c.Bind(&contentDetail); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "bad request")
	}

	content, err := contentDetail2Content(contentDetail)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "faild to parse")
	}

	pathParam := c.Param("categoryID")
	categoryID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsExistCategoryID(categoryID) {
		return c.String(http.StatusBadRequest, "invalid categoryID")
	}

	content.CategoryID = categoryID

	newContent, err := model.NewContent(&content)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "faild to save")
	}

	return c.JSON(http.StatusOK, content2ContentDetail(*newContent))
}

func GetContentDetailsHandler(c echo.Context) error {
	contents, err := model.GetContentList()
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "faild to get")
	}

	contentDetails := make([]ContentDetail, 0)
	for _, content := range contents {
		contentDetails = append(contentDetails, content2ContentDetail(*content))
	}

	return c.JSON(http.StatusOK, contentDetails)
}

func GetContentDeteil(c echo.Context) error {
	pathParam := c.Param("contentID")
	contentID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid uuid")
	}

	content, err := model.GetContentByID(contentID)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "faild to get")
	}

	return c.JSON(http.StatusOK, content2ContentDetail(*content))
}

func contentDetail2Content(contentDetail ContentDetail) (model.Content, error) {
	content := model.Content{
		Title:       contentDetail.Title,
		Image:       contentDetail.Image,
		Description: contentDetail.Description,
		Date:        contentDetail.Date,
	}

	if contentDetail.ID != "" {
		id, err := uuid.Parse(contentDetail.ID)
		if err != nil {
			return content, err
		}
		content.ID = id
	}

	if contentDetail.CategoryID != "" {
		categoryID, err := uuid.Parse(contentDetail.CategoryID)
		if err != nil {
			return content, err
		}
		content.CategoryID = categoryID
	}

	return content, nil
}

func content2ContentDetail(content model.Content) ContentDetail {
	contentDetail := ContentDetail{
		ID:          content.ID.String(),
		CategoryID:  content.CategoryID.String(),
		Title:       content.Title,
		Image:       content.Image,
		Description: content.Description,
		Date:        content.Date,
		CreatedAt:   content.CreatedAt,
		UpdatedAt:   content.UpdatedAt,
	}

	return contentDetail
}
