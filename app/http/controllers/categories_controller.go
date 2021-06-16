package controllers

import (
	"fmt"
	"github.com/zhangguangying/goblog/app/models/article"
	"github.com/zhangguangying/goblog/app/models/category"
	"github.com/zhangguangying/goblog/app/requests"
	"github.com/zhangguangying/goblog/pkg/route"
	"github.com/zhangguangying/goblog/pkg/view"
	"gorm.io/gorm"
	"net/http"
)

type CategoryController struct {
}

func (*CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

func (*CategoryController) Store(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化数据
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	// 2. 表单验证
	errors := requests.ValidateCategoryForm(_category)

	// 3. 检测错误
	if len(errors) == 0 {
		// 创建文章分类
		_category.Create()
		if _category.ID > 0 {
			fmt.Fprint(w, "创建成功！")
			// indexURL := route.Name2URL("categories.show", "id", _category.GetStringID())
			// http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章分类失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   errors,
		}, "categories.create")
	}
}

func (cc *CategoryController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 2. 读取对应的数据
	_category, err := category.Get(id)

	// 3. 获取结果集
	articles, pagerData, err := article.GetByCategoryID(_category.GetStringId(), r, 10)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "记录不存在")
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "服务器内部错误")
		}
	} else {
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}
