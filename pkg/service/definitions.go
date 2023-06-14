package service

import "github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"

type service struct{
	sorter Sorter
	repository Repository
}

type Sorter interface {
	SortByOrder(courses []handler.Course) ([]handler.OrCourse, error)
}

type Repository interface {
	StoreNewUser(name string) (int, error)
	GetUser(id int) (string, error)
	StoreCourses([]handler.Course, int) ([]handler.OrCourse, error)
	CoursesInfo(id int) ([]string, []int, []string, []bool, []float32, []bool, error)
	PassCourse(id int, courseName string, score float32) error
	DeleteUser(id int) (string, error)
	UsersInfo() ([]string, []int, error)
}