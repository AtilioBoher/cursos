package handler

// myHandler is a struct to handle request from the course API.
type myHandler struct {
	service Service
}

// Course is a pair Desired and Requires courses.
type Course struct {
	Desired  string `json:"desiredCourse"`
	Required string `json:"requiredCourse"`
}

// oReq is the request data for ordering courses.
type oReq struct {
	UserId  int   `json:"userId"`
	Courses []Course `json:"courses"`
}

// oRes is the response data for ordering courses.
type oRes struct {
	UserId  int   `json:"userId"`
	Courses []OrCourse `json:"courses"`
}

// OrCourse is a struct for ordering courses.
type OrCourse struct {
	Name  string `json:"course"`
	Order int    `json:"order"`
}

// CourseInfo has the complete information of a course.
type CourseInfo struct {
	OrCourse
	Required string `json:"requiredCourse"`
	Passed bool `json:"passed"`
	Score float32 `json:"score"`
	Available bool `json:"available"`
}

// InfoResponse is the response structure of the CoursesInfo() function.
type InfoResponse struct {
	UserId  int   `json:"userId"`
	Courses []CourseInfo `json:"courses"`
}

// Service is an interface that has all the methods needed by the myHandler handler.
type Service interface {
	SortCourses([]Course) ([]OrCourse, error)
	StoreCourses([]Course, int) ([]OrCourse, error)
	StoreNewUser(name string) (int, error)
	GetUser(id int) (string, error)
	CoursesInfo(id int) ([]string, []int, []string, []bool, []float32, []bool, error)
	PassCourse(id int, courseName string, score float32) error
	DeleteUser(id int) (string, error)
	UsersInfo() ([]string, []int, error)
}

type User struct {
	Name string `json:"name"`
	Id	int `json:"userId"`
}