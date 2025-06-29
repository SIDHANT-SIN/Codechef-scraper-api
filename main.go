package main



import (
    "github.com/gin-gonic/gin"
    "Codechef-scraper-api/routes"
)

func main() {
    r := gin.Default()

   routes.UserRoutes(r)
   routes.ContestRoutes(r)
    routes.SolvedRoutes(r)


    r.Run(":8080")
}


