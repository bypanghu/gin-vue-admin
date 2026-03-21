package core

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/core/internal"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/env"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap() (logger *zap.Logger) {
	zapLocalPath := filepath.Join(env.GetGvaDataPath(), global.GVA_CONFIG.Zap.Director)
	if ok, _ := utils.PathExists(zapLocalPath); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", zapLocalPath)
		_ = os.Mkdir(zapLocalPath, os.ModePerm)
	}
	levels := global.GVA_CONFIG.Zap.Levels()
	length := len(levels)
	cores := make([]zapcore.Core, 0, length)
	for i := 0; i < length; i++ {
		core := internal.NewZapCore(levels[i])
		cores = append(cores, core)
	}
	// 构建基础 logger（错误级别的入库逻辑已在自定义 ZapCore 中处理）
	logger = zap.New(zapcore.NewTee(cores...))
	opts := []zap.Option{zap.AddStacktrace(zapcore.ErrorLevel)}
	if global.GVA_CONFIG.Zap.ShowLine {
		opts = append(opts, zap.AddCaller())
	}
	logger = logger.WithOptions(opts...)
	return logger
}
