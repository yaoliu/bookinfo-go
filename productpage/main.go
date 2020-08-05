//__author__ = "YaoYao"
//Date: 2020/8/5
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

var (
	logger zerolog.Logger
)

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "", gin.H{"1": ""})
}

func health(c *gin.Context) {
	c.String(200, "Product page is healthy")
}

func login(c *gin.Context) {

}

func logout(c *gin.Context) {

}

func main() {
	logger = zerolog.New(os.Stdout).With().Timestamp().Logger()

	args := os.Args

	if len(args) < 2 {
		logger.Error().Msgf("usage: %s port", args[0])
		os.Exit(-1)
	}

	port := args[1]

	logger.Info().Msgf("start at port %s", port)

	app := gin.Default()
	app.Static("/static", "static")
	app.GET("/health", health)
	app.GET("/index", index)
	app.POST("/login", login)
	app.GET("/logout", logout)

	if err := app.Run(":" + port); err != nil {
		logger.Err(err)
	}
}
