package model

import (
	"encoding/json/v2"
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
func (residue *Residue) InsertData() error {
	get, b := comm.AppCache.Get("residue")
	var residueList []Residue
	if !b {
		residueList = append(residueList, *residue)
	} else {
		err := json.Unmarshal([]byte(get.(string)), &residueList)
		if err != nil {
			return err
		}
		residueList = append(residueList, *residue)
	}
	marshal, err := json.Marshal(residueList)
	if err != nil {
		return err
	}
	comm.AppCache.Set("residue", string(marshal), 30*time.Hour)
	return nil
}

// SelectAll 查询所有残余软件数据
func (residue *Residue) SelectAll() *[]Residue {
	var residueList []Residue
	for key, value := range comm.AppCache.GetAll() {
		if key != "residue" {
			continue
		}
		value := value.(string)
		err := json.Unmarshal([]byte(value), &residueList)
		if err != nil {
			return nil
		}
	}
	return &residueList
}

// DeleteByName 根据软件名称删除残余软件数据
func (residue *Residue) DeleteByName() error {
	get, b := comm.AppCache.Get("residue")
	if !b {
		return nil
	}
	var residueList []Residue
	err := json.Unmarshal([]byte(get.(string)), &residueList)
	if err != nil {
		return err
	}
	for i, v := range residueList {
		if v.ResidueAppName == residue.ResidueAppName {
			residueList = append(residueList[:i], residueList[i+1:]...)
			break
		}
	}
	marshal, err := json.Marshal(residueList)
	if err != nil {
		return err
	}
	comm.AppCache.Set("residue", string(marshal), 30*time.Hour)
	return nil
}
