package model

type College struct {
	CollegeId   int64
	CollegeName string
}

type Major struct {
	MajorId   int64
	MajorName string
	CollegeId int64
}
