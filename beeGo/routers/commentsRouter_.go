package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["beeGo/controllers:StudentController"] = append(beego.GlobalControllerRouter["beeGo/controllers:StudentController"],
        beego.ControllerComments{
            Method: "AddStudents",
            Router: `/AddStudents`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeGo/controllers:StudentController"] = append(beego.GlobalControllerRouter["beeGo/controllers:StudentController"],
        beego.ControllerComments{
            Method: "DeleteStudent",
            Router: `/DeleteStudent`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeGo/controllers:StudentController"] = append(beego.GlobalControllerRouter["beeGo/controllers:StudentController"],
        beego.ControllerComments{
            Method: "GetStudentInfo",
            Router: `/GetStudentInfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeGo/controllers:StudentController"] = append(beego.GlobalControllerRouter["beeGo/controllers:StudentController"],
        beego.ControllerComments{
            Method: "GetStudentScore",
            Router: `/GetStudentScore`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeGo/controllers:StudentController"] = append(beego.GlobalControllerRouter["beeGo/controllers:StudentController"],
        beego.ControllerComments{
            Method: "UpdateStudent",
            Router: `/UpdateStudent`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeGo/controllers:UserController"] = append(beego.GlobalControllerRouter["beeGo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/Login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beeGo/controllers:UserController"] = append(beego.GlobalControllerRouter["beeGo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/Register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
