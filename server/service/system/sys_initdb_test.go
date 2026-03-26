package system

import (
	"context"
	"errors"
	"testing"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/consts"
)

func TestInitHandlerForDBTypeRejectsUnsupportedType(t *testing.T) {
	_, _, err := initHandlerForDBType("oracle")
	if !errors.Is(err, request.ErrUnsupportedInitDBType) {
		t.Fatalf("expected unsupported db type error, got %v", err)
	}
}

func TestBuildRedisConfigUsesInitPayload(t *testing.T) {
	original := global.GVA_CONFIG
	t.Cleanup(func() {
		global.GVA_CONFIG = original
	})

	global.GVA_CONFIG = config.Server{
		Redis: config.Redis{
			Name: "legacy",
			Addr: "127.0.0.1:6379",
		},
	}

	redisCfg, useRedis := buildRedisConfig(&request.InitRedis{
		Host:     "10.0.0.8",
		Port:     "6380",
		Password: "redis-pass",
		DB:       4,
		Enable:   true,
	})

	if !useRedis {
		t.Fatalf("expected redis to be enabled")
	}
	if redisCfg.Name != consts.REIDS_NAME_DEFAULT {
		t.Fatalf("expected redis name %q, got %q", consts.REIDS_NAME_DEFAULT, redisCfg.Name)
	}
	if redisCfg.Addr != "10.0.0.8:6380" {
		t.Fatalf("expected redis addr to come from init payload, got %q", redisCfg.Addr)
	}
	if redisCfg.Password != "redis-pass" || redisCfg.DB != 4 {
		t.Fatalf("expected redis password/db to come from init payload, got %+v", redisCfg)
	}
}

func TestApplySharedInitConfigPersistsRedisEnablement(t *testing.T) {
	original := global.GVA_CONFIG
	t.Cleanup(func() {
		global.GVA_CONFIG = original
	})

	global.GVA_CONFIG = config.Server{}

	ctx := context.WithValue(context.Background(), "redisConfig", config.Redis{
		Name:     consts.REIDS_NAME_DEFAULT,
		Addr:     "10.0.0.9:6379",
		Password: "secret",
		DB:       2,
	})
	ctx = context.WithValue(ctx, "useRedis", true)

	applySharedInitConfig(ctx)

	if !global.GVA_CONFIG.System.UseRedis {
		t.Fatalf("expected system.use-redis to be true")
	}
	if global.GVA_CONFIG.Redis.Addr != "10.0.0.9:6379" {
		t.Fatalf("expected redis addr to be persisted, got %q", global.GVA_CONFIG.Redis.Addr)
	}
	if global.GVA_CONFIG.Redis.Password != "secret" || global.GVA_CONFIG.Redis.DB != 2 {
		t.Fatalf("expected redis credentials/db to be persisted, got %+v", global.GVA_CONFIG.Redis)
	}
}
