package middleware

import (
	"github.com/kataras/iris"
	"irisProject/common"
	"irisProject/model"
	"strings"
)

func AuthMiddleware() iris.Handler {
	return func(ctx iris.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(iris.Map{
				"code":    401,
				"message": "权限不足",
			})

			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParesToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(iris.Map{
				"code":    401,
				"message": "权限不足",
			})

			return
		}

		//验证通过token
		userId := claims.UserID
		DB := common.GetDbEngine()
		var user model.User
		DB.First(&user, userId)

		//用户
		if user.ID == 0 {
			ctx.JSON(iris.Map{
				"code":    401,
				"message": "权限不足",
			})

			return
		}

		//用户存在 将信息写入 context
		//ctx.Set("user", user)

		ctx.Next()
	}
}
