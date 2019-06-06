package entity

import (
	"github.com/jinzhu/gorm"
	"log"
)

//Opitions 选项表
type Opition struct {
	gorm.Model
	QuestionID uint
	Option   string
}

//AddOptions 增加选项
func (op Opition) AddOption(opition Opition) bool {
	db, err := gorm.Open("mysql", "touch:1314@/touch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("连接失败")
	}
	if db.HasTable(&Opition{}) {
		db.Create(&opition)
		return true
	} else {
		log.Println("选项表未创建，现在执行创建表操作")
		db.CreateTable(&Opition{})
		log.Println("选项表创建成功")
		db.Create(&opition)
		if db.NewRecord(opition) {

		} else {
			return true
		}
	}
	return false
}

