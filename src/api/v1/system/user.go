package systemApi

import (
	v1 "Miyo_Assignment/api/v1"
	"Miyo_Assignment/global"
	"Miyo_Assignment/model/request"
	response "Miyo_Assignment/model/response"
	model "Miyo_Assignment/model/user"
	"Miyo_Assignment/utils"
	"github.com/gin-gonic/gin"
)

type SystemApi struct{}

func (SystemApi *SystemApi) Register(c *gin.Context) {
	var r request.Register
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
	response.OkWithAll(c, response.RegisterResponse{User: userBack}, "注册成功")
}

func (SystemApi *SystemApi) Login(c *gin.Context) {
	var r request.Login
	_ = c.BindJSON(&r)
	user := model.User{UID: r.UID, Password: r.Password}
	userBack, err := v1.UserServiceApp.Login(user)
	if err != nil {
		global.GV_LOG.Errorf("登陆失败,原因：%s", err.Error())
		response.FailResponseWithMsg(c, err.Error())
		return
	}
	//如果登陆成功的话，需要添加token信息
	SystemApi.ObtainToken(c, userBack)

}

func (SystemApi *SystemApi) ObtainToken(c *gin.Context, user model.User) {
	j := utils.NewJWT()
	claims := j.CreateClaims(request.BaseClaims{
		UID:      user.UID,
		Nickname: user.Nickname,
		Username: user.Username,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		global.GV_LOG.Errorf("token generation error: %s", err.Error())
		response.FailResponseWithMsg(c, "Token generation error")
	}
	response.OkWithAll(c, response.LoginResponse{
		user,
		token,
		claims.ExpiresAt.Unix(),
	}, "login success")
	global.GV_LOG.Infof("user:%s log in successfully with token:%s", user.UID, token)
}
