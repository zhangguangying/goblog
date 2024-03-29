package requests

import (
	"goblog/app/models/user"

	"github.com/thedevsaddam/govalidator"
)

func ValidateRegistrationForm(data user.User) map[string][]string {
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:3,20"},
		"email":            []string{"required", "min:4", "max:30", "email"},
		"password":         []string{"required", "min:6"},
		"confirm_password": []string{"required"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:格式错误，只允许数字和英文",
			"between:用户名长度需在 3~20 之间",
		},
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
		},
		"password": []string{
			"required:密码为必填项",
			"min:长度需大于 6",
		},
		"confirm_password": []string{
			"required:确认密码框为必填项",
		},
	}
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}
	errs := govalidator.New(opts).ValidateStruct()
	if data.Password != data.ConfirmPassword {
		errs["confirm_password"] = append(errs["confirm_password"], "两次输入密码不匹配！")
	}
	return errs
}
