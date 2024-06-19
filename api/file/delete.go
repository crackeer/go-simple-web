package file

import (
	"go-simple-web/util"
	"os"

	"github.com/gin-gonic/gin"
)

// DeleteFile   ...
//
//	@param ctx
func DeleteFile(ctx *gin.Context) {
	file := ctx.DefaultQuery("file", "")
	if len(file) < 1 {
		util.Failure(ctx, -1, "file is required")
		return
	}
	if err := os.RemoveAll(file); err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	util.Success(ctx, nil)
}
