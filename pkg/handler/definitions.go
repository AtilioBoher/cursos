package handler

type myHandler struct {
	service Service
}

type Course struct {
	Desired  string `json:"desiredCourse"`
	Required string `json:"requiredCourse"`
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

type CourseInfo struct {
	OrCourse
	Required string `json:"requiredCourse"`
	Passed bool `json:"passed"`
	Score float32 `json:"score"`
	Available bool `json:"available"`
}

type InfoResponse struct {
	UserId  int   `json:"userId"`
	Courses []CourseInfo `json:"courses"`
}

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