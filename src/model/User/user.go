package system

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UID       int64     `gorm:"primaryKey;AUTO_INCREMENT=1000000"` //独一无二的id 作为主键
	Username  string    `gorm:"type:varchar(30)"`                  //登录用的名字唯一不可改
	Nickname  string    `gorm:"type:varchar(30)"`                  // 昵称，可以修改
	Email     string    `gorm:"unique;index"`                      //注册用的邮箱
	Phone     string    `gorm:"unique;index"`                      //注册用的手机号码
	Password  string    `gorm:"type:varchar(30)"`                  //密码
	CreatedAt time.Time //和gorm有关
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
