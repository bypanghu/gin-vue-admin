package env

import (
	"os"
	"strconv"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/pkg/consts"
	"go.uber.org/zap"
)

const ENV_PREFIX = "GVA"

// GetEnv 获取环境变量，如果环境变量不存在，则返回默认值
func GetEnv(envKey string, defaultValue string) string {
	if !strings.HasPrefix(envKey, ENV_PREFIX) {
		envKey = ENV_PREFIX + "_" + envKey
	}
	if value := os.Getenv(envKey); value != "" {
		return value
	}
	return defaultValue
}

// GetEnvAsInt 获取环境变量并转换为整数，如果环境变量不存在或转换失败，则返回默认值
func GetEnvAsInt(envKey string, defaultValue int) int {
	valueStr := GetEnv(envKey, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		zap.L().Error("环境变量转换为整数失败", zap.String("envKey", envKey), zap.String("valueStr", valueStr), zap.Error(err))
		return defaultValue
	}
	return value
}

// GetEnvAsInt64 获取环境变量并转换为 int64，如果环境变量不存在或转换失败，则返回默认值
func GetEnvAsInt64(envKey string, defaultValue int64) int64 {
	valueStr := GetEnv(envKey, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseInt(valueStr, 10, 64)
	if err != nil {
		zap.L().Error("环境变量转换为int64失败", zap.String("envKey", envKey), zap.String("valueStr", valueStr), zap.Error(err))
		return defaultValue
	}
	return value
}

// GetEnvAsUint64 获取环境变量并转换为 uint64，如果环境变量不存在或转换失败，则返回默认值
func GetEnvAsUint64(envKey string, defaultValue uint64) uint64 {
	valueStr := GetEnv(envKey, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseUint(valueStr, 10, 64)
	if err != nil {
		zap.L().Error("环境变量转换为uint64失败", zap.String("envKey", envKey), zap.String("valueStr", valueStr), zap.Error(err))
		return defaultValue
	}
	return value
}

// GetEnvAsBool 获取环境变量并转换为 bool，如果环境变量不存在或转换失败，则返回默认值
func GetEnvAsBool(envKey string, defaultValue bool) bool {
	valueStr := GetEnv(envKey, "")
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		zap.L().Error("环境变量转换为bool失败", zap.String("envKey", envKey), zap.String("valueStr", valueStr), zap.Error(err))
		return defaultValue
	}
	return value
}

// GetEnvAsSlice 获取以 separator 分隔的环境变量，如果环境变量为空，返回默认值
func GetEnvAsSlice(envKey string, defaultValue []string, separator string) []string {
	valueStr := strings.TrimSpace(GetEnv(envKey, ""))
	if valueStr == "" {
		return defaultValue
	}
	parts := strings.Split(valueStr, separator)
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item != "" {
			result = append(result, item)
		}
	}
	if len(result) == 0 {
		return defaultValue
	}
	return result
}

func GetGvaDataPath() string {
	dataPath := GetEnv(consts.ENV_GVA_DATA_PATH, consts.GVA_DATA_PATH)
	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		err := os.MkdirAll(dataPath, os.ModePerm)
		if err != nil {
			zap.L().Error("创建数据目录失败", zap.String("dataPath", dataPath), zap.Error(err))
			panic(err)
		}
	}
	return GetEnv(consts.ENV_GVA_DATA_PATH, consts.GVA_DATA_PATH)
}
