package service

import (
	"camp-course-selection/cache"
	"camp-course-selection/common/util"
	"camp-course-selection/model"
	"camp-course-selection/vo"
	amqp "github.com/rabbitmq/amqp091-go"
	rabbitmq "github.com/wagslane/go-rabbitmq"
	"os"
	"time"
)

type StudentService struct {
}

func (m *StudentService) BookCourse(v *vo.BookCourseRequest) (res vo.BookCourseResponse) {
	//判断这个学生是否已抢过课 redis中的key格式 StudentID_CourseID
	redisKey := v.StudentID + "_" + v.CourseID
	b, err := cache.RedisClient.SetNX(redisKey, 1, 5*time.Minute).Result()
	if err != nil {
		util.Log().Error("BookCourse SetNX Error : %v", err)
		res.Code = vo.UnknownError
		return
	}
	if b == false {
		//重复请求
		res.Code = vo.RepeatRequest
		return
	}
	//库存的key为课程ID
	var num int64
	num, err = cache.RedisClient.Decr(v.CourseID).Result()
	if num < 0 {
		res.Code = vo.CourseNotAvailable
		return
	}
	//抢课成功，快速返回
	publisher, err := rabbitmq.NewPublisher(os.Getenv("RABBITMQ_DSN"), amqp.Config{},
		rabbitmq.WithPublisherOptionsLogging)
	if err != nil {
		util.Log().Error("NewPublisher Error : %v", err)
		res.Code = vo.UnknownError
		return
	}
	err = publisher.Publish(
		[]byte("helloworld"),
		[]string{"book.course"},
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsExchange("student-book-course"),
	)
	if err != nil {
		util.Log().Error("Publish Error : %v", err)
		res.Code = vo.UnknownError
		return
	}
	res.Code = vo.OK
	return
}

func (m *StudentService) GetStudentCourse(v *vo.GetStudentCourseRequest) (res vo.GetStudentCourseResponse) {
	courses := []model.StudentCourse{}
	if err := model.DB.Where("STUDENT_ID = ?", v.StudentID).Find(&courses).Error; err != nil {
		res.Code = vo.StudentNotExisted
		return
	}
	var courseList = make([]vo.TCourse, len(courses))
	for i := 0; i < len(courses); i++ {
		if err := model.DB.Where("course_id = ?", courses[i].CourseID).Find(&courseList[i]).Error; err != nil {
			res.Code = vo.CourseNotExisted
			return
		}
	}
	res.Code = vo.OK
	res.Data.CourseList = courseList
	return
}
