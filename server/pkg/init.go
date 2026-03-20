package pkg

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/pkg/consts"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/env"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/initial"
	"github.com/gin-gonic/gin"
)

const prefix = "[github.com/flipped-aurora/gin-vue-admin/server] [pkg]"

func Init() {
	fmt.Printf("+------------------ %s ------------------+\n", prefix)
	fmt.Printf("%s gin-vue-admin pkg init success HasInitial=%v\n", prefix, initial.HasInitial())
	fmt.Printf("%s 后端 gin 运行模式 mode=%s\n", prefix, env.GetEnv(consts.ENV_GIN_MODE, gin.DebugMode))
	fmt.Printf("%s 数据目录 data path=%s\n", prefix, env.GetGvaDataPath())
	fmt.Printf("+------------------ %s ------------------+\n", prefix)
}
