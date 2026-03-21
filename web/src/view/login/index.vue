<template>
  <div
    id="userLayout"
    class="login-page relative min-h-screen w-full overflow-hidden"
  >
    <div class="pointer-events-none absolute inset-0">
      <div class="login-page__glow login-page__glow--primary" />
      <div class="login-page__glow login-page__glow--secondary" />
      <div class="login-page__grid" />
    </div>
    <div class="relative z-10 min-h-screen px-5 py-6 md:px-8 md:py-8 xl:px-12">
      <div
        class="mx-auto flex min-h-[calc(100vh-48px)] max-w-7xl items-center justify-center"
      >
        <section
          class="login-shell flex w-full flex-col items-center justify-center lg:max-w-[430px] xl:max-w-[460px]"
        >
          <div class="mb-8 flex flex-col items-center text-center">
            <div
              class="flex h-[72px] w-[72px] items-center justify-center rounded-[24px] text-slate-900 border border-[var(--el-color-primary-light-9)]"
            >
              <Logo :size="7" />
            </div>
            <h1
              class="mt-5 text-[40px] font-semibold tracking-tight text-[var(--el-color-primary)]"
            >
              {{ $GIN_VUE_ADMIN.appName }}
            </h1>
            <p class="mt-2 text-sm text-slate-400">后台管理平台</p>
          </div>

          <div class="login-panel w-full rounded-[28px] p-7 sm:p-8">
            <div class="mb-8 text-center">
              <h2
                class="text-[34px] font-semibold tracking-tight text-slate-900"
              >
                欢迎回来
              </h2>
              <p class="mt-2 text-sm text-slate-400">登录您的账户以继续</p>
            </div>

            <el-form
              ref="loginForm"
              :model="loginFormData"
              class="!flex flex-col gap-4"
              :rules="rules"
              :validate-on-rule-change="false"
              @keyup.enter="submitForm"
            >
              <el-form-item prop="username">
                <div class="w-full">
                  <p class="login-field-label">账号</p>
                  <el-input
                    v-model="loginFormData.username"
                    size="large"
                    placeholder="请输入用户名"
                    autofocus
                    class="login-input"
                  >
                    <template #prefix>
                      <el-icon class="text-slate-300"><User /></el-icon>
                    </template>
                  </el-input>
                </div>
              </el-form-item>
              <el-form-item prop="password">
                <div class="w-full">
                  <p class="login-field-label">密码</p>
                  <el-input
                    v-model="loginFormData.password"
                    show-password
                    size="large"
                    type="password"
                    placeholder="请输入密码"
                    class="login-input"
                  >
                    <template #prefix>
                      <el-icon class="text-slate-300"><Lock /></el-icon>
                    </template>
                  </el-input>
                </div>
              </el-form-item>
              <el-form-item v-if="loginFormData.openCaptcha" prop="captcha">
                <div class="w-full">
                  <p class="login-field-label">验证码</p>
                  <div class="flex w-full gap-3">
                    <el-input
                      v-model="loginFormData.captcha"
                      placeholder="请输入验证码"
                      size="large"
                      class="login-input flex-1"
                    >
                      <template #prefix>
                        <el-icon class="text-slate-300"><Key /></el-icon>
                      </template>
                    </el-input>
                    <div
                      class="login-captcha h-12 w-[132px] overflow-hidden rounded-2xl bg-slate-100"
                    >
                      <img
                        v-if="picPath"
                        class="h-full w-full cursor-pointer object-cover"
                        :src="picPath"
                        alt="请输入验证码"
                        @click="loginVerify()"
                      />
                    </div>
                  </div>
                </div>
              </el-form-item>
              <el-form-item class="mb-4 pt-1">
                <el-button
                  class="h-12 w-full border-0 text-base font-semibold !rounded-full"
                  type="primary"
                  size="large"
                  @click="submitForm"
                >
                  登 录
                </el-button>
              </el-form-item>
              <el-form-item v-if="isDev" class="mb-0">
                <el-button
                  class="h-12 w-full rounded-[18px] border-slate-200 bg-white text-slate-700 shadow-sm"
                  plain
                  size="large"
                  @click="checkInit"
                >
                  前往初始化
                </el-button>
              </el-form-item>
            </el-form>
          </div>

          <div class="login-note mt-5 text-center text-sm text-slate-400">
            使用平台账号继续访问系统
          </div>
        </section>
      </div>
    </div>

    <BottomInfo class="left-0 right-0 absolute bottom-10 mx-auto w-full z-20">
      <div class="links items-center justify-center gap-4 hidden md:flex">
        <a href="https://www.gin-vue-admin.com/" target="_blank">
          <img src="@/assets/docs.png" class="w-6 h-6" alt="文档" />
        </a>
        <a href="https://support.qq.com/product/371961" target="_blank">
          <img src="@/assets/kefu.png" class="w-6 h-6" alt="客服" />
        </a>
        <a
          href="https://github.com/flipped-aurora/gin-vue-admin"
          target="_blank"
        >
          <img src="@/assets/github.png" class="w-6 h-6" alt="github" />
        </a>
        <a href="https://space.bilibili.com/322210472" target="_blank">
          <img src="@/assets/video.png" class="w-6 h-6" alt="视频站" />
        </a>
      </div>
    </BottomInfo>
  </div>
