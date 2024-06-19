package database

import (
	"go-simple-web/container"
	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

// ExecSQLRequest
type ExecSQLRequest struct {
	SQL string `json:"sql"`
}

// ExecSQL
//
//	@param ctx
func ExecSQL(ctx *gin.Context) {
	request := &ExecSQLRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	database := ctx.Param("database")
	gormDB := container.GetDatabase(database)
	if gormDB == nil {
		util.Failure(ctx, -1, "db nil")
		return
	}
	if err := gormDB.Exec(request.SQL).Error; err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	util.Success(ctx, nil)
}
