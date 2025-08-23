package model

import "time"

// AppImg 应用图片
type AppImg struct {
	ID string `json:"id" gorm:"type:text;primaryKey;column:id"`
	// 应用ID
	AppID string `json:"appId" gorm:"type:text;not null;column:app_id;index"`
	// 图片URL
	ImgURL     string    `json:"imgUrl" gorm:"type:text;not null;column:img_url"`
	CreateTime time.Time `json:"createTime" gorm:"not null;column:create_time"`
}

// TableName 指定表名
func (*AppImg) TableName() string {
	return "app_imgs"
}

// GetImgByAppID 根据应用ID获取图片
func (appImg *AppImg) GetImgByAppID() (*[]AppImg, error) {
	imgs := new([]AppImg)
	err := DB.Where("app_id = ?", appImg.AppID).Find(imgs).Error
	return imgs, err
}
