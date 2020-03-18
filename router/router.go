package router

import (
	"github.com/kataras/iris"
	"irisProject/controller"
)

func CollectRouter(app *iris.Application) *iris.Application {
	//注册请求
	app.Post("/api/auth/register", controller.Register)

	return app
}
