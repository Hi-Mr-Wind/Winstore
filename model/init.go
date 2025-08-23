package model

import (
	"context"
	"os"
	"winstore/log"

	"github.com/glebarez/sqlite"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

var DB *gorm.DB

// DBinit 初始化数据库
func DBinit(c context.Context) {
	gormLogger := log.NewCustomGormLogger(c)
	db, err := gorm.Open(sqlite.Open("./lib/winstore.db"), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		runtime.LogErrorf(c, "数据加载失败：%s", err.Error())
		os.Exit(1)
	}
	DB = db
}
