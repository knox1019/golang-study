package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// StockData 定义返回给前端的数据结构
type StockData struct {
	Price float64 `json:"price"` // 当前价格
	Time  string  `json:"time"`  // 当前时间
}

// 模拟一个初始价格
var currentPrice = 100.0

func main() {
	r := gin.Default()

	// 告诉 Gin 网页模板在哪里
	r.LoadHTMLGlob("templates/*")

	// 1. 访问首页，返回 HTML 页面
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// 2. API 接口：获取当前股票价格
	// 前端 JS 会每秒调用一次这个接口
	r.GET("/api/price", func(c *gin.Context) {
		// --- 这里模拟股价波动 ---
		// 生成一个 -0.5 到 +0.5 之间的随机波动
		fluctuation := (rand.Float64() - 0.5)
		currentPrice += fluctuation
		
		// 格式化时间，只取时分秒
		nowTime := time.Now().Format("15:04:05")

		// 构造结构体数据
		data := StockData{
			Price: currentPrice,
			Time:  nowTime,
		}

		// 返回 JSON 格式给前端
		c.JSON(http.StatusOK, data)
	})

	fmt.Println("服务已启动，请访问 http://localhost:8080")
	r.Run(":8080")
}