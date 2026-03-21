<template>
  <div class="init-page relative min-h-screen overflow-hidden">
    <div class="pointer-events-none absolute inset-0">
      <div class="init-page__glow init-page__glow--primary" />
      <div class="init-page__glow init-page__glow--secondary" />
      <div class="init-page__grid" />
    </div>

    <div class="relative z-10 min-h-screen px-5 py-6 md:px-8 md:py-8 xl:px-12">
      <div
        class="mx-auto flex min-h-[calc(100vh-48px)] max-w-5xl items-center justify-center"
      >
        <section class="init-shell w-full">
          <div class="mb-8 flex flex-col items-center text-center">
            <div
              class="flex h-[72px] w-[72px] items-center justify-center rounded-[24px] text-slate-900 border border-[var(--el-color-primary-light-9)]"
            >
              <Logo :size="7" />
            </div>
            <h1
              class="mt-5 text-[36px] font-semibold tracking-tight text-sky-500"
            >
              初始化系统
            </h1>
            <p class="mt-2 text-sm text-slate-400">
              按步骤完成环境确认、数据库配置与系统设置
            </p>
          </div>

          <div class="init-panel rounded-[28px] p-6 sm:p-8 lg:p-10">
            <div class="mb-8 grid gap-3 md:grid-cols-2 xl:grid-cols-4">
              <button
                v-for="(step, index) in steps"
                :key="step.title"
                type="button"
                class="init-step-card text-left"
                :class="{
                  'is-active': currentStep === index,
                  'is-completed': currentStep > index
                }"
                @click="goStep(index)"
              >
                <span class="init-step-card__index">0{{ index + 1 }}</span>
                <h2 class="init-step-card__title">{{ step.title }}</h2>
                <p class="init-step-card__desc">{{ step.desc }}</p>
              </button>
            </div>

            <el-form :model="form" label-position="top" size="large">
              <section v-show="currentStep === 0" class="init-section">
                <div class="init-section__head">
                  <p class="init-section__eyebrow">STEP 1</p>
                  <h2 class="init-section__title">初始化提示</h2>
                  <p class="init-section__desc">
                    开始前先确认环境、文档和数据库引擎配置，避免初始化过程中出现可预期的问题。
                  </p>
                </div>
                <div class="grid gap-3">
                  <div class="init-tip-card">
                    请确认你已经具备基础的 Vue 和 Golang 使用经验。
                  </div>
                  <div class="init-tip-card">
                    初始化前建议先阅读
                    <a
                      class="init-link"
                      href="https://www.gin-vue-admin.com"
                      target="_blank"
                    >
                      官方文档
                    </a>
                    与
                    <a
                      class="init-link"
                      href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=2"
                      target="_blank"
                    >
                      初始化视频
                    </a>
                    。
                  </div>
                  <div class="init-tip-card">
                    如果你使用 MySQL，请确认数据库引擎为
                    <span class="font-semibold text-rose-500">InnoDB</span>。
                  </div>
                </div>
                <div class="mt-4 flex justify-end">
                  <el-button
                    class="init-doc-btn"
                    plain
                    size="large"
                    @click="goDoc"
                  >
                    阅读环境文档
                  </el-button>
                </div>
              </section>

              <section v-show="currentStep === 1" class="init-section">
                <div class="init-section__head">
                  <p class="init-section__eyebrow">STEP 2</p>
                  <h2 class="init-section__title">数据库设置</h2>
                  <p class="init-section__desc">
                    先选择数据库类型，再填写数据库连接参数。下方测试连接按钮当前只做前端参数检查。
                  </p>
                </div>

                <div class="grid gap-5 md:grid-cols-2">
                  <el-form-item label="数据库类型" class="mb-0 col-span-full">
                    <el-select
                      v-model="form.dbType"
                      placeholder="请选择数据库类型"
                      class="w-full init-input"
                      @change="changeDB"
                    >
                      <el-option key="mysql" label="mysql" value="mysql" />
                      <el-option key="pgsql" label="pgsql" value="pgsql" />
                      <el-option key="oracle" label="oracle" value="oracle" />
                      <el-option key="mssql" label="mssql" value="mssql" />
                      <el-option key="sqlite" label="sqlite" value="sqlite" />
                    </el-select>
                  </el-form-item>
                  <el-form-item label="数据库名称" class="mb-0">
                    <el-input
                      v-model="form.dbName"
                      placeholder="请输入数据库名称"
                      class="init-input"
                    />
                  </el-form-item>
                  <el-form-item
                    v-if="form.dbType !== 'sqlite'"
                    label="Host"
                    class="mb-0"
                  >
                    <el-input
                      v-model="form.host"
                      placeholder="请输入数据库地址"
                      class="init-input"
                    />
                  </el-form-item>
                  <el-form-item
                    v-if="form.dbType !== 'sqlite'"
                    label="Port"
                    class="mb-0"
                  >
                    <el-input
                      v-model="form.port"
                      placeholder="请输入数据库端口"
                      class="init-input"
                    />
                  </el-form-item>
                  <el-form-item
                    v-if="form.dbType !== 'sqlite'"
                    label="Username"
                    class="mb-0"
                  >
                    <el-input
                      v-model="form.userName"
                      placeholder="请输入数据库用户名"
                      class="init-input"
                    />
                  </el-form-item>
                  <el-form-item
                    v-if="form.dbType !== 'sqlite'"
                    label="Password"
                    class="mb-0"
                  >
                    <el-input
                      v-model="form.password"
                      show-password
                      placeholder="请输入数据库密码，没有可留空"
                      class="init-input"
                    />
                  </el-form-item>
                  <el-form-item
                    v-if="form.dbType === 'sqlite'"
                    label="Database Path"
                    class="mb-0"
                  >
                    <el-input
                      v-model="form.dbPath"
                      placeholder="请输入 sqlite 数据库存放路径"
                      class="init-input"
                    />
                  </el-form-item>
                  <el-form-item
                    v-if="form.dbType === 'pgsql'"
                    label="Template"
                    class="mb-0"
                  >
                    <el-input
                      v-model="form.template"
                      placeholder="请输入 PostgreSQL template"
                      class="init-input"
                    />
                  </el-form-item>
                </div>

                <div class="mt-6 flex justify-end">
                  <el-button
                    class="init-test-btn h-12 min-w-[140px]"
                    plain
                    size="large"
                    @click="testConnection"
                  >
                    测试连接
                  </el-button>
                </div>
              </section>

              <section v-show="currentStep === 2" class="init-section">
                <div class="init-section__head">
                  <p class="init-section__eyebrow">STEP 3</p>
                  <h2 class="init-section__title">Redis 配置</h2>
                  <p class="init-section__desc">
                    Redis 作为可选能力单独配置。启用后请填写连接信息。
                  </p>
                </div>

                <div class="grid gap-5 md:grid-cols-2">
                  <el-form-item label="启用 Redis" class="mb-0">
                    <div class="init-switch-wrap">
                      <el-switch v-model="form.redisEnable" />
                      <span class="init-switch-label">
                        {{ form.redisEnable ? '已启用' : '未启用' }}
                      </span>
                    </div>
                  </el-form-item>

                  <template v-if="form.redisEnable">
                    <el-form-item label="Redis Host" class="mb-0">
                      <el-input
                        v-model="form.redisHost"
                        placeholder="例如 127.0.0.1"
                        class="init-input"
                      />
                    </el-form-item>
                    <el-form-item label="Redis Port" class="mb-0">
                      <el-input
                        v-model="form.redisPort"
                        placeholder="例如 6379"
                        class="init-input"
                      />
                    </el-form-item>
                    <el-form-item label="Redis 密码" class="mb-0">
                      <el-input
                        v-model="form.redisPassword"
                        show-password
                        placeholder="没有密码可留空"
                        class="init-input"
                      />
                    </el-form-item>
                    <el-form-item label="Redis 库" class="mb-0">
                      <el-input-number
                        v-model="form.redisDB"
                        :min="0"
                        :max="16"
                        class="w-full init-input-number"
                      />
                    </el-form-item>
                  </template>
                </div>

                <div class="mt-6 flex justify-end">
                  <el-button
                    class="init-test-btn h-12 min-w-[140px]"
                    plain
                    size="large"
                    @click="testRedisConnection"
                  >
                    测试 Redis
                  </el-button>
                </div>
              </section>

              <section v-show="currentStep === 3" class="init-section">
                <div class="init-section__head">
                  <p class="init-section__eyebrow">STEP 4</p>
                  <h2 class="init-section__title">系统设置</h2>
                  <p class="init-section__desc">
                    设置管理员密码并进行二次确认，这里只做前端一致性校验。
                  </p>
                </div>

                <div class="grid gap-5 md:grid-cols-2">
                  <el-form-item label="管理员密码" class="mb-0">
                    <el-input
                      v-model="form.adminPassword"
                      show-password
                      placeholder="请输入管理员密码"
                      class="init-input"
                    />
                  </el-form-item>
                  <el-form-item label="确认管理员密码" class="mb-0">
                    <el-input
                      v-model="form.confirmAdminPassword"
                      show-password
                      placeholder="请再次输入管理员密码"
                      class="init-input"
                    />
                  </el-form-item>
                </div>
              </section>

              <div class="mt-8 flex items-center justify-between gap-3">
                <el-button
                  class="!rounded-full !px-12"
                  plain
                  size="large"
                  :disabled="currentStep === 0"
                  @click="prevStep"
                >
                  上一步
                </el-button>
                <el-button
                  v-if="currentStep < steps.length - 1"
                  class="!rounded-full !px-12"
                  type="primary"
                  size="large"
                  @click="nextStep"
                  :disabled="
                    (currentStep === 1 && !initialState.dbCheck) ||
                    (currentStep === 2 &&
                      form.redisEnable &&
                      !initialState.redisCheck)
                  "
                >
                  下一步
                </el-button>
                <el-button
                  v-else
                  class="init-submit h-12 min-w-[160px] border-0 text-base font-semibold"
                  type="primary"
                  size="large"
                  @click="onSubmit"
                >
                  立即初始化
                </el-button>
              </div>
            </el-form>
          </div>
        </section>
      </div>
    </div>
  </div>
