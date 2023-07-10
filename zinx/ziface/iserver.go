package ziface

// 定义服务器的接口
type IServer interface {
	// 启动服务器
	Start()
	// 运行服务器
	Server()
	// 停止服务器
	Stop()

	// 路由，给当前服务注册一个路由方法，供客户端的连接处理使用
	AddRouter(router IRouter)
}