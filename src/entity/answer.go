package entity

import (
	"github.com/jinzhu/gorm"
	"log"
)

type Answer struct {
	gorm.Model
	QuestionID uint `gorm:"UNIQUE"`
	Answer     string
}

//AddAnswer 增加答案
func (as Answer) AddAnswer(answer Answer) bool {
	db, err := gorm.Open("mysql", "touch:1314@/touch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("连接失败")
	}
	if db.HasTable(&Answer{}) {
		db.Create(&answer)
		return true
	} else {
		log.Println("答案表未创建，现在执行创建表操作")
		db.CreateTable(&Answer{})
		log.Println("答案表创建成功")
		db.Create(&answer)
		if db.NewRecord(answer) {

		} else {
			return true
		}
	}
	return false
}

