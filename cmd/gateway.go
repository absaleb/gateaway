package main

import (
	"github.com/absaleb/common/pkg/env"
	"github.com/absaleb/gateway/internal/impl"
)

var (
	AddressHttp = env.GetEnv("ADDR", "127.0.0.1:8080")
)

func main() {
	srv := impl.Server{AddressHttp: AddressHttp}
	srv.Run()
}
