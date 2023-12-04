package main

import (
	"bubble_todolist/DAO"
	"bubble_todolist/Models"
	"bubble_todolist/routers"
)

func main() {
	//创建数据库
	//连接数据库 bubble,这里导包的时候记得导入数据库的驱动

	err := DAO.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer DAO.Close() // 程序退出关闭数据库
	// 模型绑定
	DAO.Db.AutoMigrate(&Models.Todo{}) //todos
	// 注册路由
	r := routers.SetUpRouter()
	r.Run()
}
