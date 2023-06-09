package handler

type myHandler struct {
	service Service
}

type Course struct {
	Desired  string `json:"desiredCourse"`
	Requierd string `json:"requiredCourse"`
}

// oReq is the request for ordering courses
type oReq struct {
	UserId  int   `json:"userId"`
	Courses []Course `json:"courses"`
}

type oRes struct {
	UserId  int   `json:"userId"`
	Courses []OrCourse `json:"courses"`
}

// Ordered Course
type OrCourse struct {
	Name  string `json:"course"`
	Order int    `json:"order"`
}

type Service interface {
	SortCourses([]Course) ([]OrCourse, error)
	StoreCourses([]Course, int) ([]OrCourse, error)
	StoreNewUser(name string) (int, error)
	GetUser(id int) (string, error)

}

type User struct {
	Name string `json:"name"`
	Id	int `json:"userId"`
}