package http

import (
	"strings"
	"time"

	"go-simple-web/util"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
)

// HandleHttpProxy
//
//	@param ctx
func HandleHttpProxy(ctx *gin.Context) {
	headers := map[string]string{}
	for key, value := range ctx.Request.Header {
		if strings.HasPrefix(key, "X-Proxy-") {
			headers[strings.TrimPrefix(key, "X-Proxy-")] = strings.Join(value, "")
		}
	}
	if _, ok := headers["Content-Type"]; !ok {
		if value, ok := ctx.Request.Header["Content-Type"]; ok {
			headers["Content-Type"] = strings.Join(value, "")
		}
	}

	var schema string = "http"
	if value, ok := headers["X-Config-Schema"]; ok {
		schema = value
	}

	baseURL := schema + "://" + ctx.Param("urlpath")
	request := resty.New().SetTimeout(30 * time.Second).SetBaseURL(baseURL).R().SetHeaders(headers)

	if body, err := ctx.GetRawData(); err == nil {
		request = request.SetBody(body)

	}
	if values := ctx.Request.URL.Query(); len(values) > 0 {
		for k, v := range values {
			request = request.SetQueryParam(k, strings.Join(v, ","))
		}
	}
	response, err := request.Execute(ctx.Request.Method, "")
	if err != nil {
		util.Failure(ctx, -2, err.Error())
		return
	}

	util.Success(ctx, response.String())
}
