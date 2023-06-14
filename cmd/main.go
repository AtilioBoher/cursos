package main

import (
	"github.com/gin-gonic/gin"

	"github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"
	"github.com/AtilioBoher/ordenadoDeCursos/pkg/repository"
	"github.com/AtilioBoher/ordenadoDeCursos/pkg/service"
	"github.com/AtilioBoher/ordenadoDeCursos/pkg/utils"
)

func main() {
	sorter := utils.NewSorter()
	repository := repository.NewRepository(&sorter)
	service := service.NewService(&sorter, &repository)
	h := handler.New(&service)
	r := gin.Default()
	curso := r.Group("/curso")
	curso.POST("/usuario", h.StoreNewUser())
	curso.GET("/usuario/:id", h.GetUser())
	curso.GET("/listaUsuarios", h.UsersInfo())
	curso.PUT("/cargarCursos", h.CargarCursos())
	curso.POST("/ordenadoDeCursos", h.OrdenadoDeCursos()) // este es la consigna, el resto agregué como para practicar
	curso.GET("/listaCursos/:id", h.CoursesInfo())
	curso.PUT("/aprobarCurso", h.PassCourse())
	curso.DELETE("/borrarUsuario/:id", h.DeleteUser())
	r.Run()
}
