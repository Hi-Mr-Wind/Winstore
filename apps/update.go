package apps

import "winstore/model"

// GetDatabaseVersion 获取当前数据库版本
func (a *App) GetDatabaseVersion() (int64, error) {
	dbv := new(model.DatabaseVersion)
	err := dbv.GetVersion()
	if err != nil {
		return 0, err
	}
	return dbv.Version, nil
}

// GetNewDataBaseVersion 获取数据库版本
func (a *App) GetNewDataBaseVersion() (bool, error) {
	dbv := new(model.DatabaseVersion)
	return dbv.SelectUpdateDataBase(a.ctx)
}

// UpdateDatabase 更新软件数据库
func (a *App) UpdateDatabase() error {
	return model.DownloadNewDataBase(a.ctx)
}
