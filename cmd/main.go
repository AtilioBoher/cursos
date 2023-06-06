package main

import (
	"github.com/gin-gonic/gin"

	"github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"
	"github.com/AtilioBoher/ordenadoDeCursos/pkg/service"
	"github.com/AtilioBoher/ordenadoDeCursos/pkg/utils"
)

func main() {
	sorter := utils.NewSorter()
	service := service.NewService(&sorter)
	h := handler.New(&service)
	r := gin.Default()
	r.GET("/ordenadoDeCursos", h.OrdenadoDeCursos())
	r.Run()
}
