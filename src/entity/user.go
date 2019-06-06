package entity

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User 用户表
type User struct {
	gorm.Model
	Name     string `gorm:"UNIQUE"`
	Password string
	Email    string `gorm:"UNIQUE"`
	Role     string
}

// AddUser 添加用户
func (u User) AddUser(user User) bool {
	db, err := gorm.Open("mysql", "touch:1314@/touch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("连接失败")
	}
	if db.HasTable(&User{}) {
		db.Create(&user)
		return true
	} else {
		log.Println("用户表未创建，现在执行创建表操作")
		db.CreateTable(&User{})
		log.Println("用户表创建成功")
		db.Create(&user)
		if db.NewRecord(user) {

		} else {
			return true
		}
	}
	return false
}

// DeleteUser 删除指定用户
func (u User) DeleteUser(user User) bool {
	db, err := gorm.Open("mysql", "touch:1314@/touch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("连接失败")
	}
	db.Delete(&user)
	if db.Find(&user) == nil {
		log.Println("删除成功", user)
		return true
	}
	return false

}

// FindUser 查询指定用户
func (u User) FindUser(user User) *gorm.DB {
	db, err := gorm.Open("mysql", "touch:1314@/touch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("连接失败")
	}
	return db.First(&user)
}

// UpdateUser 更新用户
func (u User) UpdateUser(user User) bool {
	db, err := gorm.Open("mysql", "touch:1314@/touch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("连接失败")
	}
	db.Model(&user).Updates(user)
	return true
}


//FindUserName 查看用户名是否存在，用于注册
func(u User) FindUserName(name string) bool{
	db, err := gorm.Open("mysql", "touch:1314@/touch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("连接失败")
	}
	var user []User
	db.Where("name=?",name).First(&user)
	return len(user)>0
}

func(u User) FindUserEmail(email string) bool{
	db, err := gorm.Open("mysql", "touch:1314@/touch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("连接失败")
	}
	var user []User
	db.Where("email=?",email).First(&user)
	return len(user)>0
}