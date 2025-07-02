package main

import (
	middleware "Codechef-scraper-api/middle"
	"Codechef-scraper-api/routes"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)


func main() {

	if err := godotenv.Load(); err != nil {
        log.Printf("No .env file found or error loading it: %v. Using environment variables.", err)
    }

	rateLimitRequestsStr := os.Getenv("RATE_LIMIT_REQUESTS")
	rateLimitWindowHoursStr := os.Getenv("RATE_LIMIT_WINDOW_HOURS")
	port := os.Getenv("PORT")

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

	Addr := fmt.Sprintf(":%s", port)
	if err := r.Run(Addr); err != nil { 
		log.Fatalf("Failed to run HTTP server: %v", err)
	}

}






