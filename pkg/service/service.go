package service

import (
	"github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"
)

func NewService(s Sorter) service {
	return service{sorter: s}
}

func (s *service) SortCourses(courses []handler.Course) ([]handler.OrCourse, error) {
	orByOrder, err := s.sorter.SortByOrder(courses)
	if err != nil {
		return []handler.OrCourse{}, err
	}
	return orByOrder, nil
}
