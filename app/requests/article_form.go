package requests

import (
	"github.com/thedevsaddam/govalidator"
	"github.com/zhangguangying/goblog/app/models/article"
)

func ValidatorArticleForm(data article.Article) map[string][]string {
	rules := govalidator.MapData{
		"title": []string{"required", "min:3", "max:40"},
		"body":  []string{"required", "min:10"},
	}
	messages := govalidator.MapData{
		"title": []string{
			"required:标题不能为空",
			"min:标题不能小于3个字符",
			"max:标题不能大于40个字符",
		},
		"body": []string{
			"required:内容不能为空",
			"min:内容最小为10个字符",
		},
	}
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid",
	}
	return govalidator.New(opts).ValidateStruct()
}
