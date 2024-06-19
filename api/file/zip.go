package file

import (
	"fmt"
	"go-simple-web/util"
	"time"

	"github.com/gin-gonic/gin"
)

// ZipDir  ...
//
//	@param ctx
func ZipDir(ctx *gin.Context) {
	dir := ctx.DefaultQuery("dir", "")
	if len(dir) < 1 || dir == "/" {
		util.Failure(ctx, -1, "dir is required")
		return
	}
	dest := fmt.Sprintf("%s-%d.zip", dir, time.Now().Unix())

	go func() {
		util.Zip(dir, dest)
	}()
	util.Success(ctx, map[string]interface{}{
		"dest": dest,
	})
}
