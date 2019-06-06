package action

import (
	"log"
	"touch-secure/src/entity"
)

// Join 新增用户注册方法
func Join(user entity.User) bool {
	name := user.Name
	email := user.Email
	if (entity.User{}.FindUserName(name)) {
		log.Println("用户名已存在")
		return false
	} else if (entity.User{}.FindUserEmail(email)) {
		log.Println("用户邮箱已存在")
		return false
	} else {
		return entity.User{}.AddUser(user)
	}
}

func Answer() {

}
