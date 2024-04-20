package router

import (
	"github.com/gin-gonic/gin"
	"net/http/pprof"
)

/**
Go 的 pprof 工具，提供了不同类型的性能分析数据。以下是每个路由的详细作用说明：

debug.GET("pprof/", gin.WrapF(pprof.Index))：
作用：显示当前已注册的 pprof 报告的列表。
示例用途：用于查看所有可用的 pprof 报告类型。

debug.GET("pprof/cmdline", gin.WrapF(pprof.Cmdline))：
作用：返回运行进程的命令行调用。
示例用途：查看运行进程的命令行参数，可能有助于调试或了解进程的运行环境。

debug.GET("pprof/profile", gin.WrapF(pprof.Profile))：
作用：返回一个 CPU 性能分析报告。
示例用途：用于检查应用程序的 CPU 使用情况，找出性能瓶颈。

debug.GET("pprof/symbol", gin.WrapF(pprof.Symbol))：
作用：返回当前程序的符号信息。
示例用途：用于查看程序的符号信息，可能有助于调试或了解程序的运行情况。

debug.GET("pprof/trace", gin.WrapF(pprof.Trace))：
作用：返回一个运行时的跟踪报告。
示例用途：用于跟踪程序的执行路径，分析调用关系和函数执行时间等信息。

debug.GET("/pprof/allocs", gin.WrapH(pprof.Handler("allocs")))：
作用：返回过去内存分配的样本。
示例用途：用于查看应用程序的内存分配情况，了解内存使用情况。

debug.GET("/pprof/heap", gin.WrapH(pprof.Handler("heap")))：
作用：返回活动对象的内存分配示例。
示例用途：用于查看应用程序的堆内存使用情况，了解对象分配情况。

debug.GET("/pprof/goroutine", gin.WrapH(pprof.Handler("goroutine")))：
作用：返回所有当前 goroutine 的堆栈跟踪。
示例用途：用于查看应用程序中所有当前运行的 goroutine，了解并发情况。

debug.GET("/pprof/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))：
作用：返回导致创建新操作系统线程的堆栈跟踪。
示例用途：用于查看应用程序中创建线程的情况，了解线程使用情况。

debug.GET("/pprof/mutex", gin.WrapH(pprof.Handler("mutex")))：
作用：返回争用互斥锁持有者的堆栈跟踪。
示例用途：用于分析应用程序中的互斥锁情况，了解锁的竞争情况。
*/

func profiling(route *gin.Engine) {
	debug := route.Group("/debug/")
	{
		// 性能分析主页
		debug.GET("pprof/", gin.WrapF(pprof.Index))

		debug.GET("pprof/cmdline", gin.WrapF(pprof.Cmdline))

		debug.GET("pprof/profile", gin.WrapF(pprof.Profile))

		debug.GET("pprof/symbol", gin.WrapF(pprof.Symbol))

		debug.GET("pprof/trace", gin.WrapF(pprof.Trace))

		debug.GET("/pprof/allocs", gin.WrapH(pprof.Handler("allocs")))

		debug.GET("/pprof/heap", gin.WrapH(pprof.Handler("heap")))

		debug.GET("/pprof/goroutine", gin.WrapH(pprof.Handler("goroutine")))

		debug.GET("/pprof/threadcreate", gin.WrapH(pprof.Handler("threadcreate")))

		debug.GET("/pprof/mutex", gin.WrapH(pprof.Handler("mutex")))

	}

}
