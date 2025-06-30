package middleware

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type FixedWindowRateLimiter struct {
	requestTracker map[string]int 
	limit          int            
	window         time.Duration  
	mx             sync.Mutex    
}

func NewFixedWindowRateLimiter(limit int, window time.Duration) *FixedWindowRateLimiter {
	return &FixedWindowRateLimiter{
		requestTracker: make(map[string]int),
		limit:          limit,
		window:         window,
		mx:             sync.Mutex{},
	}
}

func (r *FixedWindowRateLimiter) IsRequestAllowed(ipAddress string) bool {
	r.mx.Lock()
	defer r.mx.Unlock()

	count, exists := r.requestTracker[ipAddress]

	
	if !exists {
		r.requestTracker[ipAddress] = 1 
		
		go func() {
			time.Sleep(r.window) 
			r.mx.Lock()          
			defer r.mx.Unlock()
			delete(r.requestTracker, ipAddress) 
		}()
		return true 
	}
	if count >= r.limit {
		return false 
	}
	r.requestTracker[ipAddress]++
	return true 
}

func FixedWindowRateLimitMiddleware(limit int, window time.Duration) gin.HandlerFunc {
	limiter := NewFixedWindowRateLimiter(limit, window)

	return func(c *gin.Context) {
		ipAddr := c.ClientIP()

		if !limiter.IsRequestAllowed(ipAddr) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"message": fmt.Sprintf("You have exceeded your request limit of %d requests per %s.", limit, window.String()),
			})
			return 
		}

		c.Next() 
	}
}