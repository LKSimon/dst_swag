package code

var zhCNText = map[int]string{
	Success: "成功",

	ServerError:        "内部服务器错误",
	TooManyRequests:    "请求过多",
	ParamBindError:     "参数信息错误",
	AuthorizationError: "签名信息错误",
	UrlSignError:       "参数签名错误",
	CacheSetError:      "设置缓存失败",
	CacheGetError:      "获取缓存失败",
	CacheDelError:      "删除缓存失败",
	CacheNotExist:      "缓存不存在",
	ResubmitError:      "请勿重复提交",
	HashIdsEncodeError: "HashID 加密失败",
	HashIdsDecodeError: "HashID 解密失败",
	RBACError:          "暂无访问权限",
	RedisConnectError:  "Redis 连接失败",
	PgSQLConnectError:  "PGSQL 连接失败",
	WriteConfigError:   "写入配置文件失败",
	SendEmailError:     "发送邮件失败",
	PgSQLExecError:     "SQL 执行失败",
	GoVersionError:     "Go 版本不满足要求",
	SocketConnectError: "Socket 未连接",
	SocketSendError:    "Socket 消息发送失败",

	VideoTemplateGetError: "获取模板列表失败",
	VideoCreateSyncError:  "同步合成失败",
	VideoCreateAsyncError: "异步合成失败",
	VideoTaskListError:    "获取合成状态失败",

	MaterialCreateError:     "创建素材失败",
	MaterialDeleteError:     "删除素材失败",
	MaterialUpdateNameError: "更新素材名字失败",
	MaterialListError:       "获取素材列表失败",
}
