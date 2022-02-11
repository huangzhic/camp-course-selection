package model


type StudentCourse struct {
	UserID   int64  `gorm:"primarykey;type:bigint(20);not null;comment:学生ID"`
	CourseID   int64  `gorm:"primarykey;type:bigint(20);not null;comment:学生ID"`
}

func (StudentCourse) TableName() string {
	return "t_student_course"
}

// GetUser 用ID获取用户

//func GetUser(ID interface{}) (TMember, error) {
//	var user TMember
//	result := DB.First(&user, ID)
//	return user, result.Error
//}