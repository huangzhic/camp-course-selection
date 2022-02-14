package service

import (
	"camp-course-selection/common/constants"
	"camp-course-selection/common/util"
	"camp-course-selection/model"
	"camp-course-selection/vo"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"strconv"
)

type CourService struct{}

//-----------创建课程-------------------------------------------------------
func (m *CourService) CreateCourse(courseVo *vo.CreateCourseRequest) (res vo.CreateCourseResponse) {
	course := model.TCourse{
		Name:        courseVo.Name,
		CourseStock: courseVo.Cap,
	}

	// 表单验证
	if err := model.DB.Model(course).Where("name = ?", courseVo.Name).Find(&course); err != nil {
		res.Code = vo.CourseHasExisted
		res.Data.CourseID = string(course.CourseID)
		return
	}

	// 雪花ID
	node, err := snowflake.NewNode(1)
	if err != nil {
		util.Log().Error("generate SnowFlakeID Error: %v", err)
		res.Code = vo.UnknownError
		return
	}

	id := node.Generate()
	course.CourseID = int64(id)
	course.TeacherID = "-1" // -1代表该课程未被绑定

	// 创建课程
	if err := model.DB.Create(&course).Error; err != nil {
		res.Code = vo.UnknownError
	} else {
		res.Code = vo.OK
		res.Data.CourseID = string(course.CourseID)
	}
	return
}

//-----------获取课程-------------------------------------------------------
func (m *CourService) GetCourse(v *vo.GetCourseRequest) util.R {
	course := &model.TCourse{}
	err := model.DB.Where("course_id = ?", v.CourseID).First(&course).Error
	fmt.Println(course, "111")
	if err == nil {
		return *util.Ok(course)
	}
	return *util.Error(exception.CourseNotExisted)
}

//------------绑定课程-------------------------------------------------------
func (m *CourService) BindCourseService(v *vo.BindCourseRequest) util.R {
	if err := checkBindCourseAndTeacher(v); err == exception.OK {
		return *util.Ok("绑定")
	} else {
		return *util.Error(err)
	}
}

func checkBindCourseAndTeacher(v *vo.BindCourseRequest) int {
	teacher := &model.TMember{}
	course := &model.TCourse{}
	// 检测老师是否正确
	if err := model.DB.Where("user_id = ?", v.TeacherID).First(&teacher).Error; err == nil {
		if teacher.UserType != constants.Teacher {
			return exception.StudentCannotBindOrUnBind
		}
	} else {
		return exception.UserNotExisted
	}
	// 检查课程是否正确
	if err := model.DB.Where("course_id = ?", v.CourseID).First(&course).Error; err == nil {
		if course.TeacherID != "-1" {
			return exception.CourseHasBound
		}
	} else {
		return exception.CourseNotExisted
	}
	// 绑定课程
	course.TeacherID = strconv.FormatInt(int64(teacher.UserID), 10)
	if err := model.DB.Model(&course).Update("teacher_id", course.TeacherID).Error; err == nil {
		return exception.OK
	}
	return exception.UnknownError
}

//-----------------解绑课程---------------------------------------------------
func (m *CourService) UnBindCourseService(v *vo.UnbindCourseRequest) util.R {
	if err := checkUnBindCourseAndTeacher(v); err == exception.OK {
		return *util.Ok("解绑")
	} else {
		return *util.Error(err)
	}
}

func checkUnBindCourseAndTeacher(v *vo.UnbindCourseRequest) int {
	teacher := &model.TMember{}
	course := &model.TCourse{}
	// 检测老师是否正确
	if err := model.DB.Where("user_id = ?", v.TeacherID).First(&teacher).Error; err == nil {
		if teacher.UserType != constants.Teacher {
			return exception.StudentCannotBindOrUnBind
		}
	} else {
		return exception.UserNotExisted
	}
	// 检查课程是否正确
	if err := model.DB.Where("course_id = ?", v.CourseID).First(&course).Error; err == nil {
		if course.TeacherID == "-1" {
			// 课程已经解绑
			return exception.OK
		}
	} else {
		return exception.CourseNotExisted
	}
	// 解绑课程
	course.TeacherID = "-1"
	if err := model.DB.Model(&course).Update("teacher_id", course.TeacherID).Error; err == nil {
		return exception.OK
	}
	return exception.UnknownError
}

//------------------获取老师所有课程-----------------------------------------
func (m *CourService) GetTeacherCourseService(v *vo.GetTeacherCourseRequest) util.R {
	teacher := &model.TMember{}
	// 检测老师是否正确
	if err := model.DB.Where("user_id = ?", v.TeacherID).First(&teacher).Error; err == nil {
		if teacher.UserType != constants.Teacher {
			return *util.Error(exception.StudentCannotBindOrUnBind)
		}
	} else {
		fmt.Println(teacher)
		return *util.Error(exception.UserNotExisted)
	}
	// 获取结果
	courses := make([]model.TCourse, 0)
	result := make([]string, 0)
	if err := model.DB.Where("teacher_id = ?", v.TeacherID).Find(&courses).Error; err == nil {
		if len(courses) == 0 {
			return *util.Ok("该老师无教学课程")
		}
		for _, v := range courses {
			result = append(result, v.Name)
		}
		return *util.Ok(result)
	}
	return *util.Error(exception.UnknownError)
}

//-------------------排课求解器--------------------------------------------
func (m *CourService) ScheduleCourse(schedule vo.ScheduleCourseRequest) util.R {
	// result  key:课程  val：对应的老师
	result := make(map[string]string, 0)
	teachers := make([]string, 0)
	courses := make([]string, 0)
	// set   key:课程   val： 是否被被遍历 find函数需要
	set := make(map[string]bool, 0)
	relation := &schedule.TeacherCourseRelationShip
	// 提取teacher与course集合
	for k, v := range *relation {
		teachers = append(teachers, k)
		for _, val := range v {
			if _, ok := set[val]; ok == false {
				set[val] = false
				courses = append(courses, val)
			}
		}
	}
	// 二分图匹配
	for _, v := range teachers {
		clear(&set)
		find(v, relation, &set, &result)
	}
	// result中key是课程， ans中key是老师
	ans := make(map[string]string, 0)
	for k, v := range result {
		ans[v] = k
	}
	return *util.Ok(ans)
}

func find(v string, relation *map[string][]string, set *map[string]bool, result *map[string]string) bool {
	courses, _ := (*relation)[v]
	for _, course := range courses {
		if ok, _ := (*set)[course]; ok == false {
			(*set)[course] = true
			if val, ok := (*result)[course]; ok == false || find(val, relation, set, result) == true {
				(*result)[course] = v
				return true
			}
		}
	}
	return false
}

func clear(m *map[string]bool) {
	for k, _ := range *m {
		(*m)[k] = false
	}
}
