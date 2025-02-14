// Copyright 2021 CloudJ Company Limited. All rights reserved.

package handlers

import (
	"cloudiac/portal/apps"
	"cloudiac/portal/libs/ctx"
	"cloudiac/portal/models/forms"
)

// PortalSystemStatusSearch 查询系统状态
// @Summary 查询系统状态
// @Description 查询系统状态
// @Tags 系统状态
// @Accept  json
// @Produce  json
// @Security AuthToken
// @Success 200 {object} []apps.SystemStatusResp
// @Router /systems/status [get]
func PortalSystemStatusSearch(c *ctx.GinRequest) {
	c.JSONResult(apps.SystemStatusSearch())
}

func ConsulKVSearch(c *ctx.GinRequest) {
	key := c.Query("key")
	c.JSONResult(apps.ConsulKVSearch(key))
}

// RunnerSearch 查询runner列表
// @Summary 查询runner列表
// @Description 查询runner列表
// @Tags runner
// @Accept  json
// @Produce  json
// @Security AuthToken
// @Success 200
// @Router /runners [get]
func RunnerSearch(c *ctx.GinRequest) {
	c.JSONResult(apps.RunnerSearch())
}

// ConsulTagUpdate 修改服务标签
// @Summary 修改服务标签
// @Description 修改服务标签
// @Tags 系统状态
// @Accept  json
// @Produce  json
// @Security AuthToken
// @Param data body forms.ConsulTagUpdateForm true "tag信息"
// @Success 200
// @Router /consul/tags [put]
func ConsulTagUpdate(c *ctx.GinRequest) {
	form := forms.ConsulTagUpdateForm{}
	if err := c.Bind(&form); err != nil {
		return
	}
	c.JSONResult(apps.ConsulTagUpdate(form))
}
