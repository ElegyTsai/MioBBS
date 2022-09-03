package system

import (
	"Miyo_Assignment/global"
	"Miyo_Assignment/model/user"
	"Miyo_Assignment/utils"
	"errors"
	"gorm.io/gorm"
)

type UserService struct{}

func (UserService *UserService) Register(user system.User) (userBack system.User, err error) {
	if err := UserService.isOkayRegister(user); err != nil {
		return user, err
	}
	// 是否可以注册

	user.Password = utils.BcryptHash(user.Password) //单向加密
	err = global.GV_DB.Create(&user).Error
	user.Password = "" // 密码不用返回
	return user, err
}

func (UserService *UserService) Login(user system.User) (userBack system.User, err error) {
	searchResult := system.User{}
	e := global.GV_DB.First(searchResult, "uid = ?", user.UID).Error
	if errors.Is(e, gorm.ErrRecordNotFound) {
		global.GV_LOG.Infof("UID:%d not found", user.UID)
		return user, errors.New("This account has not been registered yet ")
	}
	if !utils.BcryptCheckIsRight(user.Password, searchResult.Password) {
		global.GV_LOG.Infof("UID:%d password is wrong", user.UID)
		return user, errors.New("wrong password")
	}
	searchResult.Password = "" // 密码不用返回
	return searchResult, nil
}

func (UserService *UserService) isOkayRegister(user system.User) (err error) {
	if !errors.Is(global.GV_DB.Where("username = ?", user.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册")
	}
	if !errors.Is(global.GV_DB.Where("phone = ?", user.Phone).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("手机已注册")
	}
	if !errors.Is(global.GV_DB.Where("email = ?", user.Email).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("邮箱已注册")
	}
	return nil
}
