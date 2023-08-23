package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/huobirdcenter/huobi_golang/pkg/client"
	"github.com/jhyehuang/crontab_server/src/goft"
	"github.com/jhyehuang/crontab_server/src/models"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"github.com/jhyehuang/crontab_server/src/result"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

type RaisingPlanController struct {
	MarketClient *client.MarketClient `inject:"-"`
	RedisClient  *redis.Client        `inject:"-"`
	Db           *gorm.DB             `inject:"-"`
}

func NewRaisingPlanController() *RaisingPlanController {
	return &RaisingPlanController{}
}

func (this *RaisingPlanController) RaiseTestSealProgress(c *gin.Context) goft.Json {
	raising_id := c.Query("raising_id")
	contract_address := c.Query("contract_address")
	released_balance := c.Query("released_balance") // 可用
	will_release := c.Query("will_release")
	power := c.Query("power")
	pledge := c.Query("pledge")

	log.Infof("raising_id:%+v,contract_address:%+v,released_balance:%+v,will_release:%+v,power:%+v,pledge:%+v", raising_id, contract_address, released_balance, will_release, power, pledge)
	if raising_id == "" {
		return result.ParamError.WithParamError(errors.New("raising_id is empty"))
	}
	raising_id_i, err := decimal.NewFromString(raising_id)
	if err != nil {
		return result.ParamError.WithParamError(err)
	}

	released_balance_f, err := strconv.ParseFloat(released_balance, 64)
	if err != nil {
		return result.ParamError.WithParamError(err)
	}
	will_release_f, err := strconv.ParseFloat(will_release, 64)
	if err != nil {
		return result.ParamError.WithParamError(err)
	}
	power_f, err := strconv.ParseUint(power, 10, 64)
	if err != nil {
		return result.ParamError.WithParamError(err)
	}
	pledge_f, err := strconv.ParseFloat(pledge, 64)
	if err != nil {
		return result.ParamError.WithParamError(err)
	}

	released_balance_f = released_balance_f * 1e18
	will_release_f = will_release_f * 1e18
	pledge_f = pledge_f * 1e18
	power_f = power_f * 1024 * 1024 * 1024

	// 使用数据库事务
	tx := this.Db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 更新资产包表
	var updateMap = map[string]interface{}{
		"total_pledge_amount": pledge_f,
		"total_power":         power_f,
	}
	log.Infof("asset_pack_id: %+v ,updateMap:%+v", raising_id, updateMap)
	if err := tx.Model(&models.RaisingAssetPack{}).Where("asset_pack_id = ?", raising_id).Updates(updateMap).Error; err != nil {
		log.Error(err)
		return result.SystemError.WithParamError(err)
	}
	// 查询募集计划
	var raisingPlan models.RaisingPlan
	if err := tx.Where("raising_id = ?", raising_id).First(&raisingPlan).Error; err != nil {
		log.Errorf("query raisingPlan error:%+v", err)
		return result.SystemError.WithParamError(err)
	}

	// 随机一个 小于611440 但是6开头的值
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(611440)
	for randNum > 611440 || randNum < 600000 {
		randNum = rand.Intn(611440)
	}
	// 插入奖励释放记录
	var rewardRelease models.BlockRewardReleaseDetail
	rewardRelease.MinerId = raisingPlan.MinerId
	rewardRelease.AssetPackId = raising_id_i.BigInt().Uint64()
	rewardRelease.ReleasedReward = decimal.NewFromFloat(released_balance_f)
	rewardRelease.Height = uint64(randNum)

	if err := tx.Create(&rewardRelease).Error; err != nil {
		log.Errorf("create rewardRelease error:%+v", err)
		return result.SystemError.WithParamError(err)
	}
	var rewardWillRelease models.BlockRewardReleaseDetail
	rewardWillRelease.MinerId = raisingPlan.MinerId
	rewardWillRelease.AssetPackId = raising_id_i.BigInt().Uint64()
	rewardWillRelease.ReleasedReward = decimal.NewFromFloat(will_release_f)
	rewardWillRelease.Height = uint64(randNum) * 100

	if err := tx.Create(&rewardWillRelease).Error; err != nil {
		log.Errorf("create rewardRelease error:%+v", err)
		return result.SystemError.WithParamError(err)
	}
	tx.Commit()
	return result.OK
}

func (this *RaisingPlanController) Name() string {
	return "SProvierController"
}

func (this *RaisingPlanController) Build(goft *goft.Goft) {

	// todo。。。。 上线去掉
	goft.Handle("GET", "/v2/raise-test-seal-progress", this.RaiseTestSealProgress)

}
