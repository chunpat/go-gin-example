module github.com/FromChinaBoy/go-gin-example

go 1.12

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.0
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/creack/pty v1.1.9 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.51.1
	github.com/go-openapi/spec v0.19.5 // indirect
	github.com/go-openapi/swag v0.19.6 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20191106031601-ce3c9ade29de // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/kr/pty v1.1.8 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.0 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/robfig/cron v1.2.0
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.6.5
	github.com/unknwon/com v1.0.1
	golang.org/x/crypto v0.0.0-20200117160349-530e935923ad // indirect
	golang.org/x/mod v0.2.0 // indirect
	golang.org/x/net v0.0.0-20200114155413-6afb5195e5aa // indirect
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e // indirect
	golang.org/x/sys v0.0.0-20200121082415-34d275377bf9 // indirect
	golang.org/x/tools v0.0.0-20200122002620-f88bd7050267 // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/ini.v1 v1.51.1 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace (
	github.com/FromChinaBoy/go-gin-example => /work/zzhpeng/go-gin-example/
	github.com/FromChinaBoy/go-gin-example/conf => /work/zzhpeng/go-gin-example/pkg/conf
	github.com/FromChinaBoy/go-gin-example/middleware => /work/zzhpeng/go-gin-example/middleware
	github.com/FromChinaBoy/go-gin-example/middleware/jwt => /work/zzhpeng/go-gin-example/middleware/jwt
	github.com/FromChinaBoy/go-gin-example/models => /work/zzhpeng/go-gin-example/models
	github.com/FromChinaBoy/go-gin-example/pkg/e => /work/zzhpeng/go-gin-example/pkg/e
	github.com/FromChinaBoy/go-gin-example/pkg/file => /work/zzhpeng/go-gin-example/pkg/file
	github.com/FromChinaBoy/go-gin-example/pkg/logging => /work/zzhpeng/go-gin-example/pkg/logging
	github.com/FromChinaBoy/go-gin-example/pkg/setting => /work/zzhpeng/go-gin-example/pkg/setting
	github.com/FromChinaBoy/go-gin-example/pkg/util => /work/zzhpeng/go-gin-example/pkg/util
	github.com/FromChinaBoy/go-gin-example/routers => /work/zzhpeng/go-gin-example/routers
	github.com/FromChinaBoy/go-gin-example/routers/api => /work/zzhpeng/go-gin-example/routers/api
)
