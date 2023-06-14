package service

import "github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"

// service is a struct that implements the handler.Service interface, and requires a Sorter and Repository interface.
type service struct{
	sorter Sorter
	repository Repository
}

// Sorter provides the sorting method
type Sorter interface {
	SortByOrder(courses []handler.Course) ([]handler.OrCourse, error)
}

// Repository provides the methods to comunicate with the repository.
type Repository interface {
	StoreNewUser(name string) (int, error)
	GetUser(id int) (string, error)
	StoreCourses([]handler.Course, int) ([]handler.OrCourse, error)
	CoursesInfo(id int) ([]string, []int, []string, []bool, []float32, []bool, error)
	PassCourse(id int, courseName string, score float32) error
	DeleteUser(id int) (string, error)
	UsersInfo() ([]string, []int, error)
}