package router

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/oribe1115/portfolio/server/model"
)

type GeneralDataDetail struct {
	ID        string    `json:"id"`
	Subject   string    `json:"subject"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func PostNewGeneralDataHandler(c echo.Context) error {
	generalDataDetail := GeneralDataDetail{}
	if err := c.Bind(&generalDataDetail); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "bad request")
	}

	generalData, err := generalDataDetail2GeneralData(generalDataDetail)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to parse")
	}

	newGeneralData, err := model.NewGeneralData(&generalData)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to save")
	}

	return c.JSON(http.StatusOK, generalData2GeneralDataDetail(*newGeneralData))
}

func GetAllGeneralDataHandler(c echo.Context) error {
	generalDataList, err := model.GetAllGeneralData()

	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to get")
	}

	generalDataDetailList := make([]GeneralDataDetail, 0)
	for _, generalData := range generalDataList {
		generalDataDetailList = append(generalDataDetailList, generalData2GeneralDataDetail(*generalData))
	}

	return c.JSON(http.StatusOK, generalDataDetailList)
}

func generalDataDetail2GeneralData(generalDataDetail GeneralDataDetail) (model.GeneralData, error) {
	generalData := model.GeneralData{
		Subject: generalDataDetail.Subject,
		Content: generalDataDetail.Content,
	}

	if generalDataDetail.ID != "" {
		id, err := uuid.Parse(generalDataDetail.ID)
		if err != nil {
			return generalData, err
		}
		generalData.ID = id
	}

	return generalData, nil
}

func generalData2GeneralDataDetail(generalData model.GeneralData) GeneralDataDetail {
	generalDataDetail := GeneralDataDetail{
		ID:        generalData.ID.String(),
		Subject:   generalData.Subject,
		Content:   generalData.Content,
		CreatedAt: generalData.CreatedAt,
		UpdatedAt: generalData.UpdatedAt,
	}

	return generalDataDetail
}
