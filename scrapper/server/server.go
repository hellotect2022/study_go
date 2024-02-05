package server

import (
	"fmt"
	"net/http"
	"newProj/scrapper_goroutine"
	"os"
	"strings"

	"github.com/labstack/echo"
)

func EchoStart() {
	e := echo.New()
	e.GET("/main", handleHome)
	e.POST("/scrape", handleScrape)
	e.GET("/html", handleHtml)
	e.GET("/sample", handleSample)
	e.GET("/sample_image", handleImage)
	e.Logger.Fatal(e.Start(":5000"))
}

func handleHome(ctx echo.Context) error {
	return ctx.File("D:/Go/go/src/newProj/home.html")
	//return ctx.File("/home.html")
}
func handleScrape(ctx echo.Context) error {
	defer os.Remove("./jobs.csv")
	term := strings.ToLower(ctx.FormValue("term"))
	fmt.Println("검색값 : " + term)
	scrapper_goroutine.ScrapperMain(term)
	return ctx.Attachment("jobs.csv", fmt.Sprintf("job-search-result-%s.csv", term))
}

func handleHtml(ctx echo.Context) error {
	return ctx.HTML(http.StatusOK, "<h1>test</h1>")
}

func handleSample(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Hello World!")
}

func handleImage(ctx echo.Context) error {
	//return ctx.File("D:/Go/go/src/newProj/main_dmg.png")
	return ctx.File("/main_dmg.png")
}
