package router

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oribe1115/portfolio/server/model"
)

type ContentDetail struct {
	ID             string    `json:"id"`
	CategoryID     string    `json:"category_id"`
	Title          string    `json:"title"`
	Image          string    `json:"image"`
	Description    string    `json:"description"`
	Date           time.Time `json:"date"`
	SubImagesCount int       `json:"sub_images_count`
	SubImages      []SubImageDetail
	TaggedContents []TaggedContetDetail
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	SubCategory    SCategory
	MainCategory   MCategory
}

type ContentDetailForList struct {
	ID          string    `json:"id"`
	CategoryID  string    `json:"category_id"`
	Title       string    `json:"title"`
	Image       string    `json:"image"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type SubImageDetail struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	ContentID string    `json:"content_id"`
	URL       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaggedContetDetail struct {
	ID        string `json:"id"`
	TagID     string `json:"tag_id"`
	ContetID  string `json:"contet_id"`
	Tag       TagDetail
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var (
	defaultImagePath = "/defaultImage"
)

func PostNewContentHandler(c echo.Context) error {
	contentDetail := ContentDetail{}
	if err := c.Bind(&contentDetail); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "bad request")
	}

	content, err := contentDetail2Content(contentDetail)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to parse")
	}

	if !model.IsExistSubCategoryID(content.CategoryID) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid categoryID")
	}

	content.Title = "newcontent"
	content.Image = c.Scheme() + "://" + c.Request().Host + defaultImagePath
	content.Date = time.Now()

	newContent, err := model.NewContent(&content)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to save")
	}

	return c.JSON(http.StatusOK, content2ContentDetail(*newContent))
}

func PutContentHandler(c echo.Context) error {
	pathParam := c.Param("contentID")
	contentID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsExistContentID(contentID) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid contentID")
	}

	oldContent, err := model.GetContentByID(contentID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get old content")
	}

	contentDetail := ContentDetail{}
	if err := c.Bind(&contentDetail); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "bad request")
	}

	newContent, err := contentDetail2Content(contentDetail)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to parse")
	}

	if !model.IsExistSubCategoryID(newContent.CategoryID) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid categoryID")
	}

	oldContent.CategoryID = newContent.CategoryID
	oldContent.Title = newContent.Title
	oldContent.Image = newContent.Image
	oldContent.Description = newContent.Description
	oldContent.Date = newContent.Date

	updatedContent, err := model.SaveContent(oldContent)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to save")
	}

	return c.JSON(http.StatusOK, content2ContentDetail(*updatedContent))
}

func GetContentDetailListHandler(c echo.Context) error {
	contents, err := model.GetContentList()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get")
	}

	contentDetails := make([]ContentDetailForList, 0)
	for _, content := range contents {
		contentDetails = append(contentDetails, content2ContentDetailForList(*content))
	}

	return c.JSON(http.StatusOK, contentDetails)
}

func GetContentDetailListByMainCategoryHandler(c echo.Context) error {
	pathParam := c.Param("mainID")
	mainID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsMainCategory(mainID) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid mainID")
	}

	contents, err := model.GetContentListByMainCategory(mainID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get")
	}

	contentDetails := make([]ContentDetailForList, 0)
	for _, content := range contents {
		contentDetails = append(contentDetails, content2ContentDetailForList(*content))
	}

	return c.JSON(http.StatusOK, contentDetails)
}

func GetContentDetailListBySubCategoryHandler(c echo.Context) error {
	pathParam := c.Param("subID")
	subID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsExistSubCategoryID(subID) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid subID")
	}

	contents, err := model.GetContentListBySubCategory(subID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get")
	}

	contentDetails := make([]ContentDetailForList, 0)
	for _, content := range contents {
		contentDetails = append(contentDetails, content2ContentDetailForList(*content))
	}

	return c.JSON(http.StatusOK, contentDetails)
}

func GetContentDetailListByTag(c echo.Context) error {
	pathParam := c.Param("tagID")
	tagID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsExistTagID(tagID) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid tagID")
	}

	contents, err := model.GetContentListByTag(tagID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get")
	}

	contentDetails := make([]ContentDetailForList, 0)
	for _, content := range contents {
		contentDetails = append(contentDetails, content2ContentDetailForList(*content))
	}

	return c.JSON(http.StatusOK, contentDetails)

}

func GetContentDeteilHandler(c echo.Context) error {
	pathParam := c.Param("contentID")
	contentID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}

	content, err := model.GetContentByID(contentID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get")
	}

	return c.JSON(http.StatusOK, content2ContentDetail(*content))
}

func IGetContentDetailListByTag(c echo.Context) error {
	pathParam := c.Param("tagID")
	tagID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsExistTagID(tagID) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid tagID")
	}

	contents, err := model.IGetContentListByTag(tagID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get")
	}

	contentDetails := make([]ContentDetailForList, 0)
	for _, content := range contents {
		contentDetails = append(contentDetails, content2ContentDetailForList(*content))
	}

	return c.JSON(http.StatusOK, contentDetails)

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

	categoryID, err := uuid.Parse(contentDetail.CategoryID)
	if err != nil {
		return content, err
	}
	content.CategoryID = categoryID

	return content, nil
}

func content2ContentDetail(content model.Content) ContentDetail {
	contentDetail := ContentDetail{
		ID:          content.ID.String(),
		CategoryID:  content.CategoryID.String(),
		Title:       content.Title,
		Description: content.Description,
		Date:        content.Date,
		CreatedAt:   content.CreatedAt,
		UpdatedAt:   content.UpdatedAt,
	}

	if content.MainImage != nil {
		contentDetail.Image = content.MainImage.URL
	} else {
		contentDetail.Image = content.Image
	}

	if len(content.SubImages) != 0 {
		contentDetail.SubImages = make([]SubImageDetail, 0)
		for _, subImage := range content.SubImages {
			subImageDetail := subImage2subImageDetail(*subImage)
			contentDetail.SubImages = append(contentDetail.SubImages, subImageDetail)
		}
		contentDetail.SubImagesCount = len(contentDetail.SubImages)
	} else {
		contentDetail.SubImagesCount = 0
		contentDetail.SubImages = make([]SubImageDetail, 0)
	}

	contentDetail.TaggedContents = make([]TaggedContetDetail, 0)
	if len(content.TaggedContents) != 0 {
		for _, taggedContent := range content.TaggedContents {
			taggedContentDetail := taggedContent2TaggedContentDetail(*taggedContent)
			contentDetail.TaggedContents = append(contentDetail.TaggedContents, taggedContentDetail)
		}
	}

	contentDetail.SubCategory = subCategory2SCategory(content.SubCategory)
	contentDetail.MainCategory = mainCategory2MCategory(content.MainCategory)

	return contentDetail
}

func content2ContentDetailForList(content model.Content) ContentDetailForList {
	contentDetailForList := ContentDetailForList{
		ID:          content.ID.String(),
		CategoryID:  content.CategoryID.String(),
		Title:       content.Title,
		Description: content.Description,
		Date:        content.Date,
		CreatedAt:   content.CreatedAt,
		UpdatedAt:   content.UpdatedAt,
	}

	if content.MainImage != nil {
		contentDetailForList.Image = content.MainImage.URL
	} else {
		contentDetailForList.Image = content.Image
	}

	return contentDetailForList
}

func subImage2subImageDetail(subImage model.SubImage) SubImageDetail {
	subImageDetail := SubImageDetail{
		ID:        subImage.ID.String(),
		Name:      subImage.Name,
		ContentID: subImage.ContentID.String(),
		URL:       subImage.URL,
		CreatedAt: subImage.CreatedAt,
		UpdatedAt: subImage.UpdatedAt,
	}

	return subImageDetail
}

func taggedContent2TaggedContentDetail(taggedContent model.TaggedContent) TaggedContetDetail {
	taggedContetDetail := TaggedContetDetail{
		ID:        taggedContent.ID.String(),
		TagID:     taggedContent.TagID.String(),
		ContetID:  taggedContent.ContentID.String(),
		CreatedAt: taggedContent.CreatedAt,
		UpdatedAt: taggedContent.UpdatedAt,
	}

	taggedContetDetail.Tag = tag2TagDetail(*taggedContent.Tag)

	return taggedContetDetail
}
