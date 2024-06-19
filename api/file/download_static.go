package file

import (
	"go-simple-web/util"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// DownloadStatic    ...
//
//	@param ctx
func DownloadStatic(ctx *gin.Context) {
	urlStr := ctx.DefaultQuery("url", "")
	rootDir := ctx.DefaultQuery("root_dir", "/data1/vrdata/vrlab-public-1")
	if len(urlStr) < 1 {
		util.Failure(ctx, -1, "url is required")
		return
	}
	object, err := url.Parse(urlStr)
	if err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	dest := filepath.Join(rootDir, object.Path)

	if err := downloadTo(urlStr, dest); err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}

	util.Success(ctx, map[string]interface{}{
		"dest": dest,
	})
}

func downloadTo(urlString string, target string) error {
	dir, _ := filepath.Split(target)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	response, err := http.Get(urlString)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	return os.WriteFile(target, bytes, os.ModePerm)
}
