package main

import (
	"fmt"
	"github.com/devfeel/dotweb"
	"github.com/devfeel/dotweb/framework/file"
	"strconv"
)

func main() {
	//初始化DotServer
	app := dotweb.New()

	//设置dotserver日志目录
	app.SetLogPath(file.GetCurrentDirectory())

	//设置Debug开关
	app.SetEnabledDebug(true)

	//设置路由
	InitRoute(app.HttpServer)

	//启动 监控服务
	//pprofport := 8081
	//go app.StartPProfServer(pprofport)

	//全局容器
	app.AppContext.Set("gstring", "gvalue")
	app.AppContext.Set("gint", 1)

	// 开始服务
	port := 8080
	fmt.Println("dotweb.StartServer => " + strconv.Itoa(port))
	err := app.StartServer(port)
	fmt.Println("dotweb.StartServer error => ", err)
}

//you can curl http://127.0.0.1:8080/
func Index(ctx *dotweb.HttpContext) {
	gstring := ctx.AppContext().GetString("gstring")
	gint := ctx.AppContext().GetInt("gint")
	ctx.WriteString("index -> " + gstring + ";" + strconv.Itoa(gint))
}

func InitRoute(server *dotweb.HttpServer) {
	server.Router().GET("/", Index)
}
