package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type Student struct {
	Id         int    `orm:"pk;column(id)"`
	Name       string `orm:"column:(name)"`
	Sex        string `orm:"column:(sex)"`
	Birth      string `orm:"column:(birth)"`
	Department string `orm:"column:(department)"`
	Address    string `orm:"column:(address)"`
}
//学生成绩信息
type StuInfo struct {
	StuId int    `json:"studentId"`
	Name      string `json:"name"`
	ScoreInfo struct {
		Department string `json:"department"`
		Grade      string `json:"scoreGrade"`
		Id         string `json:"scoreId"`
		CName      string `json:"scoreName"`
	}
}

func init() {
	orm.RegisterModel(new(Student))
}

//获取全部学生信息
func GetAllStudent() ([]Student, error) {
	var students []Student
	o := orm.NewOrm()
	_, err := o.Raw("select * from student").QueryRows(&students)
	return students, err
}

//获取单个学生信息
func GetOneStudent(studentId int) (Student, error) {
	var student Student
	student.Id = studentId
	o := orm.NewOrm()
	err := o.Read(&student)
	return student, err
}

//获取学生成绩
func GetStudentScoreInfoById(studentId int) ([]*StuInfo, error) {
	var student []*StuInfo
	o := orm.NewOrm()
	//_, err := o.QueryTable("Student").Filter("Id", studentId).RelatedSel().All(&student)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(
		"student.id", "student.name", "student.sex", "student.birth", "student.department", "student.address",
		"score.id", "score.stu_id", "score.c_name", "score.grade",
	).
		From("student").
		LeftJoin("score").On("student.id = score.stu_id").
		Where("student.id = ?")

	sql := qb.String()
	o.Raw(sql, studentId).QueryRows(&student)
	return student, nil
}

func GetStudentScoreInfo()([]*StuInfo, error) {
	var students []*StuInfo
	o := orm.NewOrm()
	//_, err := o.QueryTable("Student").RelatedSel().All(&students)
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select(
		"student.id", "student.name", "student.sex", "student.birth", "student.department", "student.address",
		"score.id", "score.stu_id", "score.c_name", "score.grade",
	).
		From("student").
		LeftJoin("score").On("student.id = score.stu_id")
	sql := qb.String()
	o.Raw(sql).QueryRows(&students)
	return students, nil
}

//新增学生信息
func AddStudentInfos(studentInfo []Student) (bool,int64){
	o := orm.NewOrm()
	successNum , err := o.InsertMulti(len(studentInfo),&studentInfo)
	if err == nil{
		if successNum > 0{
			return  true,successNum
		}else {
			return false,0
		}
	}
	return false,0
}

//更新学生信息
func UpdateStuInfo(studentInfo Student ) bool {
	o := orm.NewOrm()
	num, err := o.Update(&studentInfo)
	if err == nil {
		if num > 0{
			return true
		} else {
			return false
		}
	}
	return false
}

//删除学生信息
func DeleteStu(studentId string) bool {
	id,_ := strconv.Atoi(studentId)
	if num,err := orm.NewOrm().Delete(&Student{Id:id});err == nil{
		if num >0{
			return true
		}else{
			return false
		}
	}
	return false
}