package database

import (
	"fmt"
	"go-simple-web/container"
	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

// Distinct
//
//	@param ctx
func Distinct(ctx *gin.Context) {
	request := &QueryRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	database := ctx.Param("database")
	table := ctx.Param("table")
	column := ctx.Param("column")

	gormDB := container.GetDatabase(database)
	if gormDB == nil {
		util.Failure(ctx, -1, "db nil")
		return
	}
	list := []string{}
	if err := gormDB.Table(table).Select(fmt.Sprintf("distinct %s as %s", column, column)).Where(request.Where).Pluck(column, &list).Error; err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	util.Success(ctx, list)
}
