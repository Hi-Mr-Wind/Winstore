package model

import (
	"time"
	"winstore/comm"
)

// Residue 软件卸载后残余数据
type Residue struct {
	// 残余数据路径
	ResiduePath string `json:"residue_path"`
	// 残余软件名称
	ResidueAppName string `json:"residue_app_name"`
}

// InsertData 插入残余软件数据
func (residue *Residue) InsertData() {
	comm.AppCache.Set(residue.ResidueAppName, residue.ResiduePath, 30*time.Hour)
}

// SelectAll 查询所有残余软件数据
func (residue *Residue) SelectAll() *[]Residue {
	var residueList []Residue
	for key, value := range comm.AppCache.GetAll() {
		residue := Residue{}
		residue.ResidueAppName = key
		residue.ResiduePath = value.(string)
		residueList = append(residueList, residue)
	}
	return &residueList
}

// DeleteByName 根据软件名称删除残余软件数据
func (residue *Residue) DeleteByName() error {
	return comm.AppCache.Delete(residue.ResidueAppName)
}
