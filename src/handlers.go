package main

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

func barsLastHandler(ctx *fasthttp.RequestCtx) {
	count := ctx.QueryArgs().GetUintOrZero("count")
	if count == 0 {
		count = 500
	}
	bars := barsLast(count)
	toJSON(ctx, bars)
}

func barsBetweenHandler(ctx *fasthttp.RequestCtx) {
	from := ctx.QueryArgs().Peek("from")
	to := ctx.QueryArgs().Peek("to")
	//TODO validate from and to
	bars := barsBetween(string(from), string(to))
	toJSON(ctx, bars)
}

func toJSON(ctx *fasthttp.RequestCtx, v interface{}) {
	ctx.SetContentType("application/json")
	if err := json.NewEncoder(ctx).Encode(v); err != nil {
		log.Fatalf("Json serialization exception: %s", err)
	}
}
