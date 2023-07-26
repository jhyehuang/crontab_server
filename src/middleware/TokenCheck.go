package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jhyehuang/crontab_server/src/global/redisKey"
	"github.com/jhyehuang/crontab_server/src/goft"
	"github.com/jhyehuang/crontab_server/src/result"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type TokenCheck struct {
	RedisClient *redis.Client `inject:"-"`
}

func NewTokenCheck() *TokenCheck {
	return &TokenCheck{}
}

func (this *TokenCheck) OnRequest(ctx *gin.Context) error {
	token := ctx.GetHeader(redisKey.TokenHeaderKey)
	if token == "" {
		zap.L().Error("token not found")
		goft.Throw(result.TokenInvalid, 200, ctx)
	}
	//存在 获取userId
	userId, err := this.RedisClient.Get(ctx, fmt.Sprintf(redisKey.UserTokenKey, token)).Result()
	if err != nil {
		goft.Throw(result.TokenInvalid, 200, ctx)
	}
	//重新设置过期时间
	this.RedisClient.Expire(ctx, token, redisKey.UserTokenExpireHours*time.Hour)
	userIdInt, _ := strconv.ParseInt(userId, 10, 64)
	ctx.Set(redisKey.UserTokenValue, userIdInt)
	return nil
}

func (this *TokenCheck) OnResponse(result interface{}) (interface{}, error) {
	return result, nil
}
