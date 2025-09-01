package model

import (
	"winstore/comm"
)

// 在init函数中注册Collect类型到gob
func init() {
	comm.RegisterType([]Collect{})
	comm.RegisterType(Collect{})
	comm.RegisterType(comm.CacheItem{})
}

// Collect 收藏表
type Collect struct {
	ID string `json:"id"`
	// 收藏的应用名称
	AppName string `json:"appName"`
	// 收藏的应用下载地址
	DownloadURL *string `json:"downloadUrl,omitempty"`
	// 收藏的应用地址
	AppURL *string `json:"appUrl,omitempty"`
	//应用图标，可空
	AppIcon *string `json:"appIcon,omitempty"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty"`
}

// InsertData 插入收藏数据
func (collect *Collect) InsertData() error {
	var collectList []Collect

	// 从缓存中获取现有收藏列表
	if val, ok := comm.AppCache.Get("collect"); ok {
		if cachedList, ok := val.([]Collect); ok {
			collectList = cachedList
		}
	}

	// 添加新的收藏项
	collectList = append(collectList, *collect)

	// 保存回缓存
	comm.AppCache.Set("collect", collectList, 0)
	return nil
}

// SelectAll 查询所有收藏数据
func (collect *Collect) SelectAll() (*[]Collect, error) {
	var collectList []Collect

	// 从缓存中获取收藏列表
	if val, ok := comm.AppCache.Get("collect"); ok {
		if cachedList, ok := val.([]Collect); ok {
			collectList = cachedList
		}
	}

	return &collectList, nil
}

// DeleteByID 删除收藏数据
func (collect *Collect) DeleteByID() error {
	var collectList []Collect

	// 从缓存中获取现有收藏列表
	if val, ok := comm.AppCache.Get("collect"); ok {
		if cachedList, ok := val.([]Collect); ok {
			collectList = cachedList
		}
	}

	// 查找并删除指定ID的收藏项
	for i, v := range collectList {
		if v.ID == collect.ID {
			collectList = append(collectList[:i], collectList[i+1:]...)
			break
		}
	}

	// 保存更新后的列表回缓存
	comm.AppCache.Set("collect", collectList, 0)
	return nil
}

// UpdateById 更新收藏数据
func (collect *Collect) UpdateById() error {
	var collectList []Collect
	if val, ok := comm.AppCache.Get("collect"); ok {
		if cachedList, ok := val.([]Collect); ok {
			collectList = cachedList
		}
	}
	for i, v := range collectList {
		if v.ID == collect.ID {
			collectList[i] = *collect
			break
		}
	}
	comm.AppCache.Set("collect", collectList, 0)
	return nil
}
