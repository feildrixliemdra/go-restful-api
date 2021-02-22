package main

import (
	"github.com/feildrixliemdra/go-restful-api/internal/bootstrap"
)

func main() {
	bootstrap.InitConfig()
	bootstrap.InitLogger()
	bootstrap.NewHTTPServer()
}
