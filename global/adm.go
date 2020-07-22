package global

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var GinEngine *gin.Engine

var Eloquent *gorm.DB
var Source string
var Driver string

var DBName string

var CasbinEnforcer *casbin.Enforcer

// go-admin Version Info
var Version string

func init()  {
	Version = "1.0.10"
}