package model

// Configs 配置表
type Configs struct {
	ID         string `gorm:"column:id;type:text;primaryKey" json:"id"`
	Key        string `gorm:"column:key;type:text;not null" json:"key"`
	Value      string `gorm:"column:value;type:text;not null" json:"value"`
	CreateTime string `gorm:"column:create_time;type:text;not null" json:"create_time"`
}

func (*Configs) TableName() string {
	return "configs"
}

// SelectByKey 根据key查询配置
func (configs *Configs) SelectByKey() error {
	err := DB.Where("key = ?", configs.Key).Find(configs).Error
	if err != nil {
		return err
	}
	return nil
}
