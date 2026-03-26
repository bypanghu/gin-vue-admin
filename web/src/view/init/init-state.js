const DB_OPTION_CONFIGS = {
  mysql: {
    label: 'mysql',
    defaults: {
      dbType: 'mysql',
      host: '127.0.0.1',
      port: '3306',
      userName: 'root',
      password: '',
      dbName: 'gva',
      dbPath: '',
      template: ''
    }
  },
  pgsql: {
    label: 'pgsql',
    defaults: {
      dbType: 'pgsql',
      host: '127.0.0.1',
      port: '5432',
      userName: 'postgres',
      password: '',
      dbName: 'gva',
      dbPath: '',
      template: 'template0'
    }
  },
  mssql: {
    label: 'mssql',
    defaults: {
      dbType: 'mssql',
      host: '127.0.0.1',
      port: '1433',
      userName: 'sa',
      password: '',
      dbName: 'gva',
      dbPath: '',
      template: ''
    }
  },
  sqlite: {
    label: 'sqlite',
    defaults: {
      dbType: 'sqlite',
      host: '',
      port: '',
      userName: '',
      password: '',
      dbName: 'gva',
      dbPath: '',
      template: ''
    }
  }
}

export const INIT_DB_OPTIONS = Object.values(DB_OPTION_CONFIGS).map(
  ({ label, defaults }) => ({
    label,
    value: defaults.dbType
  })
)

export const getDefaultInitForm = () => ({
  adminPassword: '',
  confirmAdminPassword: '',
  redisEnable: false,
  redisHost: '127.0.0.1',
  redisPort: '6379',
  redisPassword: '',
  redisDB: 0,
  ...DB_OPTION_CONFIGS.mysql.defaults
})

export const getFormWithDBDefaults = (form, dbType) => {
  const next = DB_OPTION_CONFIGS[dbType] || DB_OPTION_CONFIGS.mysql
  return {
    ...form,
    ...next.defaults
  }
}

export const buildTestDBPayload = (form) => ({
  dbType: form.dbType,
  host: form.host,
  port: form.port,
  userName: form.userName,
  password: form.password,
  dbName: form.dbName,
  dbPath: form.dbPath,
  template: form.template
})

export const buildRedisPayload = (form) => ({
  host: form.redisHost,
  port: form.redisPort,
  password: form.redisPassword,
  db: form.redisDB,
  enable: form.redisEnable
})

export const buildInitPayload = (form) => ({
  adminPassword: form.adminPassword,
  redis: buildRedisPayload(form),
  ...buildTestDBPayload(form)
})

export const getDatabaseTestSnapshot = (form) =>
  JSON.stringify(buildTestDBPayload(form))

export const getRedisTestSnapshot = (form) =>
  JSON.stringify(buildRedisPayload(form))
