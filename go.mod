module github.com/FormChinaBoy/go-gin-example

go 1.12

require (
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.51.1
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/gopherjs/gopherjs v0.0.0-20191106031601-ce3c9ade29de // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/json-iterator/go v1.1.9 // indirect
	github.com/kr/pretty v0.2.0 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-isatty v0.0.11 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/smartystreets/goconvey v1.6.4 // indirect
	github.com/unknwon/com v1.0.1
	golang.org/x/sys v0.0.0-20200120151820-655fe14d7479 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
	gopkg.in/go-playground/validator.v9 v9.31.0 // indirect
	gopkg.in/ini.v1 v1.51.1 // indirect
	gopkg.in/yaml.v2 v2.2.7 // indirect
)

replace (
	github.com/FromChinaBoy/go-gin-example/ => /work/zzhpeng/go-gin-example/
	github.com/FromChinaBoy/go-gin-example/conf => /work/zzhpeng/go-gin-example/pkg/conf
	github.com/FromChinaBoy/go-gin-example/middleware => /work/zzhpeng/go-gin-example/middleware
	github.com/FromChinaBoy/go-gin-example/models => /work/zzhpeng/go-gin-example/models
	github.com/FromChinaBoy/go-gin-example/pkg/e => /work/zzhpeng/go-gin-example/pkg/e
	github.com/FromChinaBoy/go-gin-example/pkg/setting => /work/zzhpeng/go-gin-example/pkg/setting
	github.com/FromChinaBoy/go-gin-example/pkg/util => /work/zzhpeng/go-gin-example/pkg/util
	github.com/FromChinaBoy/go-gin-example/routers => /work/zzhpeng/go-gin-example/routers
	github.com/FromChinaBoy/go-gin-example/routers/api => /work/zzhpeng/go-gin-example/routers/api
)
