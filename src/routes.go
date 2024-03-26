package main

import (
	docs "BrynGhiffar/urlshortener-be/docs"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var db = make(map[string]Redirect)

type Redirect struct {
	dest       string
	expiration int64
}

type CreateRedirectResponse struct {
	RedirectAlias string `json:"redirectAlias"`
}

// ShortenURL godoc
// @Summary Shorten URL
// @Description Generate a shortenned url
// @Tags URLShortener
// @Param shortenUrl query string true "URL to be shortenned"
// @Param alias query string false "URL alias, if left black will be randomly generated"
// @Success 200 {string} Helloworld
// @Router / [get]
func shortenUrlRoute(c *gin.Context) {
	redirectExpiration := getRedirectExpirationSecsEnv()
	targetUrl := c.Query("shortenUrl")
	alias := c.Query("alias")
	if targetUrl == "" {
		c.String(http.StatusBadRequest, "missing shortenUrl in query parameter")
		return
	}
	if alias == "" {
		alias = generateRandomString(5)
	}
	if !strings.HasPrefix(targetUrl, "http://") && !strings.HasPrefix(targetUrl, "https://") {
		targetUrl = fmt.Sprintf("http://%v", targetUrl)
	}
	// generate expiration
	var expiration int64 = expiresAfter(redirectExpiration)

	db[alias] = Redirect{dest: targetUrl, expiration: expiration}
	res := CreateRedirectResponse{RedirectAlias: alias}
	c.IndentedJSON(http.StatusOK, res)
}

func redirectRoute(c *gin.Context) {
	alias := c.Param("alias")
	redirect, ok := db[alias]
	if !ok {
		c.String(http.StatusNotFound, "Redirect does not exist")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, redirect.dest)
}

func setupRouter() *gin.Engine {
	getRedirectExpirationSecsEnv()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/", shortenUrlRoute)
	docs.SwaggerInfo.Host = getSwaggerHostnameEnv()
	docs.SwaggerInfo.Title = "URL Shortener Backend"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/:alias", redirectRoute)

	return r
}
