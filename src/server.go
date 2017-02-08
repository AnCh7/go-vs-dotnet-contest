package main

import (
	"log"
	"net"

	"github.com/valyala/fasthttp"
)

func startServer() {

	server := &fasthttp.Server{
		Handler: requestsHandler,
		Name:    "chartingapi",
	}

	var err error
	listener, err := net.Listen("tcp4", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err = server.Serve(listener); err != nil {
		log.Fatalf("Error when serving incoming connections: %s", err)
	}
}

func requestsHandler(context *fasthttp.RequestCtx) {
	switch string(context.Path()) {
	case "/bars/last":
		barsLastHandler(context)
	case "/bars/between":
		barsBetweenHandler(context)
	default:
		context.Error("Unexpected path", fasthttp.StatusBadRequest)
	}
}
