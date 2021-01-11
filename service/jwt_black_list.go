package service

import (
	"errors"
	"github.com/cxpgo/ginf/model"
	"github.com/cxpgo/golib/gf/util/gconv"
	"github.com/cxpgo/golib/lib"
	"gorm.io/gorm"
	"time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = lib.GGorm.Create(&jwtList).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

func IsBlacklist(jwt string) bool {
	isNotFound := errors.Is(lib.GGorm.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error, gorm.ErrRecordNotFound)
	return !isNotFound
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func GetRedisJWT(userName string) (error, string) {
	//redisJWT, err = global.GVA_REDIS.Get(userName).Result()
	redisJWT, err := lib.GRedis.Do("get",userName)
	gconv.String(redisJWT)
	return err, redisJWT.(string)
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: userName string
//@return: err error, redisJWT string

func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := 60 * 60 * 24 * 7 * time.Second
	lib.GRedis.Do("set",userName,)
	_,err = lib.GRedis.DoWithTimeout(timer,"set",userName,jwt)
	//err = global.GVA_REDIS.Set(userName, jwt, timer).Err()
	return err
}
