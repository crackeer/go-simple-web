package file

import (
	"os"
	"path/filepath"

	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

// ListFile
//
//	@param ctx
func ListFile(ctx *gin.Context) {
	dir := ctx.DefaultQuery("dir", "/")
	list, err := os.ReadDir(dir)
	if err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}

	retData := []map[string]interface{}{}
	for _, file := range list {
		info, err := file.Info()
		tmp := map[string]interface{}{
			"name":   file.Name(),
			"is_dir": file.IsDir(),
			"path":   filepath.Join(dir, file.Name()),
			"size":   info.Size(),
			"modify": info.ModTime().Format("2006-01-02 15:04:05"),
		}
		if err == nil {
			tmp["size"] = info.Size()
		}
		retData = append(retData, tmp)
	}
	util.Success(ctx, retData)
}
