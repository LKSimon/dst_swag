package video

import (
	"context"
	"dst_swag_demo/internal/pkg/core"
)

type createSyncRequest struct{}

type createSyncResponse struct{}

// CreateSync 合成素材(同步)
// @Summary 合成素材(同步)
// @Description 合成素材(同步)
// @Tags API.video
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createSyncRequest true "请求信息"
// @Success 200 {object} createSyncResponse
// @Failure 400 {object} code.Failure
// @Router /api/video/createSync [post]
func (h *handler) CreateSync() core.HandlerFunc {
	return func(ctx context.Context) {

	}
}
