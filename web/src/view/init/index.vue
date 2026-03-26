<template>
  <div
    class="relative min-h-screen overflow-hidden bg-[radial-gradient(circle_at_14%_10%,rgba(250,204,21,0.06),transparent_18%),radial-gradient(circle_at_96%_2%,rgba(59,130,246,0.12),transparent_16%),radial-gradient(circle_at_5%_100%,rgba(56,189,248,0.12),transparent_20%),linear-gradient(180deg,#fbfdfd_0%,#f7fbfb_48%,#f5fbff_100%)] lg:bg-[radial-gradient(circle_at_14%_10%,rgba(250,204,21,0.06),transparent_18%),radial-gradient(circle_at_96%_2%,rgba(59,130,246,0.12),transparent_16%),radial-gradient(circle_at_5%_100%,rgba(56,189,248,0.12),transparent_20%),linear-gradient(180deg,#fbfdfd_0%,#f7fbfb_48%,#f5fbff_100%)]"
  >
    <div class="pointer-events-none absolute inset-0">
      <div
        class="absolute -right-16 -top-32 h-88 w-88 rounded-full bg-sky-500/20 opacity-70 blur-[120px]"
      />
      <div
        class="absolute -bottom-16 -left-20 h-72 w-72 rounded-full bg-sky-300/25 opacity-70 blur-[120px]"
      />
      <div
        class="absolute inset-0 bg-[linear-gradient(rgba(148,163,184,0.045)_1px,transparent_1px),linear-gradient(90deg,rgba(148,163,184,0.045)_1px,transparent_1px)] bg-[length:54px_54px] [mask-image:linear-gradient(180deg,rgba(255,255,255,0.92),rgba(255,255,255,0.48))]"
      />
    </div>

    <div class="relative z-10 min-h-screen px-5 py-6 md:px-8 md:py-8 xl:px-12">
      <div
        class="mx-auto flex min-h-[calc(100vh-48px)] max-w-5xl items-center justify-center"
      >
        <section class="w-full pb-8 lg:pb-14">
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

          <div
            class="relative overflow-hidden rounded-[28px] border border-white/95 bg-white/92 p-6 shadow-[0_24px_48px_rgba(148,163,184,0.14),inset_0_8px_24px_rgba(255,255,255,0.7)] backdrop-blur-[22px] sm:p-8 lg:p-10"
          >
            <div class="mb-8 grid gap-3 md:grid-cols-2 xl:grid-cols-4">
              <button
                v-for="(step, index) in steps"
                :key="step.title"
                type="button"
                class="rounded-[22px] border border-slate-200/90 bg-white/75 p-[18px] pb-5 text-left shadow-[0_8px_20px_rgba(148,163,184,0.06)] transition-all duration-200 hover:-translate-y-0.5 hover:border-sky-200"
                :class="{
                  'border-sky-200 bg-[linear-gradient(180deg,var(--el-color-primary-light-10,#f0f8ff),rgba(255,255,255,0.92))] shadow-[0_12px_28px_rgba(59,130,246,0.12)]':
                    currentStep === index,
                  'border-sky-500/20 bg-slate-50/90': currentStep > index
                }"
                @click="goStep(index)"
              >
                <span
                  class="mb-2.5 inline-flex text-xs font-bold tracking-[0.24em] text-sky-500"
                >
                  0{{ index + 1 }}
                </span>
                <h2 class="text-lg font-semibold text-slate-900">
                  {{ step.title }}
                </h2>
                <p class="mt-2 text-[13px] leading-6 text-slate-500">
                  {{ step.desc }}
                </p>
              </button>
            </div>

            <el-form :model="form" label-position="top" size="large">
              <section
                v-show="currentStep === 0"
                class="rounded-3xl border border-slate-200/90 bg-white/75 p-6 shadow-[0_10px_26px_rgba(148,163,184,0.08)]"
              >
                <div class="mb-[18px]">
                  <p class="mb-2 text-xs font-bold tracking-[0.24em] text-sky-500">
                    STEP 1
                  </p>
                  <h2 class="text-2xl font-semibold leading-tight text-slate-900">
                    初始化提示
                  </h2>
                  <p class="mt-2 text-sm leading-7 text-slate-500">
                    开始前先确认环境、文档和数据库引擎配置，避免初始化过程中出现可预期的问题。
                  </p>
                </div>
                <div class="grid gap-3">
                  <div
                    class="rounded-[18px] border border-slate-200/90 bg-slate-50/90 px-[18px] py-4 leading-7 text-slate-700"
                  >
                    请确认你已经具备基础的 Vue 和 Golang 使用经验。
                  </div>
                  <div
                    class="rounded-[18px] border border-slate-200/90 bg-slate-50/90 px-[18px] py-4 leading-7 text-slate-700"
                  >
                    初始化前建议先阅读
                    <a
                      class="mx-1 font-semibold text-sky-500"
                      href="https://www.gin-vue-admin.com"
                      target="_blank"
                    >
                      官方文档
                    </a>
                    与
                    <a
                      class="mx-1 font-semibold text-sky-500"
                      href="https://www.bilibili.com/video/BV1kv4y1g7nT?p=2"
                      target="_blank"
                    >
                      初始化视频
                    </a>
                    。
                  </div>
                  <div
                    class="rounded-[18px] border border-slate-200/90 bg-slate-50/90 px-[18px] py-4 leading-7 text-slate-700"
                  >
                    如果你使用 MySQL，请确认数据库引擎为
                    <span class="font-semibold text-rose-500">InnoDB</span>。
                  </div>
                </div>
                <div class="mt-4 flex justify-end">
                  <el-button
                    class="!rounded-[18px] !border-sky-200 !bg-white/90 !text-sky-500"
                    plain
                    size="large"
                    @click="goDoc"
                  >
                    阅读环境文档
                  </el-button>
                </div>
              </section>

              <section
                v-show="currentStep === 1"
                class="rounded-3xl border border-slate-200/90 bg-white/75 p-6 shadow-[0_10px_26px_rgba(148,163,184,0.08)]"
              >
                <div class="mb-[18px]">
                  <p class="mb-2 text-xs font-bold tracking-[0.24em] text-sky-500">
                    STEP 2
                  </p>
                  <h2 class="text-2xl font-semibold leading-tight text-slate-900">
                    数据库设置
                  </h2>
                  <p class="mt-2 text-sm leading-7 text-slate-500">
                    先选择数据库类型，再填写数据库连接参数。测试连接会调用后端接口校验当前配置是否可用。
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
                      <el-option
                        v-for="option in dbOptions"
                        :key="option.value"
                        :label="option.label"
                        :value="option.value"
                      />
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
                    class="h-12 min-w-[140px] !rounded-[18px] !border-sky-200 !bg-white/90 !text-sky-500"
                    plain
                    size="large"
                    @click="testConnection"
                  >
                    测试连接
                  </el-button>
                </div>
                <p class="mt-3 text-right text-sm text-slate-400">
                  {{
                    initialState.dbCheck
                      ? '当前数据库配置已验证'
                      : '修改数据库信息后需要重新测试连接'
                  }}
                </p>
              </section>

              <section
                v-show="currentStep === 2"
                class="rounded-3xl border border-slate-200/90 bg-white/75 p-6 shadow-[0_10px_26px_rgba(148,163,184,0.08)]"
              >
                <div class="mb-[18px]">
                  <p class="mb-2 text-xs font-bold tracking-[0.24em] text-sky-500">
                    STEP 3
                  </p>
                  <h2 class="text-2xl font-semibold leading-tight text-slate-900">
                    Redis 配置
                  </h2>
                  <p class="mt-2 text-sm leading-7 text-slate-500">
                    Redis 作为可选能力单独配置。启用后可测试连接，并在初始化成功时一并写入服务端配置。
                  </p>
                </div>

                <div class="grid gap-5 md:grid-cols-2">
                  <el-form-item label="启用 Redis" class="mb-0">
                    <div
                      class="flex min-h-12 items-center gap-3 rounded-[14px] bg-white/95 px-[14px] shadow-[0_0_0_1px_rgb(226_232_240_/_0.95)]"
                    >
                      <el-switch v-model="form.redisEnable" />
                      <span class="text-sm text-slate-600">
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
                    class="h-12 min-w-[140px] !rounded-[18px] !border-sky-200 !bg-white/90 !text-sky-500"
                    plain
                    size="large"
                    @click="testRedisConnection"
                  >
                    测试 Redis
                  </el-button>
                </div>
                <p class="mt-3 text-right text-sm text-slate-400">
                  {{
                    !form.redisEnable
                      ? '未启用 Redis 时不会写入启用状态'
                      : initialState.redisCheck
                        ? '当前 Redis 配置已验证'
                        : '修改 Redis 信息后需要重新测试连接'
                  }}
                </p>
              </section>

              <section
                v-show="currentStep === 3"
                class="rounded-3xl border border-slate-200/90 bg-white/75 p-6 shadow-[0_10px_26px_rgba(148,163,184,0.08)]"
              >
                <div class="mb-[18px]">
                  <p class="mb-2 text-xs font-bold tracking-[0.24em] text-sky-500">
                    STEP 4
                  </p>
                  <h2 class="text-2xl font-semibold leading-tight text-slate-900">
                    系统设置
                  </h2>
                  <p class="mt-2 text-sm leading-7 text-slate-500">
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
                  class="!rounded-[18px] !border-slate-300/90 !bg-white/85 !px-12 !text-slate-600"
                  plain
                  size="large"
                  :disabled="currentStep === 0"
                  @click="prevStep"
                >
                  上一步
                </el-button>
                <el-button
                  v-if="currentStep < steps.length - 1"
                  class="!rounded-[18px] !px-12"
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
                  class="h-12 min-w-[160px] !rounded-[18px] !border-0 !bg-[linear-gradient(135deg,var(--el-color-primary,#3b82f6)_0%,var(--el-color-primary-dark-1,#4d8df6)_100%)] text-base font-semibold !shadow-[0_10px_24px_rgba(59,130,246,0.24)]"
                  type="primary"
                  size="large"
                  :disabled="
                    !initialState.dbCheck ||
                    (form.redisEnable && !initialState.redisCheck)
                  "
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
  import { reactive, ref, watch } from 'vue'
  import { ElLoading, ElMessage, ElMessageBox } from 'element-plus'
  import { useRouter } from 'vue-router'
  import Logo from '@/components/logo/index.vue'
  import {
    INIT_DB_OPTIONS,
    buildInitPayload,
    buildRedisPayload,
    buildTestDBPayload,
    getDatabaseTestSnapshot,
    getDefaultInitForm,
    getFormWithDBDefaults,
    getRedisTestSnapshot
  } from './init-state'

  defineOptions({
    name: 'Init'
  })

  const router = useRouter()
  const currentStep = ref(0)
  const dbOptions = INIT_DB_OPTIONS
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

  const form = reactive(getDefaultInitForm())

  const initialState = reactive({
    dbCheck: false,
    redisCheck: false
  })
  const dbTestSnapshot = ref('')
  const redisTestSnapshot = ref('')

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
      const payload = buildTestDBPayload(form)
      const res = await testDB(payload)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: res.msg || '数据库连接成功'
        })
        initialState.dbCheck = true
        dbTestSnapshot.value = getDatabaseTestSnapshot(form)
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
      const res = await testRedis(buildRedisPayload(form))
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: res.msg || 'Redis 连接成功'
        })
        initialState.redisCheck = true
        redisTestSnapshot.value = getRedisTestSnapshot(form)
      }
    } finally {
      loading.close()
    }
  }

  const changeDB = (val) => {
    Object.assign(form, getFormWithDBDefaults(form, val))
  }

  const onSubmit = async () => {
    if (
      !validateDatabaseStep() ||
      !validateRedisStep() ||
      !validateSystemStep()
    ) {
      return
    }

    if (!initialState.dbCheck) {
      ElMessage({
        type: 'error',
        message: '请先测试并确认数据库连接'
      })
      return
    }

    if (form.redisEnable && !initialState.redisCheck) {
      ElMessage({
        type: 'error',
        message: '请先测试并确认 Redis 连接'
      })
      return
    }

    const loading = ElLoading.service({
      lock: true,
      text: '正在初始化数据库，请稍候',
      spinner: 'loading',
      background: 'rgba(15, 23, 42, 0.46)'
    })

    try {
      const res = await initDB(buildInitPayload(form))
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

  watch(
    () => getDatabaseTestSnapshot(form),
    (snapshot) => {
      if (!dbTestSnapshot.value) {
        return
      }
      initialState.dbCheck = dbTestSnapshot.value === snapshot
    }
  )

  watch(
    () => getRedisTestSnapshot(form),
    (snapshot) => {
      if (!form.redisEnable) {
        initialState.redisCheck = false
        redisTestSnapshot.value = ''
        return
      }
      if (!redisTestSnapshot.value) {
        return
      }
      initialState.redisCheck = redisTestSnapshot.value === snapshot
    }
  )
</script>

<style scoped>
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
</style>
