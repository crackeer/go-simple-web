package user

import (
	"go-simple-web/container"
	"go-simple-web/util"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var apiWhiteList []string = []string{
	"api/user/login", "api/user/logout", "api/user/register",
}

// CheckAPILogin
//
//	@param ctx
func Check(ctx *gin.Context) {
	for _, item := range apiWhiteList {
		if strings.HasSuffix(ctx.Request.URL.Path, item) {
			return
		}
	}
	cfg := container.GetAppConfig()
	token, err := ctx.Cookie(cfg.LoginKey)
	if err != nil {
		util.Failure(ctx, -100, "user not login")
		ctx.Abort()
		return
	}
	jwtData, err := parseJwt(token)

	if err != nil {
		util.Failure(ctx, -100, "user not login")
		return
	}
	if jwtData.ExpireAt < time.Now().Unix() {
		util.Failure(ctx, -100, "login expired")
		return
	}

	ctx.Set("User", jwtData)
}

var pageWhiteList []string = []string{
	".css", ".js", ".png", "user/login",
}

func PageCheck(ctx *gin.Context) {
	for _, item := range pageWhiteList {
		if strings.HasSuffix(ctx.Request.URL.Path, item) {
			return
		}
	}
	redirectLogin := func(ctx *gin.Context) {
		ctx.Redirect(http.StatusTemporaryRedirect, "/user/login?jump="+ctx.Request.URL.Path)
		ctx.Abort()
	}

	cfg := container.GetAppConfig()
	token, err := ctx.Cookie(cfg.LoginKey)
	if err != nil {
		redirectLogin(ctx)
		return
	}
	jwtData, err := parseJwt(token)

	if err != nil {
		redirectLogin(ctx)
		return
	}
	if jwtData.ExpireAt < time.Now().Unix() {
		redirectLogin(ctx)
		return
	}
}
