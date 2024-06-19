package api

import (
	"context"
	"fmt"
	"go-simple-web/util"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

type ExecRequest struct {
	Params []string `json:"params"`
}

// Exec
//
//	@param ctx
func Exec(ctx *gin.Context) {
	name := ctx.Param("name")
	request := &ExecRequest{}
	if err := ctx.ShouldBindJSON(request); err != nil {
		util.Failure(ctx, -1, err.Error())
		return
	}
	fmt.Println(name, request.Params)
	cmd := exec.CommandContext(context.Background(), name, request.Params...)
	cmd.Stderr = os.Stderr
	result, err := cmd.Output()
	if err != nil {
		util.Failure(ctx, -2, err.Error())
		return
	}
	util.Success(ctx, map[string]string{
		"output": string(result),
	})
}
