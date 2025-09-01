package model

// DatabaseVersion 数据库版本
type DatabaseVersion struct {
	Version    string `gorm:"column:version;type:text;not null"                json:"version"`
	UpdateTime string `gorm:"column:update_time;type:text;not null"          json:"update_time"`
}

func (*DatabaseVersion) TableName() string {
	return "database_version"
}

// GetVersion 获取数据库版本
func (databaseVersion *DatabaseVersion) GetVersion() (*DatabaseVersion, error) {
	err := DB.Where("version = ?", databaseVersion.Version).First(databaseVersion).Error
	return databaseVersion, err
}
