package file

import (
	"go-simple-web/util"
	"os"

	"github.com/gin-gonic/gin"
)

// DeleteFile   ...
//
//	@param ctx
func MkdirAll(ctx *gin.Context) {
	file := ctx.DefaultQuery("dir", "")
	if len(file) < 1 {
		util.Failure(ctx, -1, "dur is required")
		return
	}
	if err := os.MkdirAll(file, os.ModePerm); err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	util.Success(ctx, nil)
}
