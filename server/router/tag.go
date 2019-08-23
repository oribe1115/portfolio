package router

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oribe1115/portfolio/server/model"
)

type TagDetail struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func PostNewTagHandler(c echo.Context) error {
	tagDetail := TagDetail{}
	if err := c.Bind(&tagDetail); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	tag, err := tagDetail2Tag(tagDetail)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "faild to parse")
	}

	newTag, err := model.NewTag(&tag)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "faild to save")
	}

	return c.JSON(http.StatusOK, tag2TagDetail(*newTag))
}

func GetTagListHandler(c echo.Context) error {
	tagList, err := model.GetTagList()
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "faild to get")
	}

	tagDetailList := make([]TagDetail, 0)
	for _, tag := range tagList {
		tagDetailList = append(tagDetailList, tag2TagDetail(*tag))
	}

	return c.JSON(http.StatusOK, tagDetailList)
}

func PostNewTaggedContentHandler(c echo.Context) error {
	pathParam := c.Param("contentID")
	contentID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsExistContentID(contentID) {
		return c.String(http.StatusBadRequest, "invalid contentID")
	}

	pathParam = c.Param("tagID")
	tagID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsExistTagID(tagID) {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid tagID")
	}

	taggedContent := model.TaggedContent{
		TagID:     tagID,
		ContentID: contentID,
	}
	newTaggedContent, err := model.NewTaggedContent(&taggedContent)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "faild to save")
	}

	return c.JSON(http.StatusOK, newTaggedContent)
}

func tagDetail2Tag(tagDetail TagDetail) (model.Tag, error) {
	tag := model.Tag{
		Name:        tagDetail.Name,
		Description: tagDetail.Description,
	}

	if tagDetail.ID != "" {
		id, err := uuid.Parse(tagDetail.ID)
		if err != nil {
			return tag, err
		}
		tag.ID = id
	}

	return tag, nil
}

func tag2TagDetail(tag model.Tag) TagDetail {
	tagDetail := TagDetail{
		ID:          tag.ID.String(),
		Name:        tag.Name,
		Description: tag.Description,
		CreatedAt:   tag.CreatedAt,
		UpdatedAt:   tag.UpdatedAt,
	}

	return tagDetail
}
