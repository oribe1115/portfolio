package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/oribe1115/portfolio/server/model"
	"golang.org/x/crypto/bcrypt"
)

type UserRequestBody struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func SignUpHandler(c echo.Context) error {
	if !model.IsNotExistUser() {
		return echo.NewHTTPError(http.StatusBadRequest, "user exist")
	}

	userRequestBody := UserRequestBody{}
	if err := c.Bind(&userRequestBody); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "faild to bind")
	}

	if userRequestBody.UserName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "user_name is missing")
	}

	if userRequestBody.UserName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "password is missing")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userRequestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to hash password")
	}

	user := model.User{
		UserName: userRequestBody.UserName,
		Password: hashedPassword,
	}

	_, err = model.NewUser(&user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "faild to save")
	}

	return c.NoContent(http.StatusCreated)
}

func LoginHandler(c echo.Context) error {
	userRequestBody := UserRequestBody{}
	if err := c.Bind(&userRequestBody); err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "faild to bind")
	}

	if userRequestBody.UserName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "user_name is missing")
	}

	if userRequestBody.UserName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "password is missing")
	}

	if !model.IsExistUserName(userRequestBody.UserName) {
		return echo.NewHTTPError(http.StatusBadRequest, "user_name not exist")
	}

	user, err := model.GetUser(userRequestBody.UserName)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "faild to get")
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(userRequestBody.Password))
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusBadRequest, "bad password")
	}

	return c.NoContent(http.StatusOK)
}
