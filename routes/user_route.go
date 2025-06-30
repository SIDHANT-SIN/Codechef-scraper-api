package routes

import (
	"Codechef-scraper-api/cache"
	scraper "Codechef-scraper-api/scrapers"
	"net/http"

	"github.com/gin-gonic/gin"
)



func UserRoutes(r *gin.Engine) {
    r.GET("/user/:username", func(c *gin.Context) {
        username := c.Param("username")

       
        cached, found := cache.UserCache.Get(username)
        if found {
            c.JSON(http.StatusOK, cached)
            return
        }

        user, err := scraper.ScrapeUser(username)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        cache.UserCache.Set(username, user, cache.DefaultExpiration)

        c.JSON(http.StatusOK, user)
    })
}
