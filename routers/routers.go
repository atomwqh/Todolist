package routers

import (
	"bubble_todolist/controller"
	"github.com/gin-gonic/gin"
)

func SetUpRouter() *gin.Engine {
	r := gin.Default()
	//静态文件位置
	r.Static("/static", "static")
	//模板文件位置
	r.LoadHTMLGlob("templates/*")

	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateATodo)
		//查看所有待办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return r
}
