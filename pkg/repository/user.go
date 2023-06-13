package repository

import "fmt"

func (u *User) CoursesInfo() ([]string, []int, []string, []bool, []float32, []bool) {
	var (
		courseName, reqCourseName []string
		order                     []int
		passed                    []bool
		score                     []float32
		avalilable                []bool
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

func (u *User) passCourse(courseName string, score float32) error {
	ok := false
	var c *Course
	for i := range u.Courses {
		if u.Courses[i].Name == courseName {
			ok =true
			c = &u.Courses[i]
		}
	}
	if !ok {
		return fmt.Errorf("user %v doesn't have a course named %v", u.Name, courseName)
	}
	if !c.Available {
		if c.Passed {
			return fmt.Errorf("the course %s is already approved", c.Name)
		}
		if c.Required != nil {
			return fmt.Errorf("the course %v is not available, you need to pass %v first", c.Name,
			c.Required.Name)
		}
	}
	if score < 6 {
		return fmt.Errorf("to pass the course the score must be at least 6")
	}
	c.Passed = true
	c.Score = score
	c.Available = false

	for i := range u.Courses {
		if u.Courses[i].Required != nil {
			if u.Courses[i].Required.Name == c.Name {
				u.Courses[i].Available = true
			}
		}
	}
	return nil
}
