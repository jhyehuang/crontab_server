package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jhyehuang/crontab_server/src/goft"
	"github.com/jhyehuang/crontab_server/src/models"
	"github.com/jhyehuang/crontab_server/src/pkg/log"
	"github.com/jhyehuang/crontab_server/src/result"
	"gorm.io/gorm"
	"strings"
)

type UserController struct {
	Db *gorm.DB `inject:"-"`
}

func NewUserController() *UserController {
	return &UserController{}
}

func (this *UserController) Name() string {
	return "UserController"
}

// 创建发起人信息
func (this *UserController) Create(c *gin.Context) goft.Json {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		log.Errorf("bind json error: %v", err)
		return result.ParamError
	}

	if user.Name == "" {
		return result.ParamError.WithParamError("name is empty")
	}

	user.Address = strings.ToLower(user.Address)

	// 将记录插入数据库
	err = this.Db.Create(&user).Error
	if err != nil {
		log.Errorf("create user error: %v", err)
		return result.SystemError.WithParamError(err)
	}

	return result.OK.WithData(user)
}

// 修改发起人信息
func (this *UserController) Update(c *gin.Context) goft.Json {
	address := c.Param("address")

	address = strings.ToLower(address)
	updateMap := make(map[string]interface{})
	c.Bind(&updateMap)

	err := this.Db.Model(&models.User{}).Where("address = ?", address).Updates(updateMap).Error
	if err != nil {
		log.Errorf("update user error: %v", err)
		return result.SystemError.WithParamError(err)
	}
	return result.OK.WithData(updateMap)

}

// 查询发起人信息
func (this *UserController) Query(c *gin.Context) goft.Json {
	address := c.Param("address")
	address = strings.ToLower(address)
	var user models.User
	err := this.Db.Where("address = ?", address).First(&user).Error
	if err != nil {
		log.Errorf("query user error: %v", err)
		return result.UserNotFound.WithParamError(err)
	}

	return result.OK.WithData(user)

}

func (this *UserController) Build(goft *goft.Goft) {
	goft.Handle("POST", "/v2/create", this.Create)
	goft.Handle("POST", "/v2/update/:address", this.Update)
	goft.Handle("GET", "/v2/query/:address", this.Query)
}
