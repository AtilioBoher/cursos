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
	UserId  string   `json:"userId"`
	Courses []Course `json:"courses"`
}

type oRes struct {
	UserId  string   `json:"userId"`
	Courses []OrCourse `json:"courses"`
}

// Ordered Course
type OrCourse struct {
	Name  string `json:"course"`
	Order int    `json:"order"`
}

type Service interface {
	SortCourses([]Course) ([]OrCourse, error)
}
