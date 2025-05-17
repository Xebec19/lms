package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	CourseID   string   `gorm:"column:course_id" json:"course_id"`
	CourseName string   `gorm:"column:course_name" json:"course_name"`
	CourseSlug string   `gorm:"column:course_slug" json:"course_slug"`
	CourseDesc string   `gorm:"column:course_desc" json:"course_desc"`
	CourseImg  string   `gorm:"column:course_img" json:"course_img"`
	Tags       string   `gorm:"column:tags" json:"tags"`
	Cohorts    []Cohort `gorm:"foreignKey:CourseID" json:"cohorts"`
}
