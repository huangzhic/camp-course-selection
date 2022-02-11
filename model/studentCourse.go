package model


type StudentCourse struct {
	STUDENT_ID   int64  `gorm:"primarykey;type:bigint(20);not null;comment:学生ID"`
	COURSE_ID   int64  `gorm:"primarykey;type:bigint(20);not null;comment:课程ID"`
}

func (StudentCourse) TableName() string {
	return "t_student_course"
}
