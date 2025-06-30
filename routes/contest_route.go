package routes

import (
	scraper "Codechef-scraper-api/scrapers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ContestRoutes(r *gin.Engine) {
    r.GET("/contests/:username", func(c *gin.Context) {
		username := c.Param("username")
        contests, err := scraper.ScrapeContests(username)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        
        c.JSON(http.StatusOK, contests)
    })
}
