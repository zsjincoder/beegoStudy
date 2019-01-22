// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beeGo/controllers"
	"beeGo/util/filter"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/api",
		//beego.NSBefore(filter.DumpHttpFilter),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/student",
			beego.NSCond(filter.ValidUserToken),
			beego.NSInclude(
				&controllers.StudentController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
