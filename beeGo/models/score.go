package models

import "github.com/astaxie/beego/orm"

type Score struct {
	Id        int      `orm:"pk;column(id)"`
	StuId int      `orm:"column:(stu_id)"`
	CName     string   `orm:"column:(c_name)"`
	Grade     string   `orm:"column:(grade)"`
}

func init()  {
	orm.RegisterModel(new(Score))
}

//更加stuId 获取学生成绩
func GetOneStudentScoreById(studentId int) (s []Score) {
	o := orm.NewOrm()
	o.Raw("select * from score where stu_id = ?",studentId).QueryRows(&s)
	return
}