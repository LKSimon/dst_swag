package img

import (
	"dst_swag/internal/pkg/core"
	"dst_swag/internal/service/img"
	"go.uber.org/zap"
)

type Handler interface {
	i()

	// TemplateList 视频模板列表接口
	// @Tags API.img
	// @Router /api/img/template [get]
	TemplateList() core.HandlerFunc

	// CreateSync 合成素材(同步)
	// @Tags API.img
	// @Router /api/img/createSync [post]
	CreateSync() core.HandlerFunc

	// CreateAsync 合成素材(异步)
	// @Tags API.img
	// @Router /api/img/createAsync [post]
	CreateAsync() core.HandlerFunc

	// TaskList 查询视频合成结果接口
	// @Tags API.img
	// @Router /api/img/task [get]
	TaskList() core.HandlerFunc
}


type handler struct {
	logger *zap.Logger
	service img.Service
}