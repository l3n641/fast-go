package main

import (
	"fast-go/bootstrap"
	"fast-go/cmd"
	"fast-go/global"
)

func main() {
	// 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")
	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()

	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	cmd.Execute()

}
