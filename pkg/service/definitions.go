package service

import "github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"

type service struct{
	sorter Sorter
}

type Sorter interface {
	SortByOrder(courses []handler.Course) ([]handler.OrCourse, error)
}
