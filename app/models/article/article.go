package article

import (
	"goblog/pkg/model"
	"goblog/pkg/types"
)

type Article struct {
	ID          uint64
	Title, Body string
}

func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToUInt64(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}
