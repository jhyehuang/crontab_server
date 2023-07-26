package models

import (
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"gorm.io/gorm"
)

func SetupDatabase(db *gorm.DB) (*gorm.DB, error) {

	if err := migrateSchemas(db); err != nil {
		log.Errorf("migrate schemas failed: ", err)
		return nil, err
	}
	return db, nil
}

func migrateSchemas(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&BlockRewardReleaseDetail{},
		&BlockRewardReleaseDayDetail{},
		&BlockRewardAllocRecord{},
		&BlockRewardReleaseLog{},
		&EventSync{},
		&FundDetails{},
		&MinerInfo{},
		&SPPenalty{},
		&RaisingAssetPack{},
		&AssetPackSector{},
		&CustomerAssetPack{},
		&SectorReleaseRecord{},
		&SectorSealedProgress{},
		&RaisingPlan{},
		&ServiceProvider{},
		&User{},
		&Rewards{},
		&Penalty{},
		&PushContractTxRecord{},
		&BannerBackground{},
		&RaiseSyncCount{},
	); err != nil {
		log.Errorf("migrate schemas failed: ", err)
		return err
	}
	return nil
}
