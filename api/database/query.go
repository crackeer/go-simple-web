package database

import (
	"go-simple-web/container"
	"go-simple-web/util"

	"github.com/gin-gonic/gin"
)

const DefaultPageSize = 10

// QueryRequest
type QueryRequest struct {
	Where    interface{} `json:"where"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

// Query
//
//	@param ctx
func Query(ctx *gin.Context) {
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
	if request.Page < 1 {
		request.Page = 1
	}
	if request.PageSize < 1 {
		request.PageSize = DefaultPageSize
	}
	list := []map[string]interface{}{}
	if err := gormDB.Table(table).Where(request.Where).Offset((request.Page - 1) * request.PageSize).Limit(request.PageSize).Order("id desc").Find(&list).Error; err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	util.Success(ctx, list)
}
