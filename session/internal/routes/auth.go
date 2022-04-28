package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/KoLLlaka/_learning/session/internal/model"
)

func NewMiddleware() gin.HandlerFunc {
	return AuthMiddleware
}

func AuthMiddleware(ctx *gin.Context) {
	fmt.Printf("strings.Split(ctx.FullPath(), \"/\")[1]: %v\n", strings.Split(ctx.FullPath(), "/")[1])

	if strings.Split(ctx.FullPath(), "/")[1] == "auth" {
		ctx.Next()
	}

	sess, err := store.Get(ctx.Request, "session-name")
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "not authorized",
		})
	}

	if sess.Values[AUTH_KEY] == nil {
		ctx.JSON(http.StatusUnauthorized, map[string]string{
			"message": "not authorized",
		})
	}

	ctx.Next()
}

func Register(ctx *gin.Context) {
	var user *model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("!!!!!!!!!!!!!!!User:", user)

	// password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	// if err != nil {
	// 	ctx.String(http.StatusInternalServerError, err.Error())
	// }

	// user.Password = password
	// fmt.Printf("!!!!!!!!!!!!user.Password: %v\n", user.Password, password)

	if err := model.CreateUser(user); err != nil {
		ctx.String(http.StatusInternalServerError, err.Error())
	}

	ctx.String(http.StatusOK, "registered")
}
func Login(ctx *gin.Context) {
	ctx.String(http.StatusOK, "not implement")
}
func Logout(ctx *gin.Context) {
	ctx.String(http.StatusOK, "not implement")
}
func HealthCheck(ctx *gin.Context) {
	ctx.String(http.StatusOK, "not implement")
}
func GetUser(ctx *gin.Context) {
	ctx.String(http.StatusOK, "not implement")
}
