package img

import (
"dst_swag/internal/pkg/core"
)



type templateListRequest struct {}

type templateListResponse struct {}

// TemplateList 视频模板列表接口
// @Summary 视频模板列表接口 
// @Description 视频模板列表接口 
// @Tags API.img 
// @Accept application/x-www-form-urlencoded 
// @Produce json 
// @Param Request body templateListRequest true "请求信息" 
// @Success 200 {object} templateListResponse 
// @Failure 400 {object} code.Failure 
// @Router /api/img/template [get] 
func (h *handler) TemplateList() core.HandlerFunc { 
 return func(ctx context.Context) {

}}