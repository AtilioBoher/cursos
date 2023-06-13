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
	for _, u := range r.Users {
		if u.Name == name {
			return 0, fmt.Errorf("user with name: %s already exist, it has id: %d", name, u.Id)
		}
	}
	r.LastId++
	r.Users = append(r.Users, User{
		Name:           name,
		Id:             r.LastId,
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

func (r *repo) PassCourse(id int, courseName string, score float32) error {
	u, ok := r.IdUserLookUp[id]
	if !ok {
		return fmt.Errorf("user with id: %v not found", id)
	}
	return u.passCourse(courseName, score)
}
