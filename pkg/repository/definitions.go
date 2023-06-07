package repository

type repo struct {
	Users []User
	LastId int
}

type User struct {
	Name string
	Id int
	Courses []Course
	State CoursesState
}

type Course struct {
	Name string
	Order int
	Required *Course
	Passed bool
	Score float32
	Available bool
}

type CoursesState struct {
	PassedCourses []*Course
	AvailableCourse []*Course
	NotAvailableCourse []*Course
}