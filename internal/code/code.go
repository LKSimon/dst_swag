package code

import (
	_ "embed"
)

//go:embed code.go
var ByteCodeFile []byte

// Failure 错误时返回结构
type Failure struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 描述信息
}

const (
	Success = 0

	ServerError        = 10101
	TooManyRequests    = 10102
	ParamBindError     = 10103
	AuthorizationError = 10104
	UrlSignError       = 10105
	CacheSetError      = 10106
	CacheGetError      = 10107
	CacheDelError      = 10108
	CacheNotExist      = 10109
	ResubmitError      = 10110
	HashIdsEncodeError = 10111
	HashIdsDecodeError = 10112
	RBACError          = 10113
	RedisConnectError  = 10114
	PgSQLConnectError  = 10115
	WriteConfigError   = 10116
	SendEmailError     = 10117
	PgSQLExecError     = 10118
	GoVersionError     = 10119
	SocketConnectError = 10120
	SocketSendError    = 10121

	VideoTemplateGetError = 20101
	VideoCreateSyncError  = 20102
	VideoCreateAsyncError = 20103
	VideoTaskListError    = 20104

	MaterialCreateError         = 20201
	MaterialCreateDataTypeError = 20202
	MaterialDeleteError         = 20203
	MaterialUpdateNameError     = 20204
	MaterialListError           = 20205
)

//
//func Text(code int) string {
//	lang := configs.Get().Language.Local
//
//	if lang == configs.ZhCN {
//		return zhCNText[code]
//	}
//
//	if lang == configs.EnUS {
//		return enUSText[code]
//	}
//
//	return zhCNText[code]
//}
