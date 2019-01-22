package controllers

import (
	"beeGo/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"strconv"
)

type StudentController struct {
	beego.Controller
}

// @Title 获取学生信息
// @Description 全都学生，单个学生，根据条件查询
// @Param	studentId		path 	string	false		"根据studentId查询"
// @Param	studentName		path 	string	false		"根据studentName查询"
// @Param	studentSex		path 	string	false		"根据studentSex查询"
// @Param	department		path 	string	false		"根据department查询"
// @Success 200 {code,msg} models.Student
// @Failure 403 :studentId is empty
// @router /GetStudentInfo [get]
func (s *StudentController) GetStudentInfo() {
	studentId := s.GetString("studentId")
	if studentId != "" {
		studentId, _ := strconv.Atoi(s.GetString("studentId"))
		student, err := models.GetOneStudent(studentId)
		if err == nil {
			s.Data["json"] = Response{0,"操作成功",student}
		} else {
			s.Data["json"] = ErrResponse{-1,err}
		}
	} else {
		students, err := models.GetAllStudent()
		if err == nil {
			s.Data["json"] = Response{0,"操作成功",students}
		} else {
			s.Data["json"] = ErrResponse{-1,err}
		}
	}
	s.ServeJSON()
}

// @Title 获取学生成绩信息
// @Description 根据条件查询
// @Param	studentId		path 	string	false		"根据studentId查询"
// @Success 200 {code,msg} interface{}
// @Failure 403 :studentId is empty
// @router /GetStudentScore [get]
func (s *StudentController) GetStudentScore(){
	studentId := s.GetString("studentId")
	if studentId != "" {
		studentId, _ := strconv.Atoi(studentId)
		stu, err := models.GetOneStudent(studentId)
		stus :=models.GetOneStudentScoreById(studentId)
		maps := map[string]interface{}{
			"Id":stu.Id,
			"Name":stu.Name,
			"Sex":stu.Sex,
			"Birth":stu.Birth,
			"Department":stu.Department,
			"Address":stu.Address,
			"scoreData":stus,
		}
		if err != nil{
			s.Data["json"] =ErrResponse{-1,"系统错误"}
		}else {
			s.Data["json"] = Response{0,"操作成功",maps}
		}

	}else {
		students ,err :=models.GetStudentScoreInfo()
		if err != nil{
			s.Data["json"] =ErrResponse{-1,"系统错误"}
		}else {
			s.Data["json"] = Response{0,"操作成功",students}
		}
	}
	s.ServeJSON()
}

// @Title 更新学生信息
// @Description 根据条件查询
// @Param	studentInfo		body 	models.Student	true		"更新学生数据"
// @Success 200 {code,msg} interface{}
// @Failure 403 :studentId is empty
// @router /UpdateStudent [post]
func (s *StudentController)UpdateStudent() {
	var students models.Student
	json.Unmarshal(s.Ctx.Input.RequestBody,&students)
	bl := models.UpdateStuInfo(students)
	if bl{
		s.Data["json"] = Response{0,"操作成功",bl}
	}else {
		s.Data["json"] = ErrResponse{-1,"系统错误"}
	}
	s.ServeJSON()
}

// @Title 新增学生信息
// @Description 批量新增学生信息
// @Param	studentInfo		body 	[]models.Student	true		"更新学生数据"
// @Success 200 {code,msg} interface{}
// @Failure 403 list
// @router /AddStudents [post]
func (s *StudentController)AddStudents()  {
	var students []models.Student
	json.Unmarshal(s.Ctx.Input.RequestBody,&students)
	err , snum := models.AddStudentInfos(students)
	if err {
		s.Data["json"]=Response{0,"操作成功", map[string]interface{}{"success":snum,"bl":err}}
	}else {
		s.Data["json"] = ErrResponse{-1,"操作失败"}
	}
	s.ServeJSON()
}

// @Title 删除学生信息
// @Description 根据条件查询
// @Param	studentId		path 	string	true		"根据studentId查询"
// @Success 200 {code,msg}
// @Failure 403 :studentId is empty
// @router /DeleteStudent [get]
func (s *StudentController) DeleteStudent(){
	studentId := s.GetString("studentId")
	if studentId != ""{
		bl := models.DeleteStu(studentId)
		if bl {
			s.Data["json"] = Response{0,"操作成功",bl}
		}else {
			s.Data["json"] = ErrResponse{-1,"系统错误"}
		}
	}
	s.ServeJSON()
}