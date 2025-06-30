package routes

import (
	scraper "Codechef-scraper-api/scrapers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SolvedRoutes(r *gin.Engine) {
    r.GET("/solved/:username", func(c *gin.Context) {
		username := c.Param("username")
        contests, err := scraper.ScrapeSolved(username)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        c.JSON(http.StatusOK, contests)
    })
}