</template>

<script setup>
  // @ts-ignore
  import { initDB, testDB, testRedis } from '@/api/initdb'
  import { reactive, ref } from 'vue'
  import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
  import { useRouter } from 'vue-router'
  import Logo from '@/components/logo/index.vue'

  defineOptions({
    name: 'Init'
  })

  const router = useRouter()
  const currentStep = ref(0)
  const steps = [
    {
      title: '提示说明',
      desc: '阅读初始化前的环境与文档提示'
    },
    {
      title: '数据库设置',
      desc: '配置数据库类型与连接参数'
    },
    {
      title: 'Redis 配置',
      desc: '按需启用 Redis 并填写参数'
    },
    {
      title: '系统设置',
      desc: '设置管理员密码并确认'
    }
  ]

  const form = reactive({
    adminPassword: '',
    confirmAdminPassword: '',
    redisEnable: false,
    redisHost: '127.0.0.1',
    redisPort: '6379',
    redisPassword: '',
    redisDB: 0,
    dbType: 'mysql',
    host: '127.0.0.1',
    port: '3306',
    userName: 'root',
    password: '',
    dbName: 'gva',
    dbPath: '',
    template: ''
  })

  const initialState = reactive({
    dbCheck: false,
    redisCheck: false
  })

  const goDoc = () => {
    window.open('https://www.gin-vue-admin.com/guide/start-quickly/env.html')
  }

  const goStep = (index) => {
    currentStep.value = index
  }

  const nextStep = () => {
    if (currentStep.value === 1 && !validateDatabaseStep()) {
      return
    }

    if (currentStep.value === 2 && !validateRedisStep()) {
      return
    }

    if (currentStep.value < steps.length - 1) {
      currentStep.value += 1
    }
  }

  const prevStep = () => {
    if (currentStep.value > 0) {
      currentStep.value -= 1
    }
  }

  const validateDatabaseStep = () => {
    if (!form.dbType) {
      ElMessage({
        type: 'error',
        message: '请选择数据库类型'
      })
      return false
    }

    if (!form.dbName) {
      ElMessage({
        type: 'error',
        message: '请输入数据库名称'
      })
      return false
    }

    if (form.dbType === 'sqlite') {
      return true
    }

    const requiredFields = [
      ['host', '请输入数据库地址'],
      ['port', '请输入数据库端口'],
      ['userName', '请输入数据库用户名']
    ]

    for (const [key, message] of requiredFields) {
      if (!form[key]) {
        ElMessage({
          type: 'error',
          message
        })
        return false
      }
    }

    if (form.dbType === 'pgsql' && !form.template) {
      ElMessage({
        type: 'error',
        message: '请输入 PostgreSQL template'
      })
      return false
    }

    return true
  }

  const validateSystemStep = () => {
    if (!form.adminPassword) {
      ElMessage({
        type: 'error',
        message: '请输入管理员密码'
      })
      return false
    }

    if (form.adminPassword.length < 6) {
      ElMessage({
        type: 'error',
        message: '密码长度不能小于6位'
      })
      return false
    }

    if (!form.confirmAdminPassword) {
      ElMessage({
        type: 'error',
        message: '请再次输入管理员密码'
      })
      return false
    }

    if (form.adminPassword !== form.confirmAdminPassword) {
      ElMessage({
        type: 'error',
        message: '两次输入的管理员密码不一致'
      })
      return false
    }

    return true
  }

  const validateRedisStep = () => {
    if (form.redisEnable && !form.redisHost) {
      ElMessage({
        type: 'error',
        message: '启用 Redis 后请输入 Redis Host'
      })
      return false
    }

    if (form.redisEnable && !form.redisPort) {
      ElMessage({
        type: 'error',
        message: '启用 Redis 后请输入 Redis Port'
      })
      return false
    }

    return true
  }

  const testConnection = async () => {
    if (!validateDatabaseStep()) {
      return
    }

    const loading = ElLoading.service({
      lock: true,
      text: '正在测试数据库连接，请稍候',
      spinner: 'loading',
      background: 'rgba(15, 23, 42, 0.46)'
    })

    try {
      const payload = {
        dbType: form.dbType,
        host: form.host,
        port: form.port,
        userName: form.userName,
        password: form.password,
        dbName: form.dbName,
        dbPath: form.dbPath,
        template: form.template
      }
      const res = await testDB(payload)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: res.msg || '数据库连接成功'
        })
        initialState.dbCheck = true
      }
    } finally {
      loading.close()
    }
  }

  const testRedisConnection = async () => {
    if (!validateRedisStep()) {
      return
    }

    if (!form.redisEnable) {
      ElMessage({
        type: 'info',
        message: '当前未启用 Redis'
      })
      return
    }

    const loading = ElLoading.service({
      lock: true,
      text: '正在测试 Redis 连接，请稍候',
      spinner: 'loading',
      background: 'rgba(15, 23, 42, 0.46)'
    })

    try {
      const res = await testRedis({
        host: form.redisHost,
        port: form.redisPort,
        password: form.redisPassword,
        db: form.redisDB,
        enable: form.redisEnable
      })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: res.msg || 'Redis 连接成功'
        })
        initialState.redisCheck = true
      }
    } finally {
      loading.close()
    }
  }

  const changeDB = (val) => {
    switch (val) {
      case 'mysql':
        Object.assign(form, {
          adminPassword: '123456',
          confirmAdminPassword: '123456',
          redisEnable: form.redisEnable,
          redisHost: form.redisHost,
          redisPort: form.redisPort,
          redisPassword: form.redisPassword,
          redisDB: form.redisDB,
          dbType: 'mysql',
          host: '127.0.0.1',
          port: '3306',
          userName: 'root',
          password: '',
          dbName: 'gva',
          dbPath: '',
          template: ''
        })
        break
      case 'pgsql':
        Object.assign(form, {
          adminPassword: '123456',
          confirmAdminPassword: '123456',
          redisEnable: form.redisEnable,
          redisHost: form.redisHost,
          redisPort: form.redisPort,
          redisPassword: form.redisPassword,
          redisDB: form.redisDB,
          dbType: 'pgsql',
          host: '127.0.0.1',
          port: '5432',
          userName: 'postgres',
          password: '',
          dbName: 'gva',
          dbPath: '',
          template: 'template0'
        })
        break
      case 'oracle':
        Object.assign(form, {
          adminPassword: '123456',
          confirmAdminPassword: '123456',
          redisEnable: form.redisEnable,
          redisHost: form.redisHost,
          redisPort: form.redisPort,
          redisPassword: form.redisPassword,
          redisDB: form.redisDB,
          dbType: 'oracle',
          host: '127.0.0.1',
          port: '1521',
          userName: 'oracle',
          password: '',
          dbName: 'gva',
          dbPath: '',
          template: ''
        })
        break
      case 'mssql':
        Object.assign(form, {
          adminPassword: '123456',
          confirmAdminPassword: '123456',
          redisEnable: form.redisEnable,
          redisHost: form.redisHost,
          redisPort: form.redisPort,
          redisPassword: form.redisPassword,
          redisDB: form.redisDB,
          dbType: 'mssql',
          host: '127.0.0.1',
          port: '1433',
          userName: 'mssql',
          password: '',
          dbName: 'gva',
          dbPath: '',
          template: ''
        })
        break
      case 'sqlite':
        Object.assign(form, {
          adminPassword: '123456',
          confirmAdminPassword: '123456',
          redisEnable: form.redisEnable,
          redisHost: form.redisHost,
          redisPort: form.redisPort,
          redisPassword: form.redisPassword,
          redisDB: form.redisDB,
          dbType: 'sqlite',
          host: '',
          port: '',
          userName: '',
          password: '',
          dbName: 'gva',
          dbPath: '',
          template: ''
        })
        break
      default:
        Object.assign(form, {
          adminPassword: '123456',
          confirmAdminPassword: '123456',
          redisEnable: form.redisEnable,
          redisHost: form.redisHost,
          redisPort: form.redisPort,
          redisPassword: form.redisPassword,
          redisDB: form.redisDB,
          dbType: 'mysql',
          host: '127.0.0.1',
          port: '3306',
          userName: 'root',
          password: '',
          dbName: 'gva',
          dbPath: '',
          template: ''
        })
    }
  }

  const onSubmit = async () => {
    if (
      !validateDatabaseStep() ||
      !validateRedisStep() ||
      !validateSystemStep()
    ) {
      return
    }

    const loading = ElLoading.service({
      lock: true,
      text: '正在初始化数据库，请稍候',
      spinner: 'loading',
      background: 'rgba(15, 23, 42, 0.46)'
    })

    try {
      const payload = {
        adminPassword: form.adminPassword,
        redis: {
          host: form.redisHost,
          port: form.redisPort,
          password: form.redisPassword,
          db: form.redisDB,
          enable: form.redisEnable
        },
        dbType: form.dbType,
        host: form.host,
        port: form.port,
        userName: form.userName,
        password: form.password,
        dbName: form.dbName,
        dbPath: form.dbPath,
        template: form.template
      }

      const res = await initDB(payload)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: res.msg
        })

        ElMessageBox.confirm(
          '已经完成基础数据库初始化！建议先进行编辑器AI助手配置，以获得更好的开发体验。',
          '配置完成',
          {
            confirmButtonText: '查看AI配置文档',
            cancelButtonText: '稍后配置',
            type: 'success',
            center: true
          }
        )
          .then(() => {
            window.open(
              'https://www.gin-vue-admin.com/guide/server/mcp.html',
              '_blank'
            )
            router.push({ name: 'Login' })
          })
          .catch(() => {
            router.push({ name: 'Login' })
          })
      }
      loading.close()
    } catch (_) {
      loading.close()
    }
  }
