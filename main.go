package main

import (
	"flag"

	"github.com/jhyehuang/crontab_server/src/configs"
	"github.com/jhyehuang/crontab_server/src/controller"
	"github.com/jhyehuang/crontab_server/src/goft"
	. "github.com/jhyehuang/crontab_server/src/middleware"
	"github.com/jhyehuang/crontab_server/src/pkg/glif"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
)

func main() {

	configFile := flag.String("config", "", "请输入配置文件地址")
	// 解析命令行参数
	flag.Parse()
	configs.InitConfigFile(configFile)
	log.Infof("letsfil-job config: %+v", configs.Get())
	glif.InitNode()

	goft.Ignite().
		Config(NewDbConfig(),
			NewRedisClient(),
			NewServiceConfig(),
			NewHttpsFILClient(),
			NewWsFILClient()).
		Mount("health", controller.NewHealthController()).
		Mount("user", controller.NewUserController()).
		Mount("raising-plan", controller.NewRaisingPlanController()).
		//ClassTask(chainTask, "*/30 * * * * *", chainTask.CheckTxData). // 检查交易数据
		Launch()

}
