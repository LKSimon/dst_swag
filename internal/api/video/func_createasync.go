package video

import (
	"context"
	"dst_swag/internal/pkg/core"
)

type createAsyncRequest struct{}

type createAsyncResponse struct{}

// CreateAsync 合成素材(异步)
// @Summary 合成素材(异步)
// @Description 合成素材(异步)
// @Tags API.video
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body createAsyncRequest true "请求信息"
// @Success 200 {object} createAsyncResponse
// @Failure 400 {object} code.Failure
// @Router /api/video/createAsync [post]
func (h *handler) CreateAsync() core.HandlerFunc {
	return func(ctx context.Context) {

	}
}
