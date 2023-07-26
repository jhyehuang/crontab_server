package internal

import (
	"github.com/filecoin-project/lotus/build"
	"github.com/jhyehuang/crontab_server/src/global/constant"
	"github.com/jhyehuang/crontab_server/src/models"
	"github.com/jhyehuang/crontab_server/src/pkg/filscan"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"gorm.io/gorm"
	"math/big"
	"strconv"
)

// 计算年化收益率
func CalcIncomeRate(db *gorm.DB, rPList *models.RaisingPlan) uint64 {

	var packPower = big.NewInt(0)
	// 查询资产包信息
	var raisingAssetPack models.RaisingAssetPack
	if err := db.Model(&models.RaisingAssetPack{}).Where("asset_pack_id = ? ", rPList.RaisingId).First(&raisingAssetPack).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Warnf("查询资产包不存在:%v", err)
		} else {
			log.Errorf("查询资产包失败:%v", err)
			return 0
		}

	}
	if raisingAssetPack.AssetPackStatus == constant.ASSET_PACKAGE_STATUS_SEALED {
		packPower = raisingAssetPack.TotalPower.BigInt()
	} else {
		packPower = rPList.TargetPower.BigInt()
	}
	log.Infof("packPower:%+v", packPower)

	// 查询 GetStatChainInfo
	chainInfo, err := filscan.GetStatChainInfo()
	if err != nil {
		log.Errorf("查询链信息失败:%v", err)
		return 0
	}
	filPerTera, err := strconv.ParseFloat(chainInfo.Result.Data.FilPerTera, 64)
	filPerTera = filPerTera * float64(build.FilecoinPrecision)

	pledgePerTera, err := strconv.ParseFloat(chainInfo.Result.Data.PledgePerTera, 64)
	pledgePerTera = pledgePerTera * float64(build.FilecoinPrecision)

	ecIncomeRate := uint64(filPerTera * 360 / pledgePerTera * constant.CALC_ACCURACY)

	return ecIncomeRate
}
