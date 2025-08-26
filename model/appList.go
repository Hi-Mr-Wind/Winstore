package model

import (
	"errors"

	"github.com/google/uuid"
)

type AppList struct {
	// appid
	AppID string `gorm:"column:app_id;type:text;not null;primaryKey"         json:"app_id"`
	// 软件名称
	AppName string `gorm:"column:app_name;type:text;not null"       json:"app_name"`
	// 软件图标
	AppIcon *string `gorm:"column:app_icon;type:text"                json:"app_icon,omitempty"`
	// 下载地址
	DownloadURL string `gorm:"column:download_url;type:text;not null"   json:"download_url"`
	// 官网
	OfficialWebsite *string `gorm:"column:official_website;type:text"        json:"official_website,omitempty"`
	// 制作公司
	Company *string `gorm:"column:company;type:text"                 json:"company,omitempty"`
	// 软件版本
	AppVersion *string `gorm:"column:app_version;type:text"             json:"app_version,omitempty"`
	// 排序
	Sort *int `gorm:"column:sort;type:integer"                 json:"sort,omitempty"`
	// 下载次数
	DownloadCount *int `gorm:"column:download_count;type:integer"       json:"download_count,omitempty"`
	// 更新时间
	UpdateTime *string `gorm:"column:update_time;type:text"             json:"update_time,omitempty"`
	// 软件分类
	Classify *string `gorm:"column:classify;type:text"                 json:"classify,omitempty"`
	// 软件简介
	BriefIntroduction *string `gorm:"column:brief_introduction;type:text"        json:"brief_introduction,omitempty"`
}

// TableName 指定表名
func (*AppList) TableName() string {
	return "app_list"
}

// InsertData 插入软件列表数据
func (appList *AppList) InsertData() error {
	if appList == nil {
		return errors.New("appList it cannot be null")
	}
	v7, err := uuid.NewV7()
	if err != nil {
		return err
	}
	appList.AppID = v7.String()
	return DB.Create(appList).Error
}

// SelectByID 根据id查询软件列表数据
func (appList *AppList) SelectByID(id string) error {
	return DB.First(appList, id).Error
}

// SelectAllPage 分页查询软件列表数据
func (appList *AppList) SelectAllPage(pageNum int, pageSize int, classify string) (*[]AppList, error) {
	var appLists = new([]AppList)
	tx := DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Order("sort")
	if classify != "" {
		err := tx.Where("classify = ?", classify).Find(appLists).Error
		return appLists, err
	} else {
		err := tx.Find(appLists).Error
		return appLists, err
	}
}

// UpdateData 更新软件列表数据
// data:要更新的数据
// 返回值：受影响的行数
func (appList *AppList) UpdateData(data *[]AppList) (int64, error) {
	res := DB.Model(appList).Where("1 = 1").Updates(data)
	return res.RowsAffected, res.Error
}

// DeleteByID 删除软件列表数据
func (appList *AppList) DeleteByID() error {
	return DB.Delete(appList).Error
}

// SelectByAppName 根据软件名称模糊查询软件列表数据
func (appList *AppList) SelectByAppName() (*[]AppList, error) {
	var appLists = new([]AppList)
	err := DB.Where("app_name like %?%", appList.AppName).Find(appLists).Error
	return appLists, err
}
