package initial

import (
	"os"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/pkg/consts"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/env"
	"go.uber.org/zap"
)

// HasInitial 判断是否已经初始化过了, 默认读取顺序是 环境变量 > 配置文件 > 默认值, 只要其中一个有值就认为已经初始化过了
func HasInitial() bool {
	envHasInitial := env.GetEnv(consts.ENV_HAS_INITIAL, "false")
	if envHasInitial == "true" {
		return true
	}
	// 判断文件夹中是否有文件, 如果有则认为已经初始化过了
	dataPath := env.GetGvaDataPath()
	lockFilePath := filepath.Join(dataPath, consts.INITIAL_LOCK_FILE_NAME)
	fileStat, err := os.Stat(lockFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		zap.L().Error("判断是否已经初始化失败", zap.String("lockFilePath", lockFilePath), zap.Error(err))
		panic(err)
	}
	if fileStat.IsDir() {
		_ = os.Remove(lockFilePath)
		return false
	}
	return true
}

func AutoInitial() bool {
	if !HasInitial() {
		return env.GetEnvAsBool(consts.ENV_AUTO_INITIAL, true)
	}
	return false
}