</template>

<script setup>
  import { captcha } from '@/api/user'
  import { checkDB } from '@/api/initdb'
  import BottomInfo from '@/components/bottomInfo/bottomInfo.vue'
  import { reactive, ref } from 'vue'
  import { ElMessage, ElMessageBox } from 'element-plus'
  import { Key, Lock, User } from '@element-plus/icons-vue'
  import { useRouter } from 'vue-router'
  import { useUserStore } from '@/pinia/modules/user'
  import Logo from '@/components/logo/index.vue'
  import { isDev } from '@/utils/env.js'

  defineOptions({
    name: 'Login'
  })

  const router = useRouter()
  const captchaRequiredLength = ref(6)

  const initial = reactive({
    isInitial: false,
    autoHref: false
  })
  // 验证函数
  const checkUsername = (rule, value, callback) => {
    if (value.length < 5) {
      return callback(new Error('请输入正确的用户名'))
    } else {
      callback()
    }
  }
  const checkPassword = (rule, value, callback) => {
    if (value.length < 6) {
      return callback(new Error('请输入正确的密码'))
    } else {
      callback()
    }
  }
  const checkCaptcha = (rule, value, callback) => {
    if (!loginFormData.openCaptcha) {
      return callback()
    }
    const sanitizedValue = (value || '').replace(/\s+/g, '')
    if (!sanitizedValue) {
      return callback(new Error('请输入验证码'))
    }
    if (!/^\d+$/.test(sanitizedValue)) {
      return callback(new Error('验证码须为数字'))
    }
    if (sanitizedValue.length < captchaRequiredLength.value) {
      return callback(
        new Error(`请输入至少${captchaRequiredLength.value}位数字验证码`)
      )
    }
    if (sanitizedValue !== value) {
      loginFormData.captcha = sanitizedValue
    }
    callback()
  }

  // 获取验证码
  const loginVerify = async () => {
    const ele = await captcha()
    captchaRequiredLength.value = Number(ele.data?.captchaLength) || 0
    picPath.value = ele.data?.picPath
    loginFormData.captchaId = ele.data?.captchaId
    loginFormData.openCaptcha = ele.data?.openCaptcha
    initial.autoHref = ele.data?.autoHrefToInitial || false
    initial.isInitial = ele.data?.hasInitial || false

    if (initial.autoHref && !initial.isInitial) {
      ElMessageBox.confirm(
        '系统暂未初始化，是否立即进入初始化页面？',
        '初始化提示',
        {
          confirmButtonText: '进入初始化',
          cancelButtonText: '暂不进入',
          type: 'warning'
        }
      )
        .then(async () => {
          await checkInit()
        })
        .catch(() => {})
    }
  }
  loginVerify()

  // 登录相关操作
  const loginForm = ref(null)
  const picPath = ref('')
  const loginFormData = reactive({
    username: 'admin',
    password: '',
    captcha: '',
    captchaId: '',
    openCaptcha: false
  })
  const rules = reactive({
    username: [{ validator: checkUsername, trigger: 'blur' }],
    password: [{ validator: checkPassword, trigger: 'blur' }],
    captcha: [{ validator: checkCaptcha, trigger: 'blur' }]
  })

  const userStore = useUserStore()
  const login = async () => {
    return await userStore.LoginIn(loginFormData)
  }
  const submitForm = () => {
    loginForm.value.validate(async (v) => {
      if (!v) {
        // 未通过前端静态验证
        ElMessage({
          type: 'error',
          message: '请正确填写登录信息',
          showClose: true
        })
        return false
      }

      // 通过验证，请求登陆
      const flag = await login()

      // 登陆失败，刷新验证码
      if (!flag) {
        await loginVerify()
        return false
      }

      // 登陆成功
      return true
    })
  }

  // 跳转初始化
  const checkInit = async () => {
    const res = await checkDB()
    if (res.code === 0) {
      if (res.data?.needInit) {
        userStore.NeedInit()
        await router.push({ name: 'Init' })
      } else {
        ElMessage({
          type: 'info',
          message: '已配置数据库信息，无法初始化'
        })
      }
    }
  }
