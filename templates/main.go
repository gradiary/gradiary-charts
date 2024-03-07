package main

import (
	"flag"
	"gradiary/gradiary-api-service/api"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Auto-Generated: https://www.atatus.com/blog/golang-auto-build-versioning/
// go build -ldflags "-X main.Revision=`date -u +%Y.%-m.%-d`.`git rev-list --count HEAD` -X main.Commit=`git rev-parse --short HEAD`"
var (
	Revision = ""
	Commit   = ""
)

func main() {

	// Just a workaround for glog flag parsing: https://github.com/kubernetes/kubernetes/issues/17162
	flag.CommandLine.Parse([]string{})

	log.Println("gradiary-api-service started.....")

	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.BodyLimit("64M"))
	e.Use(middleware.Recover())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}", "method":"${method}", "uri":"${uri}", "status":${status}, "took":"${latency_human}","bytes_in":${bytes_in},"bytes_out":${bytes_out}, "error":"${error}", "id":"${id}"}` + "\n",
	}))

	suffixPath := os.Getenv("API_PATH")

	// Public group - anybody can call
	p := e.Group(suffixPath)

	// API version info
	p.GET("/version", api.GetVersion)

	e.Logger.Fatal(e.Start(":80"))
}
