package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maiaaraujo5/gostart/log/logger"
)

type Server struct {
}

func NewCreateServer() *Server {
	return &Server{}
}

func (h Server) Handle(c echo.Context) error {
	logger.Info("Request Received")

	return c.JSON(http.StatusOK, nil)
}
