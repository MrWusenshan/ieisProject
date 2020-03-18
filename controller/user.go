package controller

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
	"irisProject/common"
	"irisProject/model"
	"irisProject/util"
	"log"
	"net/http"
)

//用户注册函数
func Register(ctx iris.Context) {
	//获取数据库引擎
	db := common.GetDbEngine()

	//从ctx中读取数据
	name := ctx.PostValue("name")
	telephone := ctx.PostValue("telephone")
	password := ctx.PostValue("password")

	//验证手机号
	// 1. 验证手机号码是否合法  (先简单验证电话号码长度是否是11位) todo
	if len(telephone) != 11 {
		_, _ = ctx.JSON(iris.Map{
			"code":    http.StatusUnprocessableEntity,
			"message": "无效的手机号码",
		})

		return
	}

	// 2. 验证手机号码是否已注册
	if isTelephoneExist(db, telephone) {
		_, _ = ctx.JSON(iris.Map{
			"code":    iris.StatusUnprocessableEntity,
			"message": "该手机号已被注册",
		})

		return
	}

	//验证密码是否合法	(长度不小于6)
	if len(password) < 6 {
		_, _ = ctx.JSON(iris.Map{
			"code":    http.StatusUnprocessableEntity,
			"message": "密码长度不能少于6位",
		})

		return
	}

	//用户是否提交名称,若未提交,则自动生成随机名称
	if name == "" {
		name = util.RandomString(10)
	}

	//加密用户的密码
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		_, _ = ctx.JSON(iris.Map{
			"code":    iris.StatusInternalServerError,
			"message": "对用户密码加密出错",
		})

		return
	}

	//创建用户对象
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	//将用户信息添加至user表
	db.Create(&newUser)

	_, _ = ctx.JSON(iris.Map{
		"code":    http.StatusOK,
		"message": "注册成功",
	})
}

//查询手机号是否注册函数
func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

//用户登录函数
func Login(ctx iris.Context) {
	//获取数据库引擎
	db := common.GetDbEngine()

	//从ctx中读取数据
	telephone := ctx.PostValue("telephone")
	password := ctx.PostValue("password")

	var user model.User
	//根据手机号从数据库中查找数据并返回给 user对象
	db.Where("telephone = ?", telephone).First(&user)

	//判断是否有对应的ID
	if user.ID == 0 {
		_, _ = ctx.JSON(iris.Map{
			"code":    iris.StatusUnprocessableEntity,
			"message": "用户账户不存在",
		})

		return
	}

	//判断密码是否正确
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		_, _ = ctx.JSON(iris.Map{
			"code":    iris.StatusBadRequest,
			"message": "密码错误",
		})

		return
	}

	//发放 token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(iris.Map{
			"code":    iris.StatusInternalServerError,
			"message": "系统异常",
		})

		log.Printf("token generate error : %v", err)

		return
	}

	//返回登录信息
	ctx.JSON(iris.Map{
		"code":    iris.StatusOK,
		"token":   token,
		"message": "登录成功",
	})
}

func Info(ctx iris.Context) {

}
