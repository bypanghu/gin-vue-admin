import service from '@/utils/request'
// @Tags InitDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Param data body request.InitDB true "初始化数据库参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"自动创建数据库成功"}"
// @Router /init/initdb [post]
export const initDB = (data) => {
  return service({
    url: '/init/initdb',
    method: 'post',
    data,
    donNotShowLoading: true
  })
}

// @Tags CheckDB
// @Summary 初始化用户数据库
// @Produce  application/json
// @Success 200 {string} string "{"code":0,"data":{},"msg":"探测完成"}"
// @Router /init/checkdb [post]
export const checkDB = () => {
  return service({
    url: '/init/checkdb',
    method: 'post'
  })
}

// @tags TestDb
// @Summary 测试数据库连接
// @Produce  application/json
// @Param data body request.TestDB true "测试数据库连接参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"连接成功"}"
// @Router /init/testdb [post]
export const testDB = (data) => {
  return service({
    url: '/init/testDb',
    method: 'post',
    data,
    donNotShowLoading: true
  })
}

// @tags TestRedis
// @Summary 测试Redis连接
// @Produce  application/json
// @Param data body request.TestRedis true "测试Redis连接参数"
// @Success 200 {string} string "{"code":0,"data":{},"msg":"连接成功"}"
// @Router /init/testredis [post]
export const testRedis = (data) => {
  return service({
    url: '/init/testRedis',
    method: 'post',
    data,
    donNotShowLoading: true
  })
}
