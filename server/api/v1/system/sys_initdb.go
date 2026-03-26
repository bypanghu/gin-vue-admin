package system

import (
	"errors"
	"path/filepath"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/consts"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/env"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/initial"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type DBApi struct{}

func failWithInitServiceError(c *gin.Context, err error, fallback string) {
	if errors.Is(err, request.ErrUnsupportedInitDBType) ||
		errors.Is(err, request.ErrRedisConfigRequired) ||
		errors.Is(err, request.ErrRedisHostRequired) ||
		errors.Is(err, request.ErrRedisPortRequired) {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.FailWithMessage(fallback, c)
}

// InitDB
// @Tags     InitDB
// @Summary  初始化用户数据库
// @Produce  application/json
// @Param    data  body      request.InitDB                  true  "初始化数据库参数"
// @Success  200   {object}  response.Response{data=string}  "初始化用户数据库"
// @Router   /init/initdb [post]
func (i *DBApi) InitDB(c *gin.Context) {
	if initial.HasInitial() {
		dataPath := env.GetGvaDataPath()
		lockFilePath := filepath.Join(dataPath, consts.INITIAL_LOCK_FILE_NAME)
		global.GVA_LOG.Error("服务器已初始化过! 如需重新初始化，请删除初始化标志文件后再进行初始化", zap.String("标志文件", lockFilePath))
		response.FailWithMessage("服务器已初始化过", c)
		return
	}
	var dbInfo request.InitDB
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("参数校验不通过!", zap.Error(err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := initDBService.InitDB(dbInfo); err != nil {
		global.GVA_LOG.Error("自动创建数据库失败!", zap.Error(err))
		failWithInitServiceError(c, err, "自动创建数据库失败，请查看后台日志，检查后在进行初始化")
		return
	}
	response.OkWithMessage("自动创建数据库成功", c)
}

// CheckDB
// @Tags     CheckDB
// @Summary  初始化用户数据库
// @Produce  application/json
// @Success  200  {object}  response.Response{data=map[string]interface{},msg=string}  "初始化用户数据库"
// @Router   /init/checkdb [post]
func (i *DBApi) CheckDB(c *gin.Context) {
	var (
		message  = "前往初始化数据库"
		needInit = true
	)

	if global.GVA_DB != nil {
		message = "数据库无需初始化"
		needInit = false
	}
	global.GVA_LOG.Info(message)
	response.OkWithDetailed(gin.H{"needInit": needInit}, message, c)
}

// TestDB
// @Tags     TestDB
// @Summary  测试数据库连接
// @Produce  application/json
// @Param    data  body      request.InitDB                  true  "测试数据库连接参数"
// @Success  200   {object}  response.Response{data=string}  "测试数据库连接"
// @Router   /init/testdb [post]
func (i *DBApi) TestDB(c *gin.Context) {
	var dbInfo request.TestDb
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("参数校验不通过!", zap.Error(err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := initDBService.TestDB(c.Request.Context(), dbInfo); err != nil {
		global.GVA_LOG.Error("数据库连接失败!", zap.Error(err))
		failWithInitServiceError(c, err, "数据库连接失败，请查看后台日志，检查后在进行测试")
		return
	}
	response.OkWithMessage("数据库连接成功", c)
}

// TestRedis
// @Tags     TestRedis
// @Summary  测试Redis连接
// @Produce  application/json
// @Param    data  body      request.InitDB                  true  "测试Redis连接参数"
// @Success  200   {object}  response.Response{data=string}  "测试Redis连接"
// @Router   /init/testredis [post]
func (i *DBApi) TestRedis(c *gin.Context) {
	var dbInfo request.InitRedis
	if err := c.ShouldBindJSON(&dbInfo); err != nil {
		global.GVA_LOG.Error("参数校验不通过!", zap.Error(err))
		response.FailWithMessage("参数校验不通过", c)
		return
	}
	if err := initDBService.TestRedis(c.Request.Context(), dbInfo); err != nil {
		global.GVA_LOG.Error("Redis连接失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("Redis连接成功", c)
}
