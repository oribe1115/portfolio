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

	SCategories []SCategory `json:"sub_categories"`
}

type SCategory struct {
	ID             string    `json:"id"`
	MainCategoryID string    `json:"main_category_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func PostNewMainCategoryHandler(c echo.Context) error {
	mCategory := MCategory{}
	if err := c.Bind(&mCategory); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "faild to bind")
	}

	mainCategory, err := mCategory2MainCategory(mCategory)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to cast")
	}

	newMainCategory, err := model.NewMainCategory(&mainCategory)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to make new")
	}

	ignore := model.SubCategory{
		MainCategoryID: newMainCategory.ID,
		Name:           ".ignore",
		Description:    "inital sub_category",
	}

	_, err = model.NewSubCategory(&ignore)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to init sub_category")
	}

	thisMainCategory, err := model.GetMainCategoryByID(newMainCategory.ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get this main_category")
	}

	return c.JSON(http.StatusOK, mainCategory2MCategory(*thisMainCategory))
}

func PutMainCategoryHandler(c echo.Context) error {
	pathParam := c.Param("mainID")
	mainID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsMainCategory(mainID) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid mainID")
	}

	oldMainCategory, err := model.GetMainCategoryByID(mainID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get old one")
	}

	mCategory := MCategory{}
	if err := c.Bind(&mCategory); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "faild to bind")
	}

	oldMainCategory.Name = mCategory.Name
	oldMainCategory.Description = mCategory.Description

	_, err = model.SaveMainCategory(oldMainCategory)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to save")
	}

	newMainCategory, err := model.GetMainCategoryByID(mainID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get new one")

	}

	return c.JSON(http.StatusOK, mainCategory2MCategory(*newMainCategory))
}

func GetMainCategoriesHandler(c echo.Context) error {
	mainCategories, err := model.GetMainCategories()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get")
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
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	pathParam := c.Param("mainID")
	mainID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsMainCategory(mainID) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid mainID")
	}

	subCategory, err := sCategory2SubCategory(sCategory)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to parse")
	}
	subCategory.MainCategoryID = mainID

	newSubCategory, err := model.NewSubCategory(&subCategory)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to make new")
	}

	return c.JSON(http.StatusOK, newSubCategory)
}

func PutSubCategoryHandler(c echo.Context) error {
	pathParam := c.Param("subID")
	subID, err := uuid.Parse(pathParam)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid")
	}

	if !model.IsExistSubCategoryID(subID) {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid subID")
	}

	oldSubCategory, err := model.GetSubCategory(subID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get old one")
	}

	if oldSubCategory.Name == ".ignore" {
		return echo.NewHTTPError(http.StatusBadRequest, "this is prohibited to change")
	}

	sCategory := SCategory{}
	if err := c.Bind(&sCategory); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad request")
	}

	if sCategory.MainCategoryID != oldSubCategory.MainCategoryID.String() {
		newMainID, err := uuid.Parse(sCategory.MainCategoryID)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest, "invalid uuid of new main category id")
		}
		if !model.IsMainCategory(newMainID) {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid new main category")
		}
		oldSubCategory.MainCategoryID = newMainID
	}

	oldSubCategory.Name = sCategory.Name
	oldSubCategory.Description = sCategory.Description

	newSubCategory, err := model.SaveSubCategory(oldSubCategory)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to save")
	}

	return c.JSON(http.StatusOK, subCategory2SCategory(*newSubCategory))
}

func GetSubCategoriesHandler(c echo.Context) error {
	subCategories, err := model.GetSubCategories()
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get")
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
		ID:             subCategory.ID.String(),
		MainCategoryID: subCategory.MainCategoryID.String(),
		Name:           subCategory.Name,
		Description:    subCategory.Description,
		CreatedAt:      subCategory.CreatedAt,
		UpdatedAt:      subCategory.UpdatedAt,
	}
}
