package main

import (
	middleware "Codechef-scraper-api/middle"
	"Codechef-scraper-api/routes"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {

    err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v.", err)
	}

	rateLimitRequestsStr := os.Getenv("RATE_LIMIT_REQUESTS")
	rateLimitWindowHoursStr := os.Getenv("RATE_LIMIT_WINDOW_HOURS")

	rateLimitRequests, _ := strconv.Atoi(rateLimitRequestsStr)

	rateLimitWindowHours, _ := strconv.Atoi(rateLimitWindowHoursStr)
	
	rateLimitWindow := time.Duration(rateLimitWindowHours) * time.Hour


    gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	 r.Use(gin.Recovery())
	 r.Use(gin.Logger()) 

	r.Use(middleware.FixedWindowRateLimitMiddleware(rateLimitRequests, rateLimitWindow))

	routes.UserRoutes(r)
	routes.ContestRoutes(r)
	routes.SolvedRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run HTTP server: %v", err)
	}

}






