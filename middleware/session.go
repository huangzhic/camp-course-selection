package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

// Session 初始化session
func Session(secret string) gin.HandlerFunc {
	//store := cookie.NewStore([]byte(secret))
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(secret))
	//Also set Secure: true if using SSL, you should though
	store.Options(sessions.Options{HttpOnly: true, MaxAge: 900, Path: "/"})
	return sessions.Sessions("camp-seesion", store)
}
