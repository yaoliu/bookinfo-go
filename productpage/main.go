//__author__ = "YaoYao"
//Date: 2020/8/5
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

var (
	logger          zerolog.Logger
	sessionKey      = ""
	store           = sessions.NewCookieStore([]byte(sessionKey))
	servicesDomain  = ""
	detailsHostname = "details"
	ratingsHostname = "ratings"
)

func init() {
	getservicesDomain()
	getdetailsHostname()
	getratingsHostname()
}

func getservicesDomain() {
	if os.Getenv("SERVICES_DOMAIN") != "" {
		servicesDomain = os.Getenv("SERVICES_DOMAIN")
	}
}

func getdetailsHostname() {
	if os.Getenv("DETAILS_HOSTNAME") != "" {
		servicesDomain = os.Getenv("DETAILS_HOSTNAME")
	}
}

func getratingsHostname() {
	if os.Getenv("RATINGS_HOSTNAME") != "" {
		ratingsHostname = os.Getenv("DETAILS_HOSTNAME")
	}
}



type Backend struct {
	Name     string
	Endpoint string
	ChildRen []string
}

func index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"table": "xxx"})
}

func health(c *gin.Context) {
	c.String(200, "Product page is healthy")
}

func login(c *gin.Context) {

}

func logout(c *gin.Context) {

}

func productsRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H(getProducts()))
}

func productRoute(c *gin.Context) {
	productId := c.Param("product_id")
	getProductDetails(productId)
	c.JSON(http.StatusOK, gin.H{"msg": productId})
}

func getProducts() map[string]interface{} {
	return map[string]interface{}{
		"id":              0,
		"title":           "The Comedy of Errors",
		"descriptionHtml": "<a href=\"https://en.wikipedia.org/wiki/The_Comedy_of_Errors\">Wikipedia Summary</a>: The Comedy of Errors is one of <b>William Shakespeare\\'s</b> early plays. It is his shortest and one of his most farcical comedies, with a major part of the humour coming from slapstick and mistaken identity, in addition to puns and word play.",
	}
}

func getProductDetails(productId string) {

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
	app.LoadHTMLGlob("templates/*")
	app.GET("/", index)
	app.GET("/health", health)
	app.GET("/index", index)
	app.POST("/login", login)
	app.GET("/logout", logout)
	app.GET("/api/v1/products", productsRoute)
	app.GET("/api/v1/products/:product_id", productRoute)

	if err := app.Run(":" + port); err != nil {
		logger.Err(err)
	}
}
