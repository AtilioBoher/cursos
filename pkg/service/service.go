package service

import (
	"github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"
)

func NewService(s Sorter, r Repository) service {
	return service{sorter: s, repository: r}
}

func (s *service) SortCourses(courses []handler.Course) ([]handler.OrCourse, error) {
	orByOrder, err := s.sorter.SortByOrder(courses)
	if err != nil {
		return []handler.OrCourse{}, err
	}
	return orByOrder, nil
}

func (s *service) StoreNewUser(name string) (int, error) {
	return s.repository.StoreNewUser(name)
}

func (s *service) GetUser(id int) (string, error) {
	return s.repository.GetUser(id)
}

func (s *service) StoreCourses(courses []handler.Course, id int) ([]handler.OrCourse, error) {
	return s.repository.StoreCourses(courses, id)
}

func (s *service) CoursesInfo(id int) ([]string, []int, []string, []bool, []float32, []bool,
	error) {
	return s.repository.CoursesInfo(id)
}
