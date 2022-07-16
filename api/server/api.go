package server

import (
	"github.com/gin-gonic/gin"
)

func RunAPI(address string) error {
	h, err := NewHandler()
	if err != nil {
		return err
	}

	return RunAPIWithHandler(address, h)
}

func RunAPIWithHandler(address string, h HandlerInterface) error {
	r := gin.Default()
	r.GET("/", h.GetOK)
	r.GET("/fixedcosts", h.GetFixedCosts)
	r.GET("/variablespends", h.GetVariableSpend)

	return r.Run()
}
