package file

import (
	"go-simple-web/util"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// DownloadFile
//
//	@param ctx
func DownloadFile(ctx *gin.Context) {
	filePath := ctx.DefaultQuery("file", "")
	if len(filePath) < 1 {
		util.Failure(ctx, -1, "file path nil")
		return
	}
	_, name := filepath.Split(filePath)
	ctx.Abort()
	ctx.FileAttachment(filePath, name)
}
