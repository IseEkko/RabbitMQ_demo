package main

import (
	"Iris_study/web/controller"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func main() {
	app := iris.New()
	app.RegisterView(iris.HTML("./web/view", ".html"))
	//注册控制器
	mvc.New(app.Party("/hello")).Handle(new(controller.MovieController))
	app.Run(
		iris.Addr("localhost:8080"),
	)
}
