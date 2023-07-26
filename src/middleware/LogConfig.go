package middleware

import "github.com/jhyehuang/crontab_server/src/configs"

type LogConfig struct {
	Level string `json:"level"`
}

func NewLogConfig() *LogConfig {
	return &LogConfig{Level: configs.GetValue(configs.LogLevel, configs.Get().Log.Level)}
}
