package systemApi

import (
	v1 "Miyo_Assignment/api/v1"
	"Miyo_Assignment/global"
	"Miyo_Assignment/model/Request"
	response "Miyo_Assignment/model/Response"
	model "Miyo_Assignment/model/User"
	"github.com/gin-gonic/gin"
)

type SystemApi struct{}

func (SystemApi *SystemApi) Register(c *gin.Context) {
	var r Request.Register
	_ = c.BindJSON(&r)
	user := model.User{Username: r.Username, Nickname: r.NickName, Email: r.Email, Phone: r.Phone, Password: r.Password}
	userBack, err := v1.UserServiceApp.Register(user)
	if err != nil {
		global.GV_LOG.Errorf("注册失败，原因: %s", err.Error())
		response.FailResponseWithMsg(c, err.Error())
		return
	}
	//用户注册写入成功，开始处理权限
	v1.CasbinServiceApp.AddDefaultRole(&user)
	response.OkWithAll(c, response.ResigterResponse{User: userBack}, "注册成功")
}

func (SystemApi *SystemApi) Login(c *gin.Context) {
	var r Request.Login
	_ = c.BindJSON(&r)
	user := model.User{UID: r.UID, Password: r.Password}
	userBack, err := v1.UserServiceApp.Login(user)
	if err != nil {
		global.GV_LOG.Errorf("登陆失败,原因：%s", err.Error())
		response.FailResponseWithMsg(c, err.Error())
		return
	}
	//如果登陆成功的话，需要添加token信息

}
