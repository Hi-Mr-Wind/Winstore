package apps

import (
	"fmt"
	"time"
	"winstore/model"

	"github.com/google/uuid"
)

// CollectInsert 添加收藏
func (a *App) CollectInsert(collect *model.Collect) error {
	fmt.Printf("CollectInsert:%#v", *collect)
	collect.ID = uuid.New().String()
	createTime := time.Now().Format("2006-01-02 15:04:05")
	collect.CreateTime = &createTime
	return collect.InsertData()
}

// CollectSelectAll 查询所有收藏
func (a *App) CollectSelectAll() (*[]model.Collect, error) {
	collect := &model.Collect{}
	return collect.SelectAll()
}

// CollectDelete 删除收藏
func (a *App) CollectDelete(id string) error {
	collect := &model.Collect{}
	collect.ID = id
	return collect.DeleteByID()
}

// CollectUpdate 修改收藏
func (a *App) CollectUpdate(collect *model.Collect) error {
	fmt.Printf("CollectUpdate:%#v", *collect)
	return collect.UpdateById()
}
