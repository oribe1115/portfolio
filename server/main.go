package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
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

	cookieSecret := os.Getenv("COOKIE_SECRET")
	if cookieSecret == "" {
		cookieSecret = "portfolio"
	}

	store := sessions.NewCookieStore([]byte("secret"))

	e := echo.New()
	e.Debug = true
	e.Use(middleware.Logger())

	// middleware.StaticWithConfigからほぼコピペ・一部改変
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) (err error) {
			p := c.Request().URL.Path
			if strings.HasSuffix(c.Path(), "*") { // When serving from a group, e.g. `/static*`.
				p = c.Param("*")
			}
			p, err = url.PathUnescape(p)
			if err != nil {
				return
			}
			name := filepath.Join("./static", path.Clean("/"+p)) // "/"+ for security

			fi, err := os.Stat(name)
			if err != nil {
				if os.IsNotExist(err) {
					if err = next(c); err != nil {
						if he, ok := err.(*echo.HTTPError); ok {
							if he.Code == http.StatusNotFound {
								return c.File("./static/index.html")
							}
						}
						return
					}
				}
				return
			}

			if fi.IsDir() {
				// トップページ('/')にアクセスしてきた時
				index := filepath.Join(name, "index.html")
				fi, err = os.Stat(index)

				if err != nil {
					if os.IsNotExist(err) {
						return next(c)
					}
					return
				}
				return c.File(index)
			}

			return c.File(name)
		}
	})

	e.Static("/", "/portfolio/server/static/")
	e.File("/defaultImage", "/portfolio/server/static/NoImage.png")

	e.Static("/images", "/portfolio/images")
	e.Use(session.Middleware(store))

	api := e.Group("/api")
	{
		api.GET("/category", router.IGetMainCategoriesHandler)
		api.GET("/category/sub", router.IGetSubCategoriesHandler)

		api.GET("/category/content/:mainID", router.IGetContentDetailListByMainCategoryHandler)
		api.GET("/category/content/sub/:subID", router.IGetContentDetailListBySubCategoryHandler)

		api.GET("/content", router.IGetContentDetailListHandler)
		api.GET("/content/:contentID", router.IGetContentDeteilHandler)

		api.GET("/tag", router.IGetTagListHandler)
		api.GET("/tag/content/:tagID", router.IGetContentDetailListByTag)

		api.GET("/generalData", router.GetAllGeneralDataHandler)
		api.GET("/generalData/:subject", router.GetGeneralDataBySubjectHandler)
	}

	e.POST("/api/edit/signup", router.SignUpHandler)
	e.POST("/api/edit/login", router.LoginHandler)
	e.GET("/api/edit/logout", router.LogoutHandler)

	edit := e.Group("/api/edit", router.CheckLogin)
	{
		edit.GET("/category", router.GetMainCategoriesHandler)
		edit.POST("/category/main", router.PostNewMainCategoryHandler)
		edit.PUT("/category/main/:mainID", router.PutMainCategoryHandler)
		edit.GET("/category/sub", router.GetSubCategoriesHandler)
		edit.POST("/category/:mainID/sub", router.PostNewSubCategoryHandler)
		edit.PUT("/category/sub/:subID", router.PutSubCategoryHandler)

		edit.GET("/category/content/:mainID", router.GetContentDetailListByMainCategoryHandler)
		edit.GET("/category/content/sub/:subID", router.GetContentDetailListBySubCategoryHandler)

		edit.GET("/content", router.GetContentDetailListHandler)
		edit.POST("/content", router.PostNewContentHandler)
		edit.GET("/content/:contentID", router.GetContentDeteilHandler)
		edit.PUT("/content/:contentID", router.PutContentHandler)

		edit.POST("/content/:contentID/tag/:tagID", router.PostNewTaggedContentHandler)
		edit.POST("/content/:contentID/subImage", router.PostNewSubImageHandler)
		edit.POST("/content/:contentID/mainImage", router.PostMainImageHandler)

		edit.GET("/tag", router.GetTagListHandler)
		edit.POST("/tag", router.PostNewTagHandler)
		edit.PUT("/tag/:tagID", router.PutTagHandler)
		edit.DELETE("/tag/:tagID", router.DeleteTagHandler)

		edit.GET("/tag/content/:tagID", router.GetContentDetailListByTag)

		edit.DELETE("/taggedContent/:taggedContentID", router.DeleteTaggedContentHanlder)

		edit.DELETE("/subImage/:subImageID", router.DeleteSubImageHandler)

		edit.POST("/generalData", router.PostNewGeneralDataHandler)
		edit.GET("/generalData", router.GetAllGeneralDataHandler)
		edit.GET("/generalData/:subject", router.GetGeneralDataBySubjectHandler)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(e.Start(":" + port))
}
