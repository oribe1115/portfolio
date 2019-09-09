package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/oribe1115/portfolio/server/model"
	"github.com/oribe1115/portfolio/server/router"
)

func main() {
	db, err := model.EstablishConnection()
	if err != nil {
		panic(err)
	}

	for {
		if err := db.DB().Ping(); err != nil {
			fmt.Println(err)
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}

	if err := model.Migration(); err != nil {
		panic(err)
	}

	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	e.Static("/images", "/portfolio/images")

	api := e.Group("/api")
	api.GET("/category", router.GetMainCategoriesHandler)
	api.POST("/category", router.PostNewMainCategoryHandler)
	api.GET("/category/sub", router.GetSubCategoriesHandler)
	api.POST("/category/sub/:mainID", router.PostNewSubCategoryHandler)

	api.GET("/content", router.GetContentDetailListHandler)
	api.POST("/content", router.PostNewContentHandler)
	api.GET("/content/:contentID", router.GetContentDeteilHandler)
	api.PUT("/content/:contentID", router.PutContentHandler)

	api.POST("/content/:contentID/tag/:tagID", router.PostNewTaggedContentHandler)
	api.POST("/content/:contentID/subImage", router.PostNewSubImageHandler)
	api.POST("/content/:contentID/mainImage", router.PostMainImageHandler)

	api.GET("/tag", router.GetTagListHandler)
	api.POST("/tag", router.PostNewTagHandler)

	api.DELETE("/taggedContent/:taggedContentID", router.DeleteTaggedContentHanlder)

	api.DELETE("/subImage/:subImageID", router.DeleteSubImageHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(e.Start(":" + port))
}
