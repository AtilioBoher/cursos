package service

import (
	"github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"
)

// NewService returns an instance of a service struct. To work properly, it needs a struct that
// implements the Sorter interface and a struct that implements the Repository interface.
func NewService(s Sorter, r Repository) service {
	return service{sorter: s, repository: r}
}

// SortCourses receives a list of courses and sorts them, returning the ordered courses.
func (s *service) SortCourses(courses []handler.Course) ([]handler.OrCourse, error) {
	orByOrder, err := s.sorter.SortByOrder(courses)
	if err != nil {
		return []handler.OrCourse{}, err
	}
	return orByOrder, nil
}

// StoreNewUser stores a user with the name supplied and returns the id generated.
func (s *service) StoreNewUser(name string) (int, error) {
	return s.repository.StoreNewUser(name)
}

// GetUser returns the name of the user with the id supplied.
func (s *service) GetUser(id int) (string, error) {
	return s.repository.GetUser(id)
}

// StoreCourses stores the courses supplied for the user with the id speccified, and returns the
// ordered courses.
func (s *service) StoreCourses(courses []handler.Course, id int) ([]handler.OrCourse, error) {
	return s.repository.StoreCourses(courses, id)
}

// CoursesInfo returns slices with the information of all the courses of the user with the id supplied.
func (s *service) CoursesInfo(id int) ([]string, []int, []string, []bool, []float32, []bool,
	error) {
	return s.repository.CoursesInfo(id)
}

// PassCourse approves the course of the name indicated with the score supplied.
func (s *service) PassCourse(userId int, courseName string, score float32) error {
	return s.repository.PassCourse(userId, courseName, score)
}

// DeleteUser deletes the user with the id given, and returns his name.
func (s *service) DeleteUser(id int) (string, error) {
	return s.repository.DeleteUser(id)
}

// UserInfo returns the information from all the users.
func (s *service) UsersInfo() ([]string, []int, error) {
	return s.repository.UsersInfo()
}
