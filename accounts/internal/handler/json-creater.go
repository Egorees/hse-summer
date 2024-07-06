package handler

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"log/slog"
	"net/http"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json")
)

func doJSONWrite(ctx *fasthttp.RequestCtx, code int, obj interface{}) {
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(code)
	if err := json.NewEncoder(ctx).Encode(obj); err != nil {
		slog.Error("Error while creating json", err)
		ctx.Error(err.Error(), http.StatusInternalServerError)
	}
}
