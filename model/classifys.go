package model

// Classify 分类表模型
type Classify struct {
	// 分类id
	ClassifyID string `gorm:"column:classify_id;type:text;not null;uniqueIndex" json:"classify_id"`
	// 分类名称
	ClassifyName string `gorm:"column:classify_name;type:text;not null"           json:"classify_name"`
	// 父级分类
	ParentLevel *string `gorm:"column:parent_level;type:text;index"               json:"parent_level,omitempty"`
	// 图标
	Icon *string `json:"icon,omitempty" gorm:"type:text;column:icon"`
	// 分类下软件数量
	Count int64 `json:"count" gorm:"type:int;column:count"`
}

// TableName 指定表名
func (*Classify) TableName() string {
	return "classifys"
}

// GetCount 获取分类软件数量
func getCount(classify string) (int64, error) {
	var count int64
	err := DB.
		Raw("select count(*) as count from app_list left join classifys as cl on app_list.classify = cl.classify_id where cl.classify_name = ? group by classify", classify).
		Scan(&count).Error
	return count, err
}

// SelectAll 查询所有分类
func (classify *Classify) SelectAll() (*[]Classify, error) {
	var classifys []Classify
	err := DB.Table("classifys").Select("classifys.*,count(*) as count").
		Joins("join app_list on app_list.classify = classifys.classify_id").
		Group("classifys.classify_id").Find(&classifys).Error
	return &classifys, err
}

// SelectSubClassificationById 根据父级分类id查询子级分类
func (classify *Classify) SelectSubClassificationById() (*[]Classify, error) {
	var classifys []Classify
	err := DB.Where("parent_level = ?", classify.ParentLevel).Find(&classify).Error
	return &classifys, err
}
