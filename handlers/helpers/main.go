package helpers

import (
	"github.com/pquerna/ffjson/ffjson"
	"github.com/valyala/fasthttp"
)

func WriteResponseJson(ctx *fasthttp.RequestCtx, obj interface{}) {
	buf, _ := ffjson.Marshal(obj)
	WriteResponse(ctx, buf)
	ffjson.Pool(buf)
}

func WriteResponse(ctx *fasthttp.RequestCtx, contents []byte) {
	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Response.Header.SetContentLength(len(contents))
	ctx.Write(contents)
}
