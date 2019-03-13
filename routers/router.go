// @APIVersion 1.0.0
// @Title 夸夸机器人
// @Description 夸夸机器人后台API文档
// @Contact gng@bingyan.net
// @TermsOfServiceUrl http://bingyan.net/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"kuakuaAi/controllers"
)

func init() {
	// 测试graphql
	//beego.Any("/graphql", controllers.GraphqlFilter)
	//authNS := beego.NewNamespace("/auth",
	//	beego.NSNamespace("/access_token",
	//		beego.NSInclude(&controllers.AccessTokenController{}),
	//	),
	//	beego.NSNamespace("/refresh_token",
	//		beego.NSInclude(&controllers.RefreshTokenController{}),
	//	),
	//)
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/test",
			beego.NSInclude(&controllers.TestController{}),
		),
		//beego.NSCond(func(ctx *context.Context) bool {
		//	if ctx.Request.Method == "GET" {
		//		return true
		//	}
		//	access_token := ctx.Request.Header.Get("Authorization")
		//	claims, err := utils.DecodeJWT(access_token)
		//	if err != nil || claims["IssueType"].(string) != "AccessToken" {
		//		return false
		//	}
		//	key := claims["UserId"].(string) + "AccessToken"
		//	c := models.GetRedis()
		//	defer c.Close()
		//	token, err := redis.String(c.Do("get", key))
		//	if err != nil || token != access_token {
		//		return false
		//	}
		//
		//	return true
		//}),
		//beego.NSAny("/graphql", controllers.GraphqlFilter),

	)
	//beego.AddNamespace(authNS)
	beego.AddNamespace(ns)

}
