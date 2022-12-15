package video

import (
	"context"
	"dst_swag/internal/pkg/core"
)

type taskListRequest struct{}

type taskListResponse struct{}

// TaskList 查询视频合成结果接口
// @Summary 查询视频合成结果接口
// @Description 查询视频合成结果接口
// @Tags API.video
// @Accept application/x-www-form-urlencoded
// @Produce json
// @Param Request body taskListRequest true "请求信息"
// @Success 200 {object} taskListResponse
// @Failure 400 {object} code.Failure
// @Router /api/video/task [get]
func (h *handler) TaskList() core.HandlerFunc {
	return func(ctx context.Context) {

	}
}
