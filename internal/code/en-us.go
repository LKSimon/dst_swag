package code

var enUSText = map[int]string{
	Success: "Success",

	ServerError:        "Internal server error",
	TooManyRequests:    "Too many requests",
	ParamBindError:     "Parameter error",
	AuthorizationError: "Authorization error",
	UrlSignError:       "URL signature error",
	CacheSetError:      "Failed to set cache",
	CacheGetError:      "Failed to get cache",
	CacheDelError:      "Failed to del cache",
	CacheNotExist:      "Cache does not exist",
	ResubmitError:      "Please do not submit repeatedly",
	HashIdsEncodeError: "HashID encryption failed",
	HashIdsDecodeError: "HashID decryption failed",
	RBACError:          "No access",
	RedisConnectError:  "Failed to connection Redis",
	PgSQLConnectError:  "Failed to connection PGSQL",
	WriteConfigError:   "Failed to write configuration file",
	SendEmailError:     "Failed to send mail",
	PgSQLExecError:     "SQL execution failed",
	GoVersionError:     "Go Version mismatch",
	SocketConnectError: "Socket not connected",
	SocketSendError:    "Socket message sending failed",

	VideoTemplateGetError: "Get template list failed",
	VideoCreateSyncError:  "Failed to create video sync",
	VideoCreateAsyncError: "Failed to create video async",
	VideoTaskListError:    "Failed to get status",

	MaterialCreateError:     "Failed to create material",
	MaterialDeleteError:     "Failed to delete material",
	MaterialUpdateNameError: "Failed to update material name",
	MaterialListError:       "Failed to get material list",
}
