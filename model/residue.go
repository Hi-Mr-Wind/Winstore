package model

// Residue 软件卸载后残余数据
type Residue struct {
	// 软件卸载后残余数据ID
	ID string `json:"id" gorm:"type:text;primaryKey;column:id"`
	// 残余数据路径
	ResiduePath string `json:"residue_path" gorm:"type:text;column:residue_path"`
	// 残余软件名称
	ResidueAppName string `json:"residue_app_name" gorm:"type:text;column:residue_app_name"`
	// 创建时间
	CreateTime string `json:"create_time" gorm:"type:text;column:create_time"`
}

func (*Residue) TableName() string {
	return "residue"
}

// InsertData 插入残余软件数据
func (residue *Residue) InsertData() error {
	return DB.Create(residue).Error
}

// SelectAll 查询所有残余软件数据
func (residue *Residue) SelectAll() (*[]Residue, error) {
	var residueList []Residue
	err := DB.Find(&residueList).Error
	return &residueList, err
}

// DeleteByName 根据软件名称删除残余软件数据
func (residue *Residue) DeleteByName() error {
	return DB.Where("residue_app_name = ?", residue.ResidueAppName).Delete(residue).Error
}
