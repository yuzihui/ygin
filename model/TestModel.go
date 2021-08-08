package model

// TestModel 左侧菜单栏表
type TestModel struct {
	Id uint `gorm:"column:id" form:"id"` //主键
	Pid uint `gorm:"column:pid" form:"pid"` //父类ID
}



