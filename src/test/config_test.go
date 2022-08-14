package main

import (
	"Miyo_Assignment/global"
	"Miyo_Assignment/initialize"
	model "Miyo_Assignment/model/User"
	"testing"
)

/*用于测试GORM的连接是否正常
 */
func TestGorm(t *testing.T) {
	initialize.Init()

	//t.Errorf("error")
}

func TestGormCreateUserTable(t *testing.T) {
	initialize.Init()
	global.GV_DB.Table("users").AutoMigrate(&model.User{})
}
func TestGormCreateUser(t *testing.T) {
	initialize.Init()
	u := model.User{Username: "admin", Nickname: "wei", Phone: 123, Email: "123@gmail.com", Password: "1234"}
	global.GV_DB.Create(&u)

}
func TestCasbin(t *testing.T) {

}
