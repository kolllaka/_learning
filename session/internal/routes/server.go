package routes

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

const (
	SESSION_KEY = "secret"
)

var (
	store    *sessions.CookieStore
	AUTH_KEY string = "authenticated"
	USER_ID  string = "user_id"
)

func Setup() {
	router := gin.Default()
	store = sessions.NewCookieStore([]byte(SESSION_KEY))
	store.Options = &sessions.Options{
		HttpOnly: true,
		MaxAge:   time.Now().Hour() * 5,
		// Secure:   true, // https
	}

	// router.Use(NewMiddleware(), cors.New(cors.Config{
	// 	AllowCredentials: true,
	// 	AllowOrigins:     []string{"*"},
	// 	AllowHeaders:     []string{"Acess-Control-Allow-Origin, Content-type, Origin, Accept"},
	// }))

	router.POST("/auth/register", Register)
	router.POST("/auth/login", Login)
	router.POST("/auth/logout", Logout)
	router.GET("/auth/healthcheck", HealthCheck)

	router.GET("/user", GetUser)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
