package model

import (
	"fmt"
	"winstore/comm"
)

// Configs 配置表
type Configs struct {
	Key   string `json:"configKey"`
	Value string `json:"configValue"`
}

// SelectByKey 根据key查询配置
func (configs *Configs) SelectByKey() error {
	get, b := comm.AppCache.Get(configs.Key)
	if b {
		configs.Value = get.(string)
		return nil
	}
	return fmt.Errorf("未找到该配置")
}

func (configs *Configs) UpdateByKey() {
	comm.AppCache.Set(configs.Key, configs.Value, 0)
}

func (configs *Configs) InsertData() {
	comm.AppCache.Set(configs.Key, configs.Value, 0)
}
