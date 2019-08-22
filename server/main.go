package main

import(
	"net/http"
	"os"
	"log"

	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()
	
	e.GET("/", func(c echo.Context) error{
		return c.String(http.StatusOK, "hello")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(e.Start(":" + port))
}