</script>

<style scoped>
  .login-page {
    background: radial-gradient(
        circle at 14% 10%,
        rgba(250, 204, 21, 0.06),
        transparent 18%
      ),
      radial-gradient(
        circle at 96% 2%,
        rgba(45, 212, 191, 0.16),
        transparent 16%
      ),
      radial-gradient(
        circle at 5% 100%,
        rgba(56, 189, 248, 0.12),
        transparent 20%
      ),
      linear-gradient(180deg, #fbfdfd 0%, #f7fbfb 48%, #f5fbff 100%);
  }

  .login-page__glow {
    position: absolute;
    border-radius: 9999px;
    filter: blur(120px);
    opacity: 0.72;
  }

  .login-page__glow--primary {
    top: -8rem;
    right: -4rem;
    width: 22rem;
    height: 22rem;
    background: rgba(94, 234, 212, 0.28);
  }

  .login-page__glow--secondary {
    left: -5rem;
    bottom: -4rem;
    width: 18rem;
    height: 18rem;
    background: rgba(125, 211, 252, 0.22);
  }

  .login-page__grid {
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

  .login-shell {
    padding-bottom: 72px;
  }

  .login-panel {
    position: relative;
    background: rgba(255, 255, 255, 0.92);
    border: 1px solid rgba(255, 255, 255, 0.96);
    box-shadow: 0 24px 48px rgba(148, 163, 184, 0.14),
      0 8px 24px rgba(255, 255, 255, 0.7) inset;
    backdrop-filter: blur(22px);
  }

  .login-panel::before {
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

  .login-field-label {
    margin-bottom: 2px;
    font-size: 13px;
    font-weight: 600;
    color: #334155;
  }

  .login-captcha {
    border: 1px solid rgba(226, 232, 240, 0.95);
    box-shadow: 0 8px 18px rgba(148, 163, 184, 0.08);
  }

  .login-note {
    text-shadow: 0 1px 0 rgba(255, 255, 255, 0.7);
  }

  :deep(.login-input .el-input__wrapper) {
    min-height: 48px;
    border-radius: 14px;
    background: rgba(255, 255, 255, 0.96);
    box-shadow: 0 0 0 1px rgb(226 232 240 / 95%);
    transition: box-shadow 0.2s ease, background-color 0.2s ease;
  }

  :deep(.login-input.is-focus .el-input__wrapper),
  :deep(.login-input .el-input__wrapper.is-focus) {
    background: #ffffff;
    box-shadow: 0 0 0 1px var(--el-color-primary-light-8),
      0 0 0 4px var(--el-color-primary-light-10);
  }

  :deep(.login-input .el-input__inner::placeholder) {
    color: #cbd5e1;
  }

  :deep(.el-form-item) {
    margin-bottom: 0;
  }

  @media (max-width: 1023px) {
    .login-page {
      background: radial-gradient(
          circle at top,
          rgba(94, 234, 212, 0.16),
          transparent 28%
        ),
        linear-gradient(180deg, #fbfdfd 0%, #f6fbfb 100%);
    }

    .login-page__grid {
      opacity: 0.65;
    }

    .login-shell {
      padding-bottom: 88px;
    }

    .login-panel {
      background: rgba(255, 255, 255, 0.96);
    }
  }
</style>
