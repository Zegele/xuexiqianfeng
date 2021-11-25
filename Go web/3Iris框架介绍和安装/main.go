package main

func main() {
	/*
		Iris简介
		Iris是一款Go语言中用来开发web应用的框架，该框架支持编写一次并在任何地方以最小的机器功率运行，如Android、ios、Linux、Windows等。
		该框架只需要一个可执行的服务旧可以在平台上运行了。
		Iris框架以简单而强大的api被开发者所熟悉。Iris除了为开发者提供非常简单的访问方式外，还同样支持MVC。另外，用iris构建微服务也很容易。
		在Iris框架的官方网站上，被称为速度最快的Go后端开发框架。在Iris的网站文档上，列出了该框架具备的一些特点和框架特性，列举如下：
		1. 聚集高性能
		2. 健壮的静态路由支持和通配符子域名支持
		3. 视图系统支持超过5以上模板
		4. 支持定制时间的高可扩展性Websocket API
		5. 带有GC，内存& redis提供支持的会话
		6. 方便的中间件和插件
		7. 完整 REST API
		8. 能定制 HTTP错误
		9. 源码改变后自动加载（自动部署）
		等，参考Iris官方文档

		Iris学习渠道
		1.Iris官网：https://iris-go.com
		2. Iris框架源码地址：https://github.com/kataras/iris
		3. Iris框架中文学习文档：https://studyiris.com/doc/

		Iris框架安装

		环境要求：iris框架要求golang版本至少为1.8。go version 可以查看版本。

		安装Iris框架：
		go get -u github.com/kataras/iris
		下载到 GOPATH-src-github-kataras-iris

		Iris构造服务实例
		在安装完成Iris的源码后，我们就开始来编写最简单的一个Iris的服务。在Iris中，构建并运行一个服务实例需要两步：
			1. 通过iris.New()方法可以实例化一个应用服务对象app
			2. 通过Run方法开启端口监听服务，运行服务实例

		如下是一个最简单的服务案例Demo
		package main
		import"github.com/kataras/iris"
		func main(){
			app := iris.New()
			app.Run(iris.Addr(":8000"), iris.WithoutSeverError(iris.ErrSeverClosed))
		}

	*/
}
