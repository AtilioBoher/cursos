package repository

import (
	"fmt"

	"github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"
)

// NewRepository returns a repository struct.
func NewRepository(s Sorter) repo {
	return repo{
		Users:        []User{},
		LastId:       0,
		IdUserLookUp: make(map[int]*User),
		sorter:       s,
	}
}

// StoreNewUser stores a user with the name supplied and returns the id generated.
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

// GetUser returns the name of the user with the id supplied.
func (r *repo) GetUser(id int) (string, error) {
	u, ok := r.IdUserLookUp[id]
	if ok {
		return u.Name, nil
	}
	return "", fmt.Errorf("user with id: %v not found", id)
}

// StoreCourses stores the courses supplied for the user with the id speccified, and returns the
// ordered courses.
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

// CoursesInfo returns slices with the information of all the courses of the user with the id supplied.
func (r *repo) CoursesInfo(id int) ([]string, []int, []string, []bool, []float32, []bool, error) {
	u, ok := r.IdUserLookUp[id]
	if !ok {
		return []string{}, []int{}, []string{}, []bool{}, []float32{}, []bool{},
			fmt.Errorf("user with id: %v not found", id)
	}
	courseName, order, reqCourseName, passed, score, avalilable := u.CoursesInfo()
	return courseName, order, reqCourseName, passed, score, avalilable, nil
}

// PassCourse approves the course of the name indicated with the score supplied.
func (r *repo) PassCourse(userId int, courseName string, score float32) error {
	u, ok := r.IdUserLookUp[userId]
	if !ok {
		return fmt.Errorf("user with id: %v not found", userId)
	}
	return u.passCourse(courseName, score)
}

// DeleteUser deletes the user with the id given, and returns his name.
func (r *repo) DeleteUser(id int) (string, error) {
	ok := false
	index := 0
	for i, u := range r.Users {
		if u.Id == id {
			ok = true
			index = i
			break
		}
	}
	if !ok {
		return "", fmt.Errorf("user with id: %v not found", id)
	}
	delete(r.IdUserLookUp,r.Users[index].Id)
	name := r.Users[index].Name
	r.Users = append(r.Users[:index], r.Users[index+1:]...)
	return name, nil
}

// UserInfo returns the information from all the users.
func (r *repo) UsersInfo() ([]string, []int, error){
	var names []string
	var ids []int
	for _, u := range r.Users {
		names = append(names, u.Name)
		ids = append(ids, u.Id)
	}
	return names, ids, nil

}