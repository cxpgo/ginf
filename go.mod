module github.com/cxpgo/ginf

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/casbin/casbin v1.9.1
	github.com/casbin/casbin/v2 v2.20.0
	github.com/casbin/gorm-adapter/v3 v3.0.2
	github.com/cxpgo/golib v0.0.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/mojocn/base64Captcha v1.3.1
	github.com/satori/go.uuid v1.2.0
	github.com/swaggo/gin-swagger v1.3.0
	github.com/swaggo/swag v1.5.1
	github.com/unrolled/secure v1.0.8
	go.uber.org/zap v1.10.0
	gorm.io/gorm v1.20.8
)

replace github.com/cxpgo/golib v0.0.1 => ../golib
