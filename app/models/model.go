package models

import "github.com/zhangguangying/goblog/pkg/types"

type BaseModel struct {
	ID uint64
}

func (a BaseModel) GetStringId() string {
	return types.Uint64ToString(a.ID)
}
