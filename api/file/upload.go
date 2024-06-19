package file

import (
	"fmt"
	"go-simple-web/util"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadFile ...
//
//	@param ctx
func UploadFile(ctx *gin.Context) {

	fileInfo, err := ctx.FormFile("file")
	if err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}

	uploadPath := ctx.DefaultQuery("dir", "")

	if len(uploadPath) < 1 {
		uploadPath = fmt.Sprintf("/tmp/tmp_upload_%d", time.Now().Unix())
		if err := os.MkdirAll(uploadPath, 0777); err != nil {
			util.Failure(ctx, -1, "mkdir tmp path error:"+err.Error())
			return
		}
	}
	localFile := filepath.Join(uploadPath, fileInfo.Filename)

	err = ctx.SaveUploadedFile(fileInfo, localFile)
	if err != nil {
		util.Failure(ctx, -1, "save upload file error:"+err.Error())
		return
	}
	util.Success(ctx, map[string]interface{}{
		"dest": localFile,
	})
}
