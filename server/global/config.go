package global

import (
	"path/filepath"
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/consts"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/env"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/secret"
)

func defaultGeneralDBFromEnv(
	prefixEnvKey, pathEnvKey, portEnvKey, configEnvKey, dbNameEnvKey, usernameEnvKey, passwordEnvKey, engineEnvKey, logModeEnvKey string,
	maxIdleConnsEnvKey, maxOpenConnsEnvKey, singularEnvKey, logZapEnvKey string,
) config.GeneralDB {
	return config.GeneralDB{
		Prefix:       env.GetEnv(prefixEnvKey, consts.DB_PREFIX_DEFAULT),
		Path:         env.GetEnv(pathEnvKey, consts.DB_PATH_DEFAULT),
		Port:         env.GetEnv(portEnvKey, consts.DB_PORT_DEFAULT),
		Config:       env.GetEnv(configEnvKey, consts.DB_CONFIG_DEFAULT),
		Dbname:       env.GetEnv(dbNameEnvKey, consts.DB_NAME_DEFAULT),
		Username:     env.GetEnv(usernameEnvKey, consts.DB_USERNAME_DEFAULT),
		Password:     env.GetEnv(passwordEnvKey, consts.DB_PASSWORD_DEFAULT),
		Engine:       env.GetEnv(engineEnvKey, consts.DB_ENGINE_DEFAULT),
		LogMode:      env.GetEnv(logModeEnvKey, consts.DB_LOG_MODE_DEFAULT),
		MaxIdleConns: env.GetEnvAsInt(maxIdleConnsEnvKey, consts.DB_MAX_IDLE_CONNS_DEFAULT),
		MaxOpenConns: env.GetEnvAsInt(maxOpenConnsEnvKey, consts.DB_MAX_OPEN_CONNS_DEFAULT),
		Singular:     env.GetEnvAsBool(singularEnvKey, consts.DB_SINGULAR_DEFAULT),
		LogZap:       env.GetEnvAsBool(logZapEnvKey, consts.DB_LOG_ZAP_DEFAULT),
	}
}

func mongoHostsFromEnv() []*config.MongoHost {
	hostPairs := env.GetEnvAsSlice(consts.ENV_MONGO_HOSTS, nil, ",")
	if len(hostPairs) == 0 {
		return nil
	}

	hosts := make([]*config.MongoHost, 0, len(hostPairs))
	for _, hostPair := range hostPairs {
		parts := strings.SplitN(hostPair, ":", 2)
		host := strings.TrimSpace(parts[0])
		port := ""
		if len(parts) == 2 {
			port = strings.TrimSpace(parts[1])
		}
		if host == "" {
			continue
		}
		hosts = append(hosts, &config.MongoHost{Host: host, Port: port})
	}
	if len(hosts) == 0 {
		return nil
	}
	return hosts
}

func diskListFromEnv() []config.DiskList {
	mountPoints := env.GetEnvAsSlice(consts.ENV_DISK_MOUNT_POINTS, nil, ",")
	if len(mountPoints) == 0 {
		return make([]config.DiskList, 0)
	}

	diskList := make([]config.DiskList, 0, len(mountPoints))
	for _, mountPoint := range mountPoints {
		if mountPoint == "" {
			continue
		}
		diskList = append(diskList, config.DiskList{
			Disk: config.Disk{MountPoint: mountPoint},
		})
	}
	return diskList
}

