package img

import (
"dst_swag/internal/pkg/core"
)



type createSyncRequest struct {}

type createSyncResponse struct {}

// CreateSync 合成素材(同步)
// @Summary 合成素材(同步) 
// @Description 合成素材(同步) 
// @Tags API.img 
// @Accept application/x-www-form-urlencoded 
// @Produce json 
// @Param Request body createSyncRequest true "请求信息" 
// @Success 200 {object} createSyncResponse 
// @Failure 400 {object} code.Failure 
// @Router /api/img/createSync [post] 
func (h *handler) CreateSync() core.HandlerFunc { 
 return func(ctx context.Context) {

}}