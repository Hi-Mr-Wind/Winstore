package model

// Collect 收藏表
type Collect struct {
	ID string `json:"id" gorm:"type:text;primaryKey;column:id"`
	// 收藏的应用名称
	AppName string `json:"appName" gorm:"type:text;not null;column:app_name;index"`
	// 收藏的应用图标
	DownloadURL *string `json:"downloadUrl,omitempty" gorm:"type:text;column:download_url"`
	// 收藏的应用地址
	AppURL *string `json:"appUrl,omitempty" gorm:"type:text;column:app_url"`
	//应用图标，可空
	AppIcon *string `json:"appIcon,omitempty" gorm:"type:text;column:app_icon"`
	// 创建时间
	CreateTime *string `json:"createTime,omitempty" gorm:"type:text;column:create_time"`
}

func (*Collect) TableName() string {
	return "collect"
}

// InsertData 插入收藏数据
func (collect *Collect) InsertData() error {
	return DB.Create(collect).Error
}
