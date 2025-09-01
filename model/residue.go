package model

import (
	"time"
	"winstore/comm"
)

// 在init函数中注册Residue类型到gob
func init() {
	comm.RegisterType([]Residue{})
	comm.RegisterType(Residue{})
}

// Residue 软件卸载后残余数据
type Residue struct {
	// 残余数据路径
	ResiduePath string `json:"residue_path"`
	// 残余软件名称
	ResidueAppName string `json:"residue_app_name"`
}

// InsertData 插入残余软件数据
func (residue *Residue) InsertData() error {
	var residueList []Residue

	// 从缓存中获取现有残余数据列表
	if val, ok := comm.AppCache.Get("residue"); ok {
		if cachedList, ok := val.([]Residue); ok {
			residueList = cachedList
		}
	}

	// 添加新的残余数据项
	residueList = append(residueList, *residue)

	// 保存回缓存
	comm.AppCache.Set("residue", residueList, 24*30*time.Hour)
	return nil
}

// SelectAll 查询所有残余软件数据
func (residue *Residue) SelectAll() *[]Residue {
	var residueList []Residue

	// 从缓存中获取残余数据列表
	if val, ok := comm.AppCache.Get("residue"); ok {
		if cachedList, ok := val.([]Residue); ok {
			residueList = cachedList
		}
	}

	return &residueList
}

// DeleteByName 根据软件名称删除残余软件数据
func (residue *Residue) DeleteByName() error {
	var residueList []Residue

	// 从缓存中获取现有残余数据列表
	if val, ok := comm.AppCache.Get("residue"); ok {
		if cachedList, ok := val.([]Residue); ok {
			residueList = cachedList
		}
	}

	// 查找并删除指定名称的残余数据项
	for i, v := range residueList {
		if v.ResidueAppName == residue.ResidueAppName {
			residueList = append(residueList[:i], residueList[i+1:]...)
			break
		}
	}

	// 保存更新后的列表回缓存
	comm.AppCache.Set("residue", residueList, 24*30*time.Hour)
	return nil
}
