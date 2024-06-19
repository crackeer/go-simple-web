package database

import (
	"fmt"
	"go-simple-web/container"
	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

// Query
//
//	@param ctx
func DeleteByID(ctx *gin.Context) {
	database := ctx.Param("database")
	table := ctx.Param("table")
	id := ctx.Param("id")

	gormDB := container.GetDatabase(database)
	if gormDB == nil {
		util.Failure(ctx, -1, "db nil")
		return
	}

	if err := gormDB.Exec(fmt.Sprintf("delete from %s where id = %s", table, id)).Error; err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}

	util.Success(ctx, nil)
}
