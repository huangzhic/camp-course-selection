package exception

// 说明：
// 1. 所提到的「位数」均以字节长度为准
// 2. 所有的 ID 均为 int64（以 string 方式表现）

// 通用结构

const (
	OK                 int = 0
	ParamInvalid       int = 1  // 参数不合法
	UserHasExisted     int = 2  // 该 Username 已存在
	UserHasDeleted     int = 3  // 用户已删除
	UserNotExisted     int = 4  // 用户不存在
	WrongPassword      int = 5  // 密码错误
	LoginRequired      int = 6  // 用户未登录
	CourseNotAvailable int = 7  // 课程已满
	CourseHasBound     int = 8  // 课程已绑定过
	CourseNotBind      int = 9  // 课程未绑定过
	PermDenied         int = 10 // 没有操作权限
	StudentNotExisted  int = 11 // 学生不存在
	CourseNotExisted   int = 12 // 课程不存在
	StudentHasNoCourse int = 13 // 学生没有课程
	StudentHasCourse   int = 14 // 学生有课程

	CourseHasExisted          int = 15  // 课程已经存在
	StudentCannotBindOrUnBind int = 16  // 课程不能与学生进行绑定或解绑等的相关操作

	Coursed int = 20 // 已经选过课了
	
	UnknownError              int = 255 // 未知错误
)

var Code2String = map[int]string{
	0:  "操作成功",
	1:  "参数不合法",
	2:  "该 Username 已存在",
	3:  "用户已删除",
	4:  "用户不存在",
	5:  "密码错误",
	6:  "用户未登录",
	7:  "课程已满",
	8:  "课程已绑定过",
	9:  "课程未绑定过",
	10: "没有操作权限",
	11: "学生不存在",
	12: "课程不存在",
	13: "学生没有课程",
	14: "学生有课程",

	15:  "课程已经存在",
	16:  "课程不能与学生进行绑定或解绑等的相关操作",

	20: "已经选过课了",
	255: "未知错误",
}
