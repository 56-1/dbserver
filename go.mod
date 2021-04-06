module dbserver

go 1.16

require github.com/beego/beego/v2 v2.0.1

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/smartystreets/goconvey v1.6.4
	judger v0.0.0-00010101000000-000000000000
)

replace judger => ../judger
