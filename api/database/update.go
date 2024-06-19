package database

import (
	"go-simple-web/container"
	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

// Query
//
//	@param ctx
func UpdateByID(ctx *gin.Context) {
	updateData := util.AllPostParams(ctx)
	database := ctx.Param("database")
	table := ctx.Param("table")
	id := ctx.Param("id")

	gormDB := container.GetDatabase(database)
	if gormDB == nil {
		util.Failure(ctx, -1, "db nil")
		return
	}

	affected := gormDB.Table(table).Where(map[string]interface{}{
		"id": id,
	}).Updates(updateData).RowsAffected

	util.Success(ctx, map[string]interface{}{
		"affected": affected,
	})
}
