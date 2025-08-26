package model

// Classify 分类表模型
type Classify struct {
	// 分类id
	ClassifyID string `gorm:"column:classify_id;type:text;not null;uniqueIndex" json:"classify_id"`
	// 分类名称
	ClassifyName string `gorm:"column:classify_name;type:text;not null"           json:"classify_name"`
	// 父级分类
	ParentLevel *string `gorm:"column:parent_level;type:text;index"               json:"parent_level,omitempty"`
}

// TableName 指定表名
func (*Classify) TableName() string {
	return "classifys"
}

// SelectAll 查询所有分类
func (classify *Classify) SelectAll() (*[]Classify, error) {
	var classifys []Classify
	err := DB.Find(&classifys).Error
	return &classifys, err
}

// SelectSubClassificationById 根据父级分类id查询子级分类
func (classify *Classify) SelectSubClassificationById() (*[]Classify, error) {
	var classifys []Classify
	err := DB.Where("parent_level = ?", classify.ParentLevel).Find(&classify).Error
	return &classifys, err
}
