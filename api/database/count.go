package database

import (
	"go-simple-web/container"
	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

// Query
//
//	@param ctx
func Count(ctx *gin.Context) {
	request := &QueryRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	database := ctx.Param("database")
	table := ctx.Param("table")

	gormDB := container.GetDatabase(database)
	if gormDB == nil {
		util.Failure(ctx, -1, "db nil")
		return
	}
	var total int64
	gormDB.Table(table).Where(request.Where).Count(&total)

	util.Success(ctx, map[string]interface{}{
		"total": total,
	})
}
