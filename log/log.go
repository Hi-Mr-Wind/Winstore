package log

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm/logger"
)

// CustomGormLogger 自定义GORM日志记录器
type CustomGormLogger struct {
	ctx context.Context
}

// NewCustomGormLogger 创建新的GORM日志记录器实例
func NewCustomGormLogger(ctx context.Context) *CustomGormLogger {
	return &CustomGormLogger{ctx: ctx}
}

// LogMode 设置日志级别
func (l *CustomGormLogger) LogMode(level logger.LogLevel) logger.Interface {
	return l
}

// Info 打印信息级别日志
func (l *CustomGormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	runtime.LogInfo(l.ctx, fmt.Sprintf("gorm_log INFO %s %v", msg, data))
}

// Warn 打印警告级别日志
func (l *CustomGormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	runtime.LogWarning(l.ctx, fmt.Sprintf("gorm_log WARN %s %v", msg, data))
}

// Error 打印错误级别日志
func (l *CustomGormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	runtime.LogError(l.ctx, fmt.Sprintf("gorm_log ERROR %s %v", msg, data))
}

// Trace 打印SQL跟踪日志
func (l *CustomGormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	parts := []string{}
	if err != nil {
		parts = append(parts, fmt.Sprintf("Error: %v", err))
	}
	parts = append(parts, fmt.Sprintf("SQL: %s", sql))
	parts = append(parts, fmt.Sprintf("Rows: %d", rows))
	parts = append(parts, fmt.Sprintf("Time: %v", elapsed))

	message := strings.Join(parts, " | ")
	runtime.LogDebug(l.ctx, fmt.Sprintf("gorm_log INFO %s", message))
}
