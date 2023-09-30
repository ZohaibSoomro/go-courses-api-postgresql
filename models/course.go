package models

type Course struct {
	ID               uint   `json:"id" gorm:"ID primarykey"`
	Title            string `json:"title"`
	EnrolledStudents string `json:"enrolled_no_of_students" gorm:"enrolled_students"`
}
