package main

import (
	"github.com/kataras/iris"
	"irisProject/common"
	"irisProject/router"
)

func main() {
	db := common.InitDbEngine()
	defer db.Close()

	app := iris.Default()

	app = router.CollectRouter(app)
	//app.Get("/api/info", func(ctx iris.Context) {
	//	_, _ = ctx.JSON(iris.Map{
	//		"code":    http.StatusOK,
	//		"message": "go iris server is success run",
	//	})
	//})
	_ = app.Run(iris.Addr(":8080"))
}