func DefaultConfig() config.Server {
	jwtSignKey := env.GetEnv(consts.ENV_JWT_SIGNING_KEY, "")
	if jwtSignKey == "" {
		var err error
		jwtSignKey, err = secret.GenerateSecret(32)
		if err != nil {
			panic("failed to generate JWT signing key: " + err.Error())
		}
	}

	configDefault := config.Server{
		JWT: config.JWT{
			SigningKey:  jwtSignKey,
			ExpiresTime: env.GetEnv(consts.ENV_JWT_EXPIRES_TIME, consts.JWT_EXPIRES_TIME_DEFAULT),
			BufferTime:  env.GetEnv(consts.ENV_JWT_BUFFER_TIME, consts.JWT_BUFFER_TIME_DEFAULT),
			Issuer:      env.GetEnv(consts.ENV_JWT_ISSUER, consts.JWT_ISSUER_DEFAULT),
		},
		Zap: config.Zap{
			Level:         env.GetEnv(consts.ENV_ZAP_LEVEL, consts.ZAP_LEVEL_DEFAULT),
			Prefix:        env.GetEnv(consts.ENV_ZAP_PREFIX, consts.ZAP_PREFIX_DEFAULT),
			Format:        env.GetEnv(consts.ENV_ZAP_FORMAT, consts.ZAP_FORMAT_DEFAULT),
			Director:      env.GetEnv(consts.ENV_ZAP_DIRECTOR, consts.ZAP_DIRECTOR_DEFAULT),
			EncodeLevel:   env.GetEnv(consts.ENV_ZAP_ENCODE_LEVEL, consts.ZAP_ENCODE_LEVEL_DEFAULT),
			StacktraceKey: env.GetEnv(consts.ENV_ZAP_STACKTRACE_KEY, consts.ZAP_STACKTRACE_KEY_DEFAULT),
			ShowLine:      env.GetEnvAsBool(consts.ENV_ZAP_SHOW_LINE, consts.ZAP_SHOW_LINE_DEFAULT),
			LogInConsole:  env.GetEnvAsBool(consts.ENV_ZAP_LOG_IN_CONSOLE, consts.ZAP_LOG_IN_CONSOLE_DEFAULT),
			RetentionDay:  env.GetEnvAsInt(consts.ENV_ZAP_RETENTION_DAY, consts.ZAP_RETENTION_DAY_DEFAULT),
		},
		Redis: config.Redis{
			Name:         env.GetEnv(consts.ENV_REDIS_NAME, consts.REIDS_NAME_DEFAULT),
			Addr:         env.GetEnv(consts.ENV_REDIS_ADDR, consts.REDIS_ADDR_DEFAULT),
			Password:     env.GetEnv(consts.ENV_REDIS_PASSWORD, consts.REDIS_PASSWORD_DEFAULT),
			DB:           env.GetEnvAsInt(consts.ENV_REDIS_DB, consts.REDIS_DB_DEFAULT),
			UseCluster:   env.GetEnvAsBool(consts.ENV_REDIS_USE_CLUSTER, consts.REDIS_USE_CLUSTER_DEFAULT),
			ClusterAddrs: env.GetEnvAsSlice(consts.ENV_REDIS_CLUSTER_ADDRS, []string{}, ","),
		},
		RedisList: make([]config.Redis, 0),
		Mongo: config.Mongo{
			Coll:             env.GetEnv(consts.ENV_MONGO_COLL, consts.MONGO_COLL_DEFAULT),
			Options:          env.GetEnv(consts.ENV_MONGO_OPTIONS, consts.MONGO_OPTIONS_DEFAULT),
			Database:         env.GetEnv(consts.ENV_MONGO_DATABASE, consts.MONGO_DATABASE_DEFAULT),
			Username:         env.GetEnv(consts.ENV_MONGO_USERNAME, consts.MONGO_USERNAME_DEFAULT),
			Password:         env.GetEnv(consts.ENV_MONGO_PASSWORD, consts.MONGO_PASSWORD_DEFAULT),
			AuthSource:       env.GetEnv(consts.ENV_MONGO_AUTH_SOURCE, consts.MONGO_AUTH_SOURCE_DEFAULT),
			MinPoolSize:      env.GetEnvAsUint64(consts.ENV_MONGO_MIN_POOL_SIZE, consts.MONGO_MIN_POOL_SIZE_DEFAULT),
			MaxPoolSize:      env.GetEnvAsUint64(consts.ENV_MONGO_MAX_POOL_SIZE, consts.MONGO_MAX_POOL_SIZE_DEFAULT),
			SocketTimeoutMs:  env.GetEnvAsInt64(consts.ENV_MONGO_SOCKET_TIMEOUT_MS, consts.MONGO_SOCKET_TIMEOUT_MS_DEFAULT),
			ConnectTimeoutMs: env.GetEnvAsInt64(consts.ENV_MONGO_CONNECT_TIMEOUT_MS, consts.MONGO_CONNECT_TIMEOUT_MS_DEFAULT),
			IsZap:            env.GetEnvAsBool(consts.ENV_MONGO_IS_ZAP, consts.MONGO_IS_ZAP_DEFAULT),
			Hosts:            mongoHostsFromEnv(),
		},
		Email: config.Email{
			To:          env.GetEnv(consts.ENV_EMAIL_TO, consts.EMAIL_TO_DEFAULT),
			From:        env.GetEnv(consts.ENV_EMAIL_FROM, consts.EMAIL_FROM_DEFAULT),
			Host:        env.GetEnv(consts.ENV_EMAIL_HOST, consts.EMAIL_HOST_DEFAULT),
			Secret:      env.GetEnv(consts.ENV_EMAIL_SECRET, consts.EMAIL_SECRET_DEFAULT),
			Nickname:    env.GetEnv(consts.ENV_EMAIL_NICKNAME, consts.EMAIL_NICKNAME_DEFAULT),
			Port:        env.GetEnvAsInt(consts.ENV_EMAIL_PORT, consts.EMAIL_PORT_DEFAULT),
			IsSSL:       env.GetEnvAsBool(consts.ENV_EMAIL_IS_SSL, consts.EMAIL_IS_SSL_DEFAULT),
			IsLoginAuth: env.GetEnvAsBool(consts.ENV_EMAIL_IS_LOGIN_AUTH, consts.EMAIL_IS_LOGIN_AUTH_DEFAULT),
		},
		System: config.System{
			DbType:             env.GetEnv(consts.ENV_SYSTEM_DB_TYPE, consts.SYSTEM_DB_TYPE_DEFAULT),
			OssType:            env.GetEnv(consts.ENV_SYSTEM_OSS_TYPE, consts.SYSTEM_OSS_TYPE_DEFAULT),
			RouterPrefix:       env.GetEnv(consts.ENV_SYSTEM_ROUTER_PREFIX, consts.SYSTEM_ROUTER_PREFIX_DEFAULT),
			Addr:               env.GetEnvAsInt(consts.ENV_SYSTEM_ADDR, consts.SYSTEM_ADDR_DEFAULT),
			LimitCountIP:       env.GetEnvAsInt(consts.ENV_SYSTEM_LIMIT_COUNT_IP, consts.SYSTEM_LIMIT_COUNT_IP_DEFAULT),
			LimitTimeIP:        env.GetEnvAsInt(consts.ENV_SYSTEM_LIMIT_TIME_IP, consts.SYSTEM_LIMIT_TIME_IP_DEFAULT),
			UseMultipoint:      env.GetEnvAsBool(consts.ENV_SYSTEM_USE_MULTIPOINT, consts.SYSTEM_USE_MULTIPOINT_DEFAULT),
			UseRedis:           env.GetEnvAsBool(consts.ENV_SYSTEM_USE_REDIS, consts.SYSTEM_USE_REDIS_DEFAULT),
			UseMongo:           env.GetEnvAsBool(consts.ENV_SYSTEM_USE_MONGO, consts.SYSTEM_USE_MONGO_DEFAULT),
			UseStrictAuth:      env.GetEnvAsBool(consts.ENV_SYSTEM_USE_STRICT_AUTH, consts.SYSTEM_USE_STRICT_AUTH_DEFAULT),
			DisableAutoMigrate: env.GetEnvAsBool(consts.ENV_SYSTEM_DISABLE_AUTO_MIGRATE, consts.SYSTEM_DISABLE_AUTO_MIGRATE_DEFAULT),
		},
		Captcha: config.Captcha{
			KeyLong:            env.GetEnvAsInt(consts.ENV_CAPTCHA_KEY_LONG, consts.CAPTCHA_KEY_LONG_DEFAULT),
			ImgWidth:           env.GetEnvAsInt(consts.ENV_CAPTCHA_IMG_WIDTH, consts.CAPTCHA_IMG_WIDTH_DEFAULT),
			ImgHeight:          env.GetEnvAsInt(consts.ENV_CAPTCHA_IMG_HEIGHT, consts.CAPTCHA_IMG_HEIGHT_DEFAULT),
			OpenCaptcha:        env.GetEnvAsInt(consts.ENV_CAPTCHA_OPEN_CAPTCHA, consts.CAPTCHA_OPEN_CAPTCHA_DEFAULT),
			OpenCaptchaTimeOut: env.GetEnvAsInt(consts.ENV_CAPTCHA_OPEN_CAPTCHA_TIMEOUT, consts.CAPTCHA_OPEN_CAPTCHA_TIMEOUT_DEFAULT),
		},
		AutoCode: config.Autocode{
			Web:    env.GetEnv(consts.ENV_AUTOCODE_WEB, consts.AUTOCODE_WEB_DEFAULT),
			Root:   env.GetEnv(consts.ENV_AUTOCODE_ROOT, consts.AUTOCODE_ROOT_DEFAULT),
			Server: env.GetEnv(consts.ENV_AUTOCODE_SERVER, consts.AUTOCODE_SERVER_DEFAULT),
			Module: env.GetEnv(consts.ENV_AUTOCODE_MODULE, consts.AUTOCODE_MODULE_DEFAULT),
			AiPath: env.GetEnv(consts.ENV_AUTOCODE_AI_PATH, consts.AUTOCODE_AI_PATH_DEFAULT),
		},
		Mysql: config.Mysql{GeneralDB: defaultGeneralDBFromEnv(
			consts.ENV_MYSQL_PREFIX, consts.ENV_MYSQL_PATH, consts.ENV_MYSQL_PORT, consts.ENV_MYSQL_CONFIG,
			consts.ENV_MYSQL_DB_NAME, consts.ENV_MYSQL_USERNAME, consts.ENV_MYSQL_PASSWORD, consts.ENV_MYSQL_ENGINE,
			consts.ENV_MYSQL_LOG_MODE, consts.ENV_MYSQL_MAX_IDLE_CONNS, consts.ENV_MYSQL_MAX_OPEN_CONNS,
			consts.ENV_MYSQL_SINGULAR, consts.ENV_MYSQL_LOG_ZAP,
		)},
		Mssql: config.Mssql{GeneralDB: defaultGeneralDBFromEnv(
			consts.ENV_MSSQL_PREFIX, consts.ENV_MSSQL_PATH, consts.ENV_MSSQL_PORT, consts.ENV_MSSQL_CONFIG,
			consts.ENV_MSSQL_DB_NAME, consts.ENV_MSSQL_USERNAME, consts.ENV_MSSQL_PASSWORD, consts.ENV_MSSQL_ENGINE,
			consts.ENV_MSSQL_LOG_MODE, consts.ENV_MSSQL_MAX_IDLE_CONNS, consts.ENV_MSSQL_MAX_OPEN_CONNS,
			consts.ENV_MSSQL_SINGULAR, consts.ENV_MSSQL_LOG_ZAP,
		)},
		Pgsql: config.Pgsql{GeneralDB: defaultGeneralDBFromEnv(
			consts.ENV_PGSQL_PREFIX, consts.ENV_PGSQL_PATH, consts.ENV_PGSQL_PORT, consts.ENV_PGSQL_CONFIG,
			consts.ENV_PGSQL_DB_NAME, consts.ENV_PGSQL_USERNAME, consts.ENV_PGSQL_PASSWORD, consts.ENV_PGSQL_ENGINE,
			consts.ENV_PGSQL_LOG_MODE, consts.ENV_PGSQL_MAX_IDLE_CONNS, consts.ENV_PGSQL_MAX_OPEN_CONNS,
			consts.ENV_PGSQL_SINGULAR, consts.ENV_PGSQL_LOG_ZAP,
		)},
		Oracle: config.Oracle{GeneralDB: defaultGeneralDBFromEnv(
			consts.ENV_ORACLE_PREFIX, consts.ENV_ORACLE_PATH, consts.ENV_ORACLE_PORT, consts.ENV_ORACLE_CONFIG,
			consts.ENV_ORACLE_DB_NAME, consts.ENV_ORACLE_USERNAME, consts.ENV_ORACLE_PASSWORD, consts.ENV_ORACLE_ENGINE,
			consts.ENV_ORACLE_LOG_MODE, consts.ENV_ORACLE_MAX_IDLE_CONNS, consts.ENV_ORACLE_MAX_OPEN_CONNS,
			consts.ENV_ORACLE_SINGULAR, consts.ENV_ORACLE_LOG_ZAP,
		)},
		Sqlite: config.Sqlite{GeneralDB: defaultGeneralDBFromEnv(
			consts.ENV_SQLITE_PREFIX, consts.ENV_SQLITE_PATH, consts.ENV_SQLITE_PORT, consts.ENV_SQLITE_CONFIG,
			consts.ENV_SQLITE_DB_NAME, consts.ENV_SQLITE_USERNAME, consts.ENV_SQLITE_PASSWORD, consts.ENV_SQLITE_ENGINE,
			consts.ENV_SQLITE_LOG_MODE, consts.ENV_SQLITE_MAX_IDLE_CONNS, consts.ENV_SQLITE_MAX_OPEN_CONNS,
			consts.ENV_SQLITE_SINGULAR, consts.ENV_SQLITE_LOG_ZAP,
		)},
		DBList: make([]config.SpecializedDB, 0),
		Local: config.Local{
			Path:      env.GetEnv(consts.ENV_LOCAL_PATH, filepath.Join(consts.GVA_DATA_PATH, consts.LOCAL_PATH_DEFAULT)),
			StorePath: env.GetEnv(consts.ENV_LOCAL_STORE_PATH, filepath.Join(consts.GVA_DATA_PATH, consts.LOCAL_STORE_PATH_DEFAULT)),
		},
		Qiniu: config.Qiniu{
			Zone:          env.GetEnv(consts.ENV_QINIU_ZONE, consts.QINIU_ZONE_DEFAULT),
			Bucket:        env.GetEnv(consts.ENV_QINIU_BUCKET, consts.QINIU_BUCKET_DEFAULT),
			ImgPath:       env.GetEnv(consts.ENV_QINIU_IMG_PATH, consts.QINIU_IMG_PATH_DEFAULT),
			AccessKey:     env.GetEnv(consts.ENV_QINIU_ACCESS_KEY, consts.QINIU_ACCESS_KEY_DEFAULT),
			SecretKey:     env.GetEnv(consts.ENV_QINIU_SECRET_KEY, consts.QINIU_SECRET_KEY_DEFAULT),
			UseHTTPS:      env.GetEnvAsBool(consts.ENV_QINIU_USE_HTTPS, consts.QINIU_USE_HTTPS_DEFAULT),
			UseCdnDomains: env.GetEnvAsBool(consts.ENV_QINIU_USE_CDN_DOMAINS, consts.QINIU_USE_CDN_DOMAINS_DEFAULT),
		},
		AliyunOSS: config.AliyunOSS{
			Endpoint:        env.GetEnv(consts.ENV_ALIYUN_OSS_ENDPOINT, consts.ALIYUN_OSS_ENDPOINT_DEFAULT),
			AccessKeyId:     env.GetEnv(consts.ENV_ALIYUN_OSS_ACCESS_KEY_ID, consts.ALIYUN_OSS_ACCESS_KEY_ID_DEFAULT),
			AccessKeySecret: env.GetEnv(consts.ENV_ALIYUN_OSS_ACCESS_KEY_SECRET, consts.ALIYUN_OSS_ACCESS_KEY_SECRET_DEFAULT),
			BucketName:      env.GetEnv(consts.ENV_ALIYUN_OSS_BUCKET_NAME, consts.ALIYUN_OSS_BUCKET_NAME_DEFAULT),
			BucketUrl:       env.GetEnv(consts.ENV_ALIYUN_OSS_BUCKET_URL, consts.ALIYUN_OSS_BUCKET_URL_DEFAULT),
			BasePath:        env.GetEnv(consts.ENV_ALIYUN_OSS_BASE_PATH, consts.ALIYUN_OSS_BASE_PATH_DEFAULT),
		},
		HuaWeiObs: config.HuaWeiObs{
			Path:      env.GetEnv(consts.ENV_HUAWEI_OBS_PATH, consts.HUAWEI_OBS_PATH_DEFAULT),
			Bucket:    env.GetEnv(consts.ENV_HUAWEI_OBS_BUCKET, consts.HUAWEI_OBS_BUCKET_DEFAULT),
			Endpoint:  env.GetEnv(consts.ENV_HUAWEI_OBS_ENDPOINT, consts.HUAWEI_OBS_ENDPOINT_DEFAULT),
			AccessKey: env.GetEnv(consts.ENV_HUAWEI_OBS_ACCESS_KEY, consts.HUAWEI_OBS_ACCESS_KEY_DEFAULT),
			SecretKey: env.GetEnv(consts.ENV_HUAWEI_OBS_SECRET_KEY, consts.HUAWEI_OBS_SECRET_KEY_DEFAULT),
		},
		TencentCOS: config.TencentCOS{
			Bucket:     env.GetEnv(consts.ENV_TENCENT_COS_BUCKET, consts.TENCENT_COS_BUCKET_DEFAULT),
			Region:     env.GetEnv(consts.ENV_TENCENT_COS_REGION, consts.TENCENT_COS_REGION_DEFAULT),
			SecretID:   env.GetEnv(consts.ENV_TENCENT_COS_SECRET_ID, consts.TENCENT_COS_SECRET_ID_DEFAULT),
			SecretKey:  env.GetEnv(consts.ENV_TENCENT_COS_SECRET_KEY, consts.TENCENT_COS_SECRET_KEY_DEFAULT),
			BaseURL:    env.GetEnv(consts.ENV_TENCENT_COS_BASE_URL, consts.TENCENT_COS_BASE_URL_DEFAULT),
			PathPrefix: env.GetEnv(consts.ENV_TENCENT_COS_PATH_PREFIX, consts.TENCENT_COS_PATH_PREFIX_DEFAULT),
		},
		AwsS3: config.AwsS3{
			Bucket:           env.GetEnv(consts.ENV_AWS_S3_BUCKET, consts.AWS_S3_BUCKET_DEFAULT),
			Region:           env.GetEnv(consts.ENV_AWS_S3_REGION, consts.AWS_S3_REGION_DEFAULT),
			Endpoint:         env.GetEnv(consts.ENV_AWS_S3_ENDPOINT, consts.AWS_S3_ENDPOINT_DEFAULT),
			SecretID:         env.GetEnv(consts.ENV_AWS_S3_SECRET_ID, consts.AWS_S3_SECRET_ID_DEFAULT),
			SecretKey:        env.GetEnv(consts.ENV_AWS_S3_SECRET_KEY, consts.AWS_S3_SECRET_KEY_DEFAULT),
			BaseURL:          env.GetEnv(consts.ENV_AWS_S3_BASE_URL, consts.AWS_S3_BASE_URL_DEFAULT),
			PathPrefix:       env.GetEnv(consts.ENV_AWS_S3_PATH_PREFIX, consts.AWS_S3_PATH_PREFIX_DEFAULT),
			S3ForcePathStyle: env.GetEnvAsBool(consts.ENV_AWS_S3_FORCE_PATH_STYLE, consts.AWS_S3_FORCE_PATH_STYLE_DEFAULT),
			DisableSSL:       env.GetEnvAsBool(consts.ENV_AWS_S3_DISABLE_SSL, consts.AWS_S3_DISABLE_SSL_DEFAULT),
		},
		CloudflareR2: config.CloudflareR2{
			Bucket:          env.GetEnv(consts.ENV_CLOUDFLARE_R2_BUCKET, consts.CLOUDFLARE_R2_BUCKET_DEFAULT),
			BaseURL:         env.GetEnv(consts.ENV_CLOUDFLARE_R2_BASE_URL, consts.CLOUDFLARE_R2_BASE_URL_DEFAULT),
			Path:            env.GetEnv(consts.ENV_CLOUDFLARE_R2_PATH, consts.CLOUDFLARE_R2_PATH_DEFAULT),
			AccountID:       env.GetEnv(consts.ENV_CLOUDFLARE_R2_ACCOUNT_ID, consts.CLOUDFLARE_R2_ACCOUNT_ID_DEFAULT),
			AccessKeyID:     env.GetEnv(consts.ENV_CLOUDFLARE_R2_ACCESS_KEY_ID, consts.CLOUDFLARE_R2_ACCESS_KEY_ID_DEFAULT),
			SecretAccessKey: env.GetEnv(consts.ENV_CLOUDFLARE_R2_SECRET_ACCESS_KEY, consts.CLOUDFLARE_R2_SECRET_ACCESS_KEY_DEFAULT),
		},
		Minio: config.Minio{
			Endpoint:        env.GetEnv(consts.ENV_MINIO_ENDPOINT, consts.MINIO_ENDPOINT_DEFAULT),
			AccessKeyId:     env.GetEnv(consts.ENV_MINIO_ACCESS_KEY_ID, consts.MINIO_ACCESS_KEY_ID_DEFAULT),
			AccessKeySecret: env.GetEnv(consts.ENV_MINIO_ACCESS_KEY_SECRET, consts.MINIO_ACCESS_KEY_SECRET_DEFAULT),
			BucketName:      env.GetEnv(consts.ENV_MINIO_BUCKET_NAME, consts.MINIO_BUCKET_NAME_DEFAULT),
			UseSSL:          env.GetEnvAsBool(consts.ENV_MINIO_USE_SSL, consts.MINIO_USE_SSL_DEFAULT),
			BasePath:        env.GetEnv(consts.ENV_MINIO_BASE_PATH, consts.MINIO_BASE_PATH_DEFAULT),
			BucketUrl:       env.GetEnv(consts.ENV_MINIO_BUCKET_URL, consts.MINIO_BUCKET_URL_DEFAULT),
		},
		Excel: config.Excel{
			Dir: env.GetEnv(consts.ENV_EXCEL_DIR, consts.EXCEL_DIR_DEFAULT),
		},
		DiskList: diskListFromEnv(),
		Cors: config.CORS{
			Mode:      env.GetEnv(consts.ENV_CORS_MODE, consts.CORS_MODE_DEFAULT),
			Whitelist: nil,
		},
		MCP: config.MCP{
			Name:        env.GetEnv(consts.ENV_MCP_NAME, consts.MCP_NAME_DEFAULT),
			Version:     env.GetEnv(consts.ENV_MCP_VERSION, consts.MCP_VERSION_DEFAULT),
			SSEPath:     env.GetEnv(consts.ENV_MCP_SSE_PATH, consts.MCP_SSE_PATH_DEFAULT),
			MessagePath: env.GetEnv(consts.ENV_MCP_MESSAGE_PATH, consts.MCP_MESSAGE_PATH_DEFAULT),
			UrlPrefix:   env.GetEnv(consts.ENV_MCP_URL_PREFIX, consts.MCP_URL_PREFIX_DEFAULT),
			Addr:        env.GetEnvAsInt(consts.ENV_MCP_ADDR, consts.MCP_ADDR_DEFAULT),
			Separate:    env.GetEnvAsBool(consts.ENV_MCP_SEPARATE, consts.MCP_SEPARATE_DEFAULT),
		},
	}

	return configDefault
}
