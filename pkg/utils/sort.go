package utils

import (
	"fmt"

	"github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"
)

type sorter struct{}

func NewSorter() sorter {
	return sorter{}
}

func (s *sorter) SortByLevel(courses []handler.Course) ([][]handler.OrCourse, error) {
	var orByLevel [][]handler.OrCourse
	order := 0
	orByLevel = append(orByLevel, []handler.OrCourse{})
	for _, c1 := range courses {
		isLevelZero := true
		for _, c2 := range courses {
			if c1.Requierd == c2.Desired {
				isLevelZero = false
				break
			}
		}
		if isLevelZero {
			registered := false
			for _, v := range orByLevel[0] {
				if v.Name == c1.Requierd {
					registered = true
					break
				}
			}
			if !registered {
				orByLevel[0] = append(orByLevel[0],
					handler.OrCourse{Name: c1.Requierd, Order: order})
				order++
			}

		}
	}
	if order == 0 {
		return [][]handler.OrCourse{}, fmt.Errorf("there is no independent course (Order zero " +
			"course)")
	}

	i := 0
	remaining := len(courses)
	for 0 < remaining {
		orByLevel = append(orByLevel, []handler.OrCourse{})
		for _, c1 := range orByLevel[i] {
			for _, c2 := range courses {
				if c1.Name == c2.Requierd {
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
