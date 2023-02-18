package article

import (
	"goblog/pkg/model"
	"goblog/pkg/route"
	"goblog/pkg/types"
	"strconv"
)

type Article struct {
	ID          uint64
	Title, Body string
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatUint(a.ID, 10))
}

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUInt64(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

func GetAll() ([]Article, error) {
	var articles []Article
	if err := model.DB.Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}
