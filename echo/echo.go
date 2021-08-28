package echo

import (
	"github.com/hashicorp/consul/connect"
	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/consul-golang/app/server/handler"
	"log"
	"net/http"
)

func New(handler *handler.Server, consul *connect.Service) {
	e := echo.New()
	e.Add(http.MethodGet, "/handle", handler.Handle)

	//server := http.Server{
	//	Addr:      ":8080",
	//	TLSConfig: consul.ServerTLSConfig(),
	//}

	err := e.Start(":8080")
	log.Fatal(err)
}
