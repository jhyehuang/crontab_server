package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/jhyehuang/crontab_server/src/goft"
	"github.com/jhyehuang/crontab_server/src/result"
)

type HealthController struct {
	MarketClient *client.MarketClient `inject:"-"`
}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (this *HealthController) Health(c *gin.Context) goft.Json {
	return result.OK
}

func (this *HealthController) Name() string {
	return "HealthController"
}

func (this *HealthController) Build(goft *goft.Goft) {
	goft.Handle("GET", "/check", this.Health)
}
