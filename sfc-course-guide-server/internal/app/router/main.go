package router

import (
	"fmt"
	"github.com/LuckyWindsck/sfc-course-guide/sfc-course-guide-server/internal/pkg/httphandler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os/exec"
	"runtime"
)

const (
	Scheme    = "http"
	Host      = "localhost"
	Port      = "8000"
	Authority = "//" + Host + ":" + Port
	Index     = "/"
	Search    = "/search"
	EntryLink = Scheme + ":" + Authority + Index
)

func openLink(url string) (err error) {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "windows":
		// cmd = "cmd"
		// args = []string{"/c", "start"}
		// https://stackoverflow.com/questions/39320371/how-start-web-server-to-open-page-in-browser-in-golang#answer-39324149

		// cmd = "rundll32"
		// args = []string{"url.dll,FileProtocolHandler"}
		// https://gist.github.com/hyg/9c4afcd91fe24316cbf0
		err = fmt.Errorf("unsupported platform")
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}

func Route() {
	router := gin.Default()

	// MiddleWare
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://localhost:3000",
		},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Load Template
	router.LoadHTMLGlob("web/*.tmpl")

	// Route
	router.GET(Index, httphandler.GetIndex)
	router.GET(Search, httphandler.GetSearch)

	// openLink(EntryLink)

	router.Run(":" + Port)
}
