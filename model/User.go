package model

import (
	"boKe/utils/errmsg"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(20);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type:int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色码"`
}

//查询用户是否存在

func CheckUser(name string) (code int) {
	var users User
	db.Select("id").Where("username = ?", name).First(&users)
	if users.ID > 0 {
		return errmsg.ERROR_USERNAME //1001
	}
	return errmsg.SUCCESS //200
}

//新增用户

func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password)
	err := db.Debug().Create(&data).Error
	fmt.Println(data)
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS // 200
}

//查询单个用户

func GetUser(id int) (User, int) {
	var user User
	err := db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

//查询用户列表

func GetUsers(username string, pageSize int, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offset = -1
	}
	if username == "" {
		err = db.Limit(pageSize).Offset(offset).Find(&users).Count(&total).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return nil, 0
		}
	} else {
		err1 := db.Where("username like ?", username+"%").Limit(pageSize).Offset(offset).Find(&users).Count(&total).Error
		if err1 != nil && err1 != gorm.ErrRecordNotFound {
			return nil, 0
		}
	}

	return users, total
}

//	编辑用户

func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Debug().Model(&user).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//	删除用户

func DeleteUser(id int) int {
	var user User
	err = db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//	密码加密

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return
}

func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}

	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// 登陆验证

func CheckLogin(username string, password string) int {
	var user User

	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		return errmsg.ERROR_USERNAME
	}
	if ScryptPw(password) != user.Password {
		return errmsg.ERROR_USERPASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NO_RIGHT
	}

	return errmsg.SUCCESS

}
