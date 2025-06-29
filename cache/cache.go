package cache

import (
    "time"
    "github.com/patrickmn/go-cache"
)

var UserCache = cache.New(10*time.Minute, 15*time.Minute)
var DefaultExpiration = 10 * time.Minute