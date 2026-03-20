package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/env"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/file"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

// Viper 配置
func Viper() *viper.Viper {
	config := getConfigPath()

	v := viper.New()
	setViperDefaultsAndEnv(v)
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&global.GVA_CONFIG); err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	global.GVA_CONFIG.AutoCode.Root, _ = filepath.Abs("..")
	return v
}

func setViperDefaultsAndEnv(v *viper.Viper) {
	defaultConfig := global.DefaultConfig()
	defaultConfigMap := structToMap(defaultConfig)

	v.SetEnvPrefix(env.ENV_PREFIX)
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	applyMapToViper(defaultConfigMap, "", func(key string, value any) {
		v.SetDefault(key, value)
		_ = v.BindEnv(key)
	})
}

func structToMap(value any) map[string]any {
	raw, err := yaml.Marshal(value)
	if err != nil {
		panic(fmt.Errorf("marshal default config failed: %w", err))
	}

	var result map[string]any
	if err = yaml.Unmarshal(raw, &result); err != nil {
		panic(fmt.Errorf("unmarshal default config map failed: %w", err))
	}
	return result
}

func applyMapToViper(value map[string]any, parentKey string, apply func(key string, value any)) {
	for key, item := range value {
		currentKey := key
		if parentKey != "" {
			currentKey = parentKey + "." + key
		}

		switch v := item.(type) {
		case map[string]any:
			applyMapToViper(v, currentKey, apply)
		default:
			apply(currentKey, item)
		}
	}
}

// getConfigPath 获取配置文件路径, 优先级: 命令行 > 环境变量 > 默认值
func getConfigPath() (config string) {
	flag.StringVar(&config, "c", "", "choose config file.")
	flag.Parse()
	if config != "" {
		fmt.Printf("您正在使用命令行的 '-c' 参数传递的值, config 的路径为 %s\n", config)
		return
	}
	config = filepath.Join(env.GetGvaDataPath(), "config.yaml")
	_, err := os.Stat(config)
	if err != nil || os.IsNotExist(err) {
		fmt.Printf("未找到配置文件, 将创建默认配置文件")
		defaultConfig := generateDefaultConfig()
		err = file.WriteFileAtomic(config, []byte(defaultConfig), 0644)
		if err != nil {
			panic(fmt.Errorf("写入默认配置失败: %w", err))
		}
		fmt.Printf("默认配置文件已创建, config 的路径为 %s\n", config)
		return

	}

	return
}

// generateDefaultConfig 生成默认配置文件, 如果配置文件不存在则创建一个默认的配置文件
func generateDefaultConfig() string {
	global.GVA_CONFIG = global.DefaultConfig()
	// 将配置文件转换成 yml
	out, mErr := yaml.Marshal(global.GVA_CONFIG)
	if mErr != nil {
		panic(fmt.Errorf("生成默认配置失败: %w", mErr))
	}
	return string(out)
}
