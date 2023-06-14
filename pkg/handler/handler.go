// myHdlrPkg (My Handler Package) is a package to handle the request...
package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func New(s Service) myHandler {
	return myHandler{service: s}
}

func (hdlr *myHandler) OrdenadoDeCursos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := oReq{}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := checkRequestData(req)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		orCourses, err := hdlr.service.SortCourses(req.Courses)
		if err != nil {
			ctx.JSON(422, gin.H{"error": err.Error()})
			return
		}
		res := oRes{
			UserId:  req.UserId,
			Courses: orCourses,
		}
		ctx.JSON(http.StatusAccepted, res)
	}
}

func checkRequestData(r oReq) error {
	if r.UserId == 0 {
		return fmt.Errorf("es necesario especificar el ID del usuario")
	}
	for _, c := range r.Courses {
		if c.Desired == "" {
			return fmt.Errorf("falta especificar el nombre de uno de los cursos")
		}
		if c.Required == "" {
			return fmt.Errorf("el curso %s no especifica el curso requerido", c.Desired)
		}
	}
	for _, c1 := range r.Courses {
		i := 0
		for _, c2 := range r.Courses {
			if c1.Desired == c2.Desired {
				i++
			}
			if i > 1 {
				return fmt.Errorf("el curso %s se encuentra duplicado", c1.Desired)
			}
		}
	}
	return nil
}

func (hdlr *myHandler) StoreNewUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := User{}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if req.Name == "" {
			ctx.JSON(404, gin.H{"error": "no se especifica el nombre del usuario"})
			return
		}
		id, err := hdlr.service.StoreNewUser(req.Name)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		req.Id = id
		ctx.JSON(http.StatusAccepted, req)
	}
}

func (hdlr *myHandler) GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, gin.H{"error": ("invalid ID" + err.Error())})
			return
		}
		name, err := hdlr.service.GetUser(int(id))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		res := User{
			Name: name,
			Id:   int(id),
		}
		ctx.JSON(http.StatusAccepted, res)
	}
}

func (hdlr *myHandler) CargarCursos() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := oReq{}
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := checkRequestData(req); err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		orCourses, err := hdlr.service.StoreCourses(req.Courses, req.UserId)
		if err != nil {
			ctx.JSON(422, gin.H{"error": err.Error()})
			return
		}
		res := oRes{
			UserId:  req.UserId,
			Courses: orCourses,
		}
		ctx.JSON(http.StatusAccepted, res)
	}
}

func (hdlr *myHandler) CoursesInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, gin.H{"error": ("invalid ID" + err.Error())})
			return
		}
		courseName, order, reqCourseName, passed, score, avalilable,
		err := hdlr.service.CoursesInfo(int(id))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		courses := []CourseInfo{}
		for i := range courseName {
			courses = append(courses, CourseInfo{
				OrCourse:  OrCourse{
					Name:  courseName[i],
					Order: order[i],
				},
				Required:  reqCourseName[i],
				Passed:    passed[i],
				Score:     score[i],
				Available: avalilable[i],
			})
		}
		res := InfoResponse{
			UserId:  int(id),
			Courses: courses,
		}
		ctx.JSON(http.StatusAccepted, res)
	}
}

func (hdlr *myHandler) PassCourse() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Query("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, gin.H{"error": ("invalid ID" + err.Error())})
			return
		}
		name := ctx.Query("name")
		if name == "" {
			ctx.JSON(404, gin.H{"error": "name required"})
			return
		}
		score, err := strconv.ParseFloat(ctx.Query("score"), 32)
		if err != nil {
			ctx.JSON(404, gin.H{"error": ("invalid score" + err.Error())})
			return
		}
		if err := hdlr.service.PassCourse(int(id),name, float32(score)); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"messege": fmt.Sprintf("course %s passed with score: %f",
		name, score)})
	}
}

func (hdlr *myHandler) DeleteUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(404, gin.H{"error": ("invalid ID" + err.Error())})
			return
		}
		name, err := hdlr.service.DeleteUser(int(id));
		if  err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusAccepted, gin.H{"messege": fmt.Sprintf("user %s, with id: %d was " + 
		"successfully deleted",
		name, id)})
	}
}

func (hdlr *myHandler) UsersInfo() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		names, ids, err := hdlr.service.UsersInfo()
		if  err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var users []User
		for i := range names {
			users = append(users, User{
				Name: names[i],
				Id:   ids[i],
			})
		}
		ctx.JSON(http.StatusAccepted, users)
	}
}