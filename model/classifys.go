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
func (Classify) TableName() string {
	return "classifys"
}
