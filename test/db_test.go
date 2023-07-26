package tests

import (
	"fmt"
	"github.com/jhyehuang/crontab_server/src/pkg/util"
	_ "github.com/lib/pq"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"testing"
	"time"
)

type ContentType int64
type Content struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Cid         util.DbCID  `gorm:"cid" json:"cid"`
	Name        string      `json:"name"`
	UserID      uint        `json:"userId" gorm:"index"`
	Description string      `json:"description"`
	Size        int64       `json:"size"`
	Type        ContentType `json:"type"`
	Active      bool        `json:"active"`
	Offloaded   bool        `json:"offloaded"`
	Replication int         `json:"replication"`

	// TODO: shift most of the 'state' booleans in here into a single state
	// field, should make reasoning about things much simpler
	AggregatedIn uint `json:"aggregatedIn" gorm:"index:,option:CONCURRENTLY"`
	Aggregate    bool `json:"aggregate"`

	Pinning bool   `json:"pinning"`
	PinMeta string `json:"pinMeta"`
	Replace bool   `json:"replace" gorm:"default:0"`
	Origins string `json:"origins"`

	Failed bool `json:"failed"`

	Location string `json:"location"`
	// TODO: shift location tracking to just use the ID of the shuttle
	// Also move towards recording content movement intentions in the database,
	// making that process more resilient to failures
	// LocID     uint   `json:"locID"`
	// LocIntent uint   `json:"locIntent"`

	// If set, this content is part of a split dag.
	// In such a case, the 'root' content should be advertised on the dht, but
	// not have deals made for it, and the children should have deals made for
	// them (unlike with aggregates)
	DagSplit  bool `json:"dagSplit"`
	SplitFrom uint `json:"splitFrom"`

	Collection string     `json:"collection"`
	IpfsCid    util.DbCID `json:"ipfsCid"`
	EnCid      util.DbCID `json:"enCid"`
}

func TestDb(t *testing.T) {
	dsn := "user=postgres password=XxvRmI6DwyBCUxCZ dbname=postgres host=swarm.crab.merak.sxxfuture.net port=5433 sslmode=disable"

	// 连接数据库
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Content{})

	var users []*Content
	if err := db.Table("contents").Find(&users, "user_id=12 and  created_at >= '2023-03-31' and aggregate is not true ").Error; err != nil {
		panic(err)
	}

	// 将查询结果写入Excel文件
	f := excelize.NewFile()
	sheetName := "Sheet1"
	rowIndex := 1
	for _, user := range users {
		f.SetCellValue(sheetName, "A"+strconv.Itoa(rowIndex), user.Name)
		f.SetCellValue(sheetName, "B"+strconv.Itoa(rowIndex), user.Cid.CID.String())
		rowIndex++
	}

	// 保存Excel文件
	err = f.SaveAs("contents_18156675885.xlsx")
	if err != nil {
		panic(err)
	}

	fmt.Println("成功将查询结果写入Excel文件！")

}
