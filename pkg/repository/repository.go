package repository

import (
	"fmt"

	"github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"
)

func NewRepository(s Sorter) repo {
	return repo{
		Users:        []User{},
		LastId:       0,
		IdUserLookUp: make(map[int]*User),
		sorter:       s,
	}
}

func (r *repo) StoreNewUser(name string) (int, error) {
	r.LastId++
	r.Users = append(r.Users, User{
		Name:           name,
		Id:             r.LastId,
		IdCourseLookUp: make(map[string]*Course),
	})
	r.IdUserLookUp[r.LastId] = &r.Users[len(r.Users)-1]
	return r.LastId, nil
}

func (r *repo) GetUser(id int) (string, error) {
	u, ok := r.IdUserLookUp[id]
	if ok {
		return u.Name, nil
	}
	return "", fmt.Errorf("user with id: %v not found", id)
}

func (r *repo) StoreCourses(courses []handler.Course, id int) ([]handler.OrCourse, error) {
	u, ok := r.IdUserLookUp[id]
	if !ok {
		return []handler.OrCourse{}, fmt.Errorf("user with id: %v not found", id)
	}
	// clear previous courses
	u.Courses = []Course{}
	u.State = CoursesState{}
	u.IdCourseLookUp = make(map[string]*Course)
	// load courses
	orCourses, err := r.sorter.StoreCourses(courses, u)
	return orCourses, err
}

func (r *repo) CoursesInfo(id int) ([]string, []int, []string, []bool, []float32, []bool, error) {
	u, ok := r.IdUserLookUp[id]
	if !ok {
		return []string{}, []int{}, []string{}, []bool{}, []float32{}, []bool{},
		fmt.Errorf("user with id: %v not found", id)
	}
	courseName, order, reqCourseName, passed, score, avalilable := u.CoursesInfo()
	return courseName, order, reqCourseName, passed, score, avalilable, nil
}