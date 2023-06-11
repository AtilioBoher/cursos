package utils

import (
	"fmt"

	"github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"
	"github.com/AtilioBoher/ordenadoDeCursos/pkg/repository"
)

type sorter struct{}

func NewSorter() sorter {
	return sorter{}
}

func (s *sorter) SortByLevel(courses []handler.Course) ([][]handler.OrCourse, error) {
	var orByLevel [][]handler.OrCourse
	order := 0
	orByLevel = append(orByLevel, []handler.OrCourse{})
	// check and register as level zero all of the courses that have no required courses
	for _, c1 := range courses {
		isLevelZero := true
		for _, c2 := range courses {
			if c1.Required == c2.Desired {
				isLevelZero = false
				break
			}
		}
		if isLevelZero {
			registered := false
			for _, v := range orByLevel[0] {
				if v.Name == c1.Required {
					registered = true
					break
				}
			}
			if !registered {
				orByLevel[0] = append(orByLevel[0],
					handler.OrCourse{Name: c1.Required, Order: order})
				order++
			}
		}
	}
	if order == 0 {
		return [][]handler.OrCourse{}, fmt.Errorf("there is no independent course (zero order " +
			"course)")
	}
	i := 0
	remaining := len(courses)
	for 0 < remaining {
		orByLevel = append(orByLevel, []handler.OrCourse{})
		for _, c1 := range orByLevel[i] {
			for _, c2 := range courses {
				if c1.Name == c2.Required {
					orByLevel[i+1] = append(orByLevel[i+1],
						handler.OrCourse{Name: c2.Desired, Order: order})
					order++
					remaining--
				}
			}
		}
		if len(orByLevel[i+1]) == 0 {
			return [][]handler.OrCourse{}, fmt.Errorf("there is a circular dependency among " +
				"courses")
		}
		i++
	}
	return orByLevel, nil
}

func (s *sorter) SortByOrder(courses []handler.Course) ([]handler.OrCourse, error) {
	orByLevel, err := s.SortByLevel(courses)
	if err != nil {
		return []handler.OrCourse{}, err
	}
	var orByOrder []handler.OrCourse
	for _, level := range orByLevel {
		orByOrder = append(orByOrder, level...)
	}
	return orByOrder, nil

}

func (s *sorter) StoreCourses(courses []handler.Course, u *repository.User) ([]handler.OrCourse, error) {
	var orByLevel [][]handler.OrCourse
	order := 0
	orByLevel = append(orByLevel, []handler.OrCourse{})
	// check and register as level zero all of the courses that have no required courses
	for _, c1 := range courses {
		isLevelZero := true
		for _, c2 := range courses {
			if c1.Required == c2.Desired {
				isLevelZero = false
				break
			}
		}
		if isLevelZero {
			registered := false
			for _, v := range orByLevel[0] {
				if v.Name == c1.Required {
					registered = true
					break
				}
			}
			if !registered {
				orByLevel[0] = append(orByLevel[0],
					handler.OrCourse{Name: c1.Required, Order: order})
					u.Courses = append(u.Courses,
					repository.Course{
						Name:      c1.Required,
						Order:     order,
						Required:  nil,
						Passed:    false,
						Score:     0,
						Available: true,
					})
					u.State.AvailableCourses = append(u.State.AvailableCourses, &u.Courses[order])
					u.IdCourseLookUp[u.Courses[order].Name] = &u.Courses[order]
				order++
			}
		}
	}
	if order == 0 {
		return []handler.OrCourse{}, fmt.Errorf("there is no independent course (zero order course)")
	}
	i := 0
	remaining := len(courses)
	for 0 < remaining {
		orByLevel = append(orByLevel, []handler.OrCourse{})
		for _, c1 := range orByLevel[i] {
			for _, c2 := range courses {
				if c1.Name == c2.Required {
					orByLevel[i+1] = append(orByLevel[i+1],
						handler.OrCourse{Name: c2.Desired, Order: order})
						u.Courses = append(u.Courses,
							repository.Course{
								Name:      c2.Desired,
								Order:     order,
								Required:  u.IdCourseLookUp[c1.Name],
								Passed:    false,
								Score:     0,
								Available: false,
							})
							u.State.NotAvailableCourses = append(u.State.NotAvailableCourses, &u.Courses[order])
							u.IdCourseLookUp[u.Courses[order].Name] = &u.Courses[order]
					order++
					remaining--
				}
			}
		}
		if len(orByLevel[i+1]) == 0 {
			return []handler.OrCourse{}, fmt.Errorf("there is a circular dependency among courses")
		}
		i++
	}
	var orByOrder []handler.OrCourse
	for _, level := range orByLevel {
		orByOrder = append(orByOrder, level...)
	}
	return orByOrder, nil
}
