package entity

import (
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)
//Question 问题表
type Question struct {
	gorm.Model
	Name string `gorm:"UNIQUE","NOT NULL"`
	Options []Opition `gorm:"foreignkey:QuestionID"`
}

//AddQuestions 增加问题
func (qs Question) AddQuestions(question Question) bool {
	db, err := gorm.Open("mysql", "touch:1314@/touch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("连接失败")
	}
	if db.HasTable(&Question{}) {
		db.Create(&question)
		return true
	} else {
		log.Println("问题表未创建，现在执行创建表操作")
		db.CreateTable(&Question{})
		log.Println("问题表创建成功")
		db.Create(&question)
		if db.NewRecord(question) {

		} else {
			return true
		}
	}
	return false
}

func (qs Question) RadomOneQuestion() Question {
	db, err := gorm.Open("mysql", "touch:1314@/touch?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("连接失败")
	}
	var count int
	db.Table("questions").Count(&count)
	rand.Seed(time.Now().Unix())
	qid := rand.Intn(count) + 1
	var quetion Question
	db.First(&quetion,qid)
	db.Model(&quetion).Related(&quetion.Options)
	return quetion
}
