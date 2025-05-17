package models

import "gorm.io/gorm"

type Cohort struct {
	gorm.Model
	CohortID   string `gorm:"column:cohort_id" json:"cohort_id"`
	CohortName string `gorm:"column:cohort_name" json:"cohort_name"`
	CohortSlug string `gorm:"column:cohort_slug" json:"cohort_slug"`
	Course     Course `gorm:"foreignKey:CourseID" json:"course"`
}
