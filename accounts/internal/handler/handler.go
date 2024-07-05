package handler

import (
	"SSO/internal/service"
	"fmt"
	"github.com/valyala/fasthttp"
)

func BuildHandler(service *service.Service) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		fmt.Fprintf(ctx, "Hi there! RequestURI is %q", ctx.RequestURI())
	}
}
