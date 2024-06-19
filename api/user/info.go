package user

import (
	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

func Info(ctx *gin.Context) {
	util.Success(ctx, getCurrentUser(ctx))
}

func getCurrentUser(ctx *gin.Context) *User {
	value, exists := ctx.Get("CurrentUser")
	if !exists {
		return nil
	}

	if user, ok := value.(*User); ok {
		return user
	}
	return nil
}
