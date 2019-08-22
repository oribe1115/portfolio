package router

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oribe1115/portfolio/server/model"
)

type MCategory struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	SCategories []SCategory
}

type SCategory struct {
	ID          string    `json:"id"`
	MainCategoryID      string    `json:"main_category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func PostNewMainCategoryHandler(c echo.Context) error {
	mCategory := MCategory{}
	if err := c.Bind(&mCategory); err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusBadRequest)
	}

	mainCategory, err := mCategory2MainCategory(mCategory)
	if err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	newMainCategory, err := model.NewMainCategory(&mainCategory)
	if err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, mainCategory2MCategory(*newMainCategory))
}

func GetMainCategoriesHandler(c echo.Context) error {
	mainCategories, err := model.GetMainCategories()
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "faild to get")
	}

	mCategories := make([]MCategory, 0)
	for _, mainCategory := range mainCategories {
		mCategories = append(mCategories, mainCategory2MCategory(*mainCategory))
	}

	return c.JSON(http.StatusOK, mCategories)
}

func PostNewSubCategoryHandler(c echo.Context) error {
	sCategory := SCategory{}
	if err := c.Bind(&sCategory); err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "bad request")
	}

	pathParam := c.Param("mainID")
	mainID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsMainCategory(mainID) {
		return c.String(http.StatusBadRequest, "invalid mainID")
	}

	subCategory, err := sCategory2SubCategory(sCategory)
	if err != nil {
		c.Logger().Error(err)
		return c.String(http.StatusInternalServerError, "faild to parse")
	}
	subCategory.MainCategoryID = mainID

	newSubCategory, err := model.NewSubCategory(&subCategory)
	if err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.JSON(http.StatusOK, newSubCategory)
}

func GetSubCategoriesHandler(c echo.Context) error {
	subCategories, err := model.GetSubCategories()
	if err != nil {
		c.Logger().Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	sCategories := make([]SCategory, 0)
	for _, subCategory := range subCategories {
		sCategories = append(sCategories, subCategory2SCategory(*subCategory))
	}

	return c.JSON(http.StatusOK, sCategories)
}

func mCategory2MainCategory(mCategory MCategory) (model.MainCategory, error) {
	mainCategory := model.MainCategory{
		Name:        mCategory.Name,
		Description: mCategory.Description,
	}

	if mCategory.ID != "" {
		id, err := uuid.Parse(mCategory.ID)
		if err != nil {
			return mainCategory, err
		}
		mainCategory.ID = id
	}

	return mainCategory, nil
}

func mainCategory2MCategory(mainCategory model.MainCategory) MCategory {
	mCategory := MCategory{
		ID:          mainCategory.ID.String(),
		Name:        mainCategory.Name,
		Description: mainCategory.Description,
		CreatedAt:   mainCategory.CreatedAt,
		UpdatedAt:   mainCategory.UpdatedAt,
	}

	mCategory.SCategories = make([]SCategory, 0)
	for _, subCategory := range mainCategory.SubCategories {
		mCategory.SCategories = append(mCategory.SCategories, subCategory2SCategory(subCategory))
	}

	return mCategory
}

func sCategory2SubCategory(sCategory SCategory) (model.SubCategory, error) {
	subCategory := model.SubCategory{
		Name:        sCategory.Name,
		Description: sCategory.Description,
	}

	if sCategory.ID != "" {
		id, err := uuid.Parse(sCategory.ID)
		if err != nil {
			return subCategory, nil
		}
		subCategory.ID = id
	}

	return subCategory, nil
}

func subCategory2SCategory(subCategory model.SubCategory) SCategory {
	return SCategory{
		ID:          subCategory.ID.String(),
		MainCategoryID:      subCategory.MainCategoryID.String(),
		Name:        subCategory.Name,
		Description: subCategory.Description,
		CreatedAt:   subCategory.CreatedAt,
		UpdatedAt:   subCategory.UpdatedAt,
	}
}
