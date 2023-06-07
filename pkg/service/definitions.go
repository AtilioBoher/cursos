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
}