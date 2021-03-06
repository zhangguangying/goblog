package auth

import (
	"errors"
	"github.com/zhangguangying/goblog/app/models/user"
	"github.com/zhangguangying/goblog/pkg/session"
	"gorm.io/gorm"
)

func _getUID() string {
	_uid := session.Get("uid")
	uid, ok := _uid.(string)
	if ok && len(uid) > 0 {
		return uid
	}
	return ""
}

func User() user.User {
	uid := _getUID()
	if len(uid) > 0 {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		}
	}
	return user.User{}
}

func Attempt(email, password string) error {
	_user, err := user.GetByEmail(email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("账号不存在")
		} else {
			return errors.New("内部错误，请稍后重试")
		}
	}
	if !_user.ComparePassword(password) {
		return errors.New("账号不存在或密码错误")
	}
	session.Put("uid", _user.GetStringId())
	return nil
}

func Login(_user user.User) {
	session.Put("uid", _user.GetStringId())
}

func Logout() {
	session.Forget("uid")
}

func Check() bool {
	return len(_getUID()) > 0
}
