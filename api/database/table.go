package database

import (
	"go-simple-web/container"
	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

// Table ...
//
//	@param ctx
func Table(ctx *gin.Context) {
	database := ctx.Param("database")
	gormDB := container.GetDatabase(database)
	if gormDB == nil {
		util.Failure(ctx, -1, "db nil")
		return
	}
	list := []map[string]interface{}{}
	if err := gormDB.Raw("show tables").Scan(&list).Error; err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	tables := []string{}
	for _, value := range list {
		for _, value := range value {
			tables = append(tables, value.(string))
		}
	}

	util.Success(ctx, tables)
}
