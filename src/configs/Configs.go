package configs

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"os"

	"strings"

	"github.com/filecoin-project/go-address"
	"github.com/fsnotify/fsnotify"
	"github.com/jhyehuang/crontab_server/src/global/constant"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"github.com/spf13/viper"
)

var config = new(Config)

type Config struct {
	Server struct {
		Name string `yaml:"name"`
		Port string `yaml:"port"`
	} `yaml:"server"`

	MySQL struct {
		Username     string `yaml:"username" required:"false"`
		Password     string `yaml:"password" required:"false"`
		Host         string `yaml:"host" required:"false"`
		Port         string `yaml:"port" required:"false"`
		Db           string `yaml:"db" required:"false"`
		MaxIdleConns int    `yaml:"maxIdleConns"`
		MaxOpenConns int    `yaml:"maxOpenConns"`
	} `yaml:"mysql"`

	Redis struct {
		Host     string `yaml:"host" required:"false"`
		Port     string `yaml:"port" required:"false"`
		Database int    `yaml:"database"`
		Password string `yaml:"password" required:"false"`
		PoolSize int    `yaml:"poolSize"`
		Timeout  string `yaml:"timeout"`
	} `yaml:"redis"`

	Log struct {
		Level string `yaml:"level"`
	} `yaml:"log"`

	Influxdb struct {
		Url    string `yaml:"url" required:"false"`
		Token  string `yaml:"token" required:"false"`
		Bucket string `yaml:"bucket"`
		Org    string `yaml:"org" required:"false"`
	} `yaml:"influxdb"`

	Lotus struct {
		Url       string `yaml:"url" required:"false"`
		AuthToken string `yaml:"authToken" required:"false"`
		NameSpace string `yaml:"nameSpace"`
	} `yaml:"lotus"`

	Glif struct {
		UrlV0     string `yaml:"urlV0" required:"false"`
		UrlV1     string `yaml:"urlV1" required:"false"`
		AuthToken string `yaml:"authToken" required:"false"`
	} `yaml:"glif"`

	FilData struct {
		Url       string `yaml:"url" required:"false"`
		AuthToken string `yaml:"authToken" required:"false"`
	} `yaml:"fildata"`

	FilFevm struct {
		HttpsUrl        string `yaml:"httpsUrl" required:"false"`
		WsUrl           string `yaml:"wsUrl" required:"false"`
		AuthToken       string `yaml:"authToken" required:"false"`
		FactoryContract string `yaml:"factoryContract" required:"false"`
	} `yaml:"filfevm"`

	FIL struct {
		Network string `yaml:"network" required:"false"`
	} `yaml:"fil"`

	AWS struct {
		AwsAccessKeyId     string `yaml:"awsAccessKeyId" required:"false"`
		AwsSecretAccessKey string `yaml:"awsSecretAccessKey" required:"false"`
		AwsDefaultRegion   string `yaml:"awsDefaultRegion" required:"false"`
		StaticResBucket    string `yaml:"StaticResBucket" required:"false"`
	} `yaml:"aws"`

	InternalService struct {
		Miner string `yaml:"miner" required:"false"`
		Key   string `yaml:"key" required:"false"`
	}
}

var (

	//go:embed application.yaml
	defaultConfigs []byte
)

func init() {
	InitConfig()
}

func InitConfig() {
	var r io.Reader
	var filename string
	fmt.Printf("ENV_MODE: %s \n", os.Getenv(constant.ENV_MODE))

	filename = "application"
	r = bytes.NewReader(defaultConfigs)

	fmt.Printf("filename: %s \n", filename)

	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func InitConfigFile(configPath *string) {
	// 读取配置文件
	viper.SetConfigFile(*configPath) // 配置文件路径
	viper.SetConfigType("yaml")      // 配置文件类型
	viper.AutomaticEnv()
	viper.SetEnvPrefix("FILFI")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	// 解析配置文件
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})

	if GetValue(FilNetwork, Get().FIL.Network) == "testnet" {
		address.CurrentNetwork = address.Testnet
	} else {
		address.CurrentNetwork = address.Mainnet
	}
}

func Get() Config {
	return *config
}

func GetValue(key string, defaultValue string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	} else {
		return defaultValue
	}
}
