package assets

import (
	"os"
	"path/filepath"

	"go-simple-web/container"
	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

// SaveRequest ...
type SaveRequest struct {
	File    string `json:"file"`
	Content string `json:"content"`
}

// Login
//
//	@param ctx
func Save(ctx *gin.Context) {
	saveRequest := &SaveRequest{}
	if err := ctx.ShouldBindJSON(saveRequest); err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	cfg := container.GetAppConfig()

	if err := os.WriteFile(filepath.Join(cfg.FrontendDir, saveRequest.File), []byte(saveRequest.Content), 0644); err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	util.Success(ctx, nil)
}