</script>

<style scoped>
  .init-page {
    background: radial-gradient(
        circle at 14% 10%,
        rgba(250, 204, 21, 0.06),
        transparent 18%
      ),
      radial-gradient(
        circle at 96% 2%,
        rgba(59, 130, 246, 0.12),
        transparent 16%
      ),
      radial-gradient(
        circle at 5% 100%,
        rgba(56, 189, 248, 0.12),
        transparent 20%
      ),
      linear-gradient(180deg, #fbfdfd 0%, #f7fbfb 48%, #f5fbff 100%);
  }

  .init-page__glow {
    position: absolute;
    border-radius: 9999px;
    filter: blur(120px);
    opacity: 0.72;
  }

  .init-page__glow--primary {
    top: -8rem;
    right: -4rem;
    width: 22rem;
    height: 22rem;
    background: rgba(59, 130, 246, 0.2);
  }

  .init-page__glow--secondary {
    left: -5rem;
    bottom: -4rem;
    width: 18rem;
    height: 18rem;
    background: rgba(125, 211, 252, 0.22);
  }

  .init-page__grid {
    position: absolute;
    inset: 0;
    background-image: linear-gradient(
        rgba(148, 163, 184, 0.045) 1px,
        transparent 1px
      ),
      linear-gradient(90deg, rgba(148, 163, 184, 0.045) 1px, transparent 1px);
    background-size: 54px 54px;
    mask-image: linear-gradient(
      180deg,
      rgba(255, 255, 255, 0.92),
      rgba(255, 255, 255, 0.48)
    );
  }

  .init-shell {
    padding-bottom: 56px;
  }

  .init-brand__logo {
    background: linear-gradient(
      180deg,
      rgba(30, 41, 59, 0.98),
      rgba(30, 64, 175, 0.92)
    );
    color: #f8fafc;
    font-size: 1.125rem;
    font-weight: 700;
    letter-spacing: 0.08em;
    box-shadow: 0 10px 28px rgba(37, 99, 235, 0.18),
      0 0 0 1px rgba(255, 255, 255, 0.52) inset;
  }

  .init-panel {
    position: relative;
    background: rgba(255, 255, 255, 0.92);
    border: 1px solid rgba(255, 255, 255, 0.96);
    box-shadow: 0 24px 48px rgba(148, 163, 184, 0.14),
      0 8px 24px rgba(255, 255, 255, 0.7) inset;
    backdrop-filter: blur(22px);
  }

  .init-panel::before {
    content: '';
    position: absolute;
    inset: 0;
    border-radius: inherit;
    padding: 1px;
    background: linear-gradient(
      135deg,
      rgba(226, 232, 240, 0.9),
      rgba(255, 255, 255, 1)
    );
    -webkit-mask: linear-gradient(#fff 0 0) content-box,
      linear-gradient(#fff 0 0);
    -webkit-mask-composite: xor;
    mask-composite: exclude;
    pointer-events: none;
  }

  .init-step-card {
    border-radius: 22px;
    padding: 18px 18px 20px;
    background: rgba(255, 255, 255, 0.72);
    border: 1px solid rgba(226, 232, 240, 0.88);
    box-shadow: 0 8px 20px rgba(148, 163, 184, 0.06);
    transition: transform 0.2s ease, border-color 0.2s ease,
      box-shadow 0.2s ease, background-color 0.2s ease;
  }

  .init-step-card:hover {
    transform: translateY(-1px);
    border-color: var(--el-color-primary-bg, rgba(59, 130, 246, 0.4));
  }

  .init-step-card.is-active {
    background: linear-gradient(
      180deg,
      var(--el-color-primary-light-10, #f0f8ff),
      rgba(255, 255, 255, 0.92)
    );
    border-color: var(--el-color-primary-bg, rgba(59, 130, 246, 0.4));
    box-shadow: 0 12px 28px rgba(59, 130, 246, 0.12);
  }

  .init-step-card.is-completed {
    border-color: rgba(59, 130, 246, 0.22);
    background: rgba(248, 250, 252, 0.92);
  }

  .init-step-card__index {
    display: inline-flex;
    margin-bottom: 10px;
    font-size: 12px;
    font-weight: 700;
    letter-spacing: 0.24em;
    color: var(--el-color-primary, #3b82f6);
  }

  .init-step-card__title {
    font-size: 18px;
    font-weight: 600;
    color: #0f172a;
  }

  .init-step-card__desc {
    margin-top: 8px;
    font-size: 13px;
    line-height: 1.6;
    color: #64748b;
  }

  .init-section {
    border-radius: 24px;
    padding: 24px;
    background: rgba(255, 255, 255, 0.72);
    border: 1px solid rgba(226, 232, 240, 0.86);
    box-shadow: 0 10px 26px rgba(148, 163, 184, 0.08);
  }

  .init-section__head {
    margin-bottom: 18px;
  }

  .init-section__eyebrow {
    margin-bottom: 8px;
    font-size: 12px;
    font-weight: 700;
    letter-spacing: 0.24em;
    color: var(--el-color-primary, #3b82f6);
  }

  .init-section__title {
    font-size: 24px;
    font-weight: 600;
    line-height: 1.2;
    color: #0f172a;
  }

  .init-section__desc {
    margin-top: 8px;
    font-size: 14px;
    line-height: 1.7;
    color: #64748b;
  }

  .init-tip-card {
    border-radius: 18px;
    padding: 16px 18px;
    background: rgba(248, 250, 252, 0.9);
    border: 1px solid rgba(226, 232, 240, 0.9);
    color: #334155;
    line-height: 1.7;
  }

  .init-switch-wrap {
    display: flex;
    min-height: 48px;
    align-items: center;
    gap: 12px;
    border-radius: 14px;
    background: rgba(255, 255, 255, 0.96);
    box-shadow: 0 0 0 1px rgb(226 232 240 / 95%);
    padding: 0 14px;
  }

  .init-switch-label {
    font-size: 14px;
    color: #475569;
  }

  .init-link {
    margin: 0 4px;
    color: var(--el-color-primary, #3b82f6);
    font-weight: 600;
  }

  .init-doc-btn,
  .init-test-btn {
    border-radius: 18px;
    border-color: var(--el-color-primary-bg, rgba(59, 130, 246, 0.4));
    color: var(--el-color-primary, #3b82f6);
    background: rgba(255, 255, 255, 0.9);
  }

  .init-doc-btn:hover,
  .init-test-btn:hover {
    border-color: var(--el-color-primary, #3b82f6);
    color: var(--el-color-primary-dark-1, #4d8df6);
    background: var(--el-color-primary-light-10, #f0f8ff);
  }

  .init-submit {
    border-radius: 18px;
    background: linear-gradient(
      135deg,
      var(--el-color-primary, #3b82f6) 0%,
      var(--el-color-primary-dark-1, #4d8df6) 100%
    );
    box-shadow: 0 10px 24px rgba(59, 130, 246, 0.24);
  }

  .init-submit:hover {
    background: linear-gradient(
      135deg,
      var(--el-color-primary-dark-1, #4d8df6) 0%,
      var(--el-color-primary-dark-2, #5f99f7) 100%
    );
  }

  .init-nav-btn {
    border-radius: 18px;
    border-color: rgba(203, 213, 225, 0.9);
    color: #475569;
    background: rgba(255, 255, 255, 0.86);
  }

  .init-nav-btn:hover {
    border-color: rgba(148, 163, 184, 0.9);
    color: #0f172a;
    background: #ffffff;
  }

  :deep(.el-form-item) {
    margin-bottom: 0;
  }

  :deep(.el-form-item__label) {
    padding-bottom: 10px;
    font-size: 13px;
    font-weight: 600;
    color: #334155;
  }

  :deep(.init-input .el-input__wrapper),
  :deep(.init-input .el-select__wrapper) {
    min-height: 48px;
    border-radius: 14px;
    background: rgba(255, 255, 255, 0.96);
    box-shadow: 0 0 0 1px rgb(226 232 240 / 95%);
    transition: box-shadow 0.2s ease, background-color 0.2s ease;
  }

  :deep(.init-input .el-input__wrapper:hover),
  :deep(.init-input .el-select__wrapper:hover) {
    background: #ffffff;
    box-shadow: 0 0 0 1px var(--el-color-primary-bg, rgba(59, 130, 246, 0.4));
  }

  :deep(.init-input.is-focus .el-input__wrapper),
  :deep(.init-input .el-input__wrapper.is-focus),
  :deep(.init-input .el-select__wrapper.is-focused) {
    background: #ffffff;
    box-shadow: 0 0 0 1px var(--el-color-primary-light-8),
      0 0 0 4px var(--el-color-primary-light-10);
  }

  :deep(.init-input .el-input__inner::placeholder) {
    color: #cbd5e1;
  }

  :deep(.init-input-number .el-input__wrapper) {
    min-height: 48px;
    border-radius: 14px;
    box-shadow: 0 0 0 1px rgb(226 232 240 / 95%);
  }

  @media (max-width: 1023px) {
    .init-page {
      background: radial-gradient(
          circle at top,
          rgba(59, 130, 246, 0.12),
          transparent 28%
        ),
        linear-gradient(180deg, #fbfdfd 0%, #f6fbfb 100%);
    }

    .init-page__grid {
      opacity: 0.65;
    }

    .init-shell {
      padding-bottom: 32px;
    }
  }
</style>
