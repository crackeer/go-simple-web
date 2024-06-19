package main

import (
	"encoding/json"
	"fmt"
	"go-simple-web/api/database"
	"go-simple-web/api/file"
	httpHandler "go-simple-web/api/http"
	"go-simple-web/api/user"
	"go-simple-web/container"
	"go-simple-web/util"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	rollRender "github.com/unrolled/render"
)

const tmplExtension = ".tmpl"
const tmplDataExtension = ".json"

var (
	renderer   *rollRender.Render
	fileServer http.Handler
)

func initRender(dir, layout string) {
	renderer = rollRender.New(rollRender.Options{
		Directory:  dir,                           // Specify what path to load the templates from.
		FileSystem: &rollRender.LocalFileSystem{}, // Specify filesystem from where files are loaded.
		Layout:     layout,                        // Specify a layout template. Layouts can call {{ yield }} to render the current template or {{ partial "css" }} to render a partial from the current template.
		Extensions: []string{".tmpl"},             // Specify extensions to load for templates.
		Delims: rollRender.Delims{
			Left:  "{[{",
			Right: "}]}",
		},
		IsDevelopment:               true,
		Asset:                       nil,
		AssetNames:                  nil,
		RenderPartialsWithoutPrefix: true,
	})

	fileServer = http.StripPrefix("", http.FileServer(http.Dir(dir)))
}

func main() {
	if err := container.Initialize(); err != nil {
		panic(err)
	}
	cfg := container.GetAppConfig()
	router := gin.Default()
	router.Any("/proxy/*urlpath", util.DoResponseJSON(), httpHandler.HandleHttpProxy)

	apiRouter := router.Group("api", user.Check, util.DoResponseJSON())
	apiRouter.POST("database/query/:database/:table", database.Query)
	apiRouter.POST("database/count/:database/:table", database.Count)
	apiRouter.POST("database/create/:database/:table", database.Create)
	apiRouter.POST("database/desc/:database/:table", database.Desc)
	apiRouter.POST("database/exec_sql/:database", database.ExecSQL)
	apiRouter.POST("database/table/:database", database.Table)
	apiRouter.POST("database/distinct/:database/:table/:column", database.Distinct)
	apiRouter.POST("database/single_update/:database/:table/:id", database.UpdateByID)
	apiRouter.POST("database/single_delete/:database/:table/:id", database.DeleteByID)
	apiRouter.POST("user/login", user.Login)

	apiRouter.GET("file/list", file.ListFile)
	apiRouter.GET("file/download", file.DownloadFile)
	apiRouter.POST("file/upload", file.UploadFile)
	apiRouter.GET("file/zip", file.ZipDir)
	apiRouter.GET("file/delete", file.DeleteFile)
	apiRouter.GET("file/mkdirall", file.MkdirAll)
	apiRouter.GET("file/download_static", file.DownloadStatic)

	router.Use(user.PageCheck)
	initRender(cfg.FrontendDir, "layout")
	router.NoRoute(createStaticHandler(cfg.FrontendDir))
	router.Run(fmt.Sprintf(":%d", cfg.Port))
}

func createStaticHandler(dir string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file := strings.TrimLeft(ctx.Request.URL.Path, "/")
		if len(file) < 1 {
			file = "index"
		}

		if isMarkdown(file) {
			file = "markdown"
		}

		fullPath := filepath.Join(dir, file)
		if fileExists(fullPath) {
			fileServer.ServeHTTP(ctx.Writer, ctx.Request)
			return
		}
		if !fileExists(fullPath + tmplExtension) {
			ctx.String(http.StatusNotFound, "page not found")
			ctx.Abort()
			return
		}
		data := map[string]interface{}{}
		if fileExists(fullPath + tmplDataExtension) {
			if bytes, err := os.ReadFile(fullPath + tmplDataExtension); err == nil {
				json.Unmarshal(bytes, &data)
			}
		}
		renderer.HTML(ctx.Writer, http.StatusOK, file, data)
	}
}

func isMarkdown(value string) bool {
	return strings.HasSuffix(value, ".mdx")
}

func fileExists(filename string) bool {
	if value, err := os.Stat(filename); err == nil && !value.IsDir() {
		return true
	} else if os.IsNotExist(err) {
		return false
	}
	return false
}
