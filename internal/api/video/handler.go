package video

import (
	"dst_swag_demo/internal/pkg/core"
	"dst_swag_demo/internal/service/video"
	"go.uber.org/zap"
)

//var _ Handler = (*handler)(nil)

type Handler interface {
	i()

	// TemplateList 视频模板列表接口
	// @Tags API.video
	// @Router /api/video/template [get]
	TemplateList() core.HandlerFunc

	// CreateSync 合成素材(同步)
	// @Tags API.video
	// @Router /api/video/createSync [post]
	CreateSync() core.HandlerFunc

	// CreateAsync 合成素材(异步)
	// @Tags API.video
	// @Router /api/video/createAsync [post]
	CreateAsync() core.HandlerFunc

	// TaskList 查询视频合成结果接口
	// @Tags API.video
	// @Router /api/video/task [get]
	TaskList() core.HandlerFunc
}

type handler struct {
	logger       *zap.Logger
	videoService video.Service
}

//
//func New(logger *zap.Logger) Handler {
//	return &handler{
//		logger:       logger,
//		videoService: video.New(),
//	}
//}

//func (h *handler) i() {}
