package requests

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/zhangguangying/goblog/app/models/user"
)

func ValidateRegistrationForm(data user.User) map[string][]string {
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:3,20"},
		"email":            []string{"required", "min:4", "max:30", "email"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}
	messages := govalidator.MapData{
		"name":             []string{"required:name 不能为空", "alpha_num:格式错误,只允许英文或数字", "between:3,20:用户名长度需再3-20之间"},
		"email":            []string{"required:email 不能为空", "min:Email 长度最小为4", "max:Email 长度最大为30", "email:Email格式不正确"},
		"password":         []string{"required:Paasword 不能为空", "min: Password 长度最小为6"},
		"password_confirm": []string{"required: password_confirm 不能为空"},
	}
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}
	errs := govalidator.New(opts).ValidateStruct()
	if data.Password != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次密码不匹配")
	}
	return errs
}
