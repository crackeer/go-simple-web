package database

import (
	"go-simple-web/container"
	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

// Query
//
//	@param ctx
func Create(ctx *gin.Context) {
	data := util.AllPostParams(ctx)
	database := ctx.Param("database")
	table := ctx.Param("table")

	gormDB := container.GetDatabase(database)
	if gormDB == nil {
		util.Failure(ctx, -1, "db nil")
		return
	}

	if err := gormDB.Table(table).Create(&data).Error; err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	util.Success(ctx, data)
}
