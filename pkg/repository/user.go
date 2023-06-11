package repository

func (u *User) CoursesInfo()([]string, []int, []string, []bool, []float32, []bool) {
	var (
		courseName, reqCourseName []string
		order []int
		passed []bool
		score []float32
		avalilable []bool
	)
	for _, c := range u.Courses {
		courseName = append(courseName, c.Name)
		order = append(order, c.Order)
		if c.Required != nil {
			reqCourseName = append(reqCourseName, c.Required.Name)
		} else {
			reqCourseName = append(reqCourseName, "this course doesn't have requisites")
		}
		passed = append(passed, c.Passed)
		score = append(score, c.Score)
		avalilable = append(avalilable, c.Available)
	}
	return courseName, order, reqCourseName, passed, score, avalilable
}