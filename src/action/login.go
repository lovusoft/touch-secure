package action
import (
	"touch-secure/src/entity"
)
// Login 用户登录方法
func Login(user entity.User)  {
	user.FindUser(user)
}
