package repository

import "github.com/AtilioBoher/ordenadoDeCursos/pkg/handler"

type repo struct {
	Users    []User
	LastId   int
	IdUserLookUp map[int]*User
	sorter   Sorter
}

type User struct {
	Name    string
	Id      int
	Courses []Course
	IdCourseLookUp map[string]*Course
	State   CoursesState
}

type Course struct {
	Name      string
	Order     int
	Required  *Course
	Passed    bool
	Score     float32
	Available bool
}

type CoursesState struct {
	PassedCourses      []*Course
	AvailableCourses    []*Course
	NotAvailableCourses []*Course
}

type Sorter interface {
	StoreCourses(courses []handler.Course, u *User) ([]handler.OrCourse, error)
}
