package database

import (
	"go-simple-web/container"
	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

// Desc  ...
//
//	@param ctx
func Desc(ctx *gin.Context) {
	database := ctx.Param("database")
	table := ctx.Param("table")
	gormDB := container.GetDatabase(database)
	if gormDB == nil {
		util.Failure(ctx, -1, "db nil")
		return
	}
	fields := []map[string]interface{}{}
	if err := gormDB.Raw("desc " + table).Find(&fields).Error; err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	util.Success(ctx, fields)
}
