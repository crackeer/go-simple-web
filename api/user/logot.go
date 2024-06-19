package user

import (
	"go-simple-web/container"
	"go-simple-web/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Logout(ctx *gin.Context) {
	domain := util.GetCookieDomain(ctx, "")
	cfg := container.GetAppConfig()
	ctx.SetCookie(cfg.LoginKey, "", -1, "/", domain, true, false)
	ctx.Redirect(http.StatusTemporaryRedirect, "/")
}
