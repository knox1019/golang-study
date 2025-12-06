package main

import (
	"net/http"
	"time"
	"io/ioutil"
	"strings"
	"strconv"
	"github.com/gin-gonic/gin"
)

// 股票数据结构体（添加 PrevClose 字段，与赋值对应）
type StockData struct {
	Code      string  `json:"code"`      // 股票代码
	Name      string  `json:"name"`      // 股票名称
	Price     float64 `json:"price"`     // 当前价格
	Change    float64 `json:"change"`    // 涨跌幅(%)
	Open      float64 `json:"open"`      // 今开
	High      float64 `json:"high"`      // 最高
	Low       float64 `json:"low"`       // 最低
	PrevClose float64 `json:"prev_close"`// 昨收价（新增字段，匹配赋值）
	UpdateAt  string  `json:"update_at"` // 更新时间
}

func main() {
	// 初始化 Gin 引擎（开发模式，生产环境用 gin.ReleaseMode）
	r := gin.Default()

	// 1. 静态文件服务（用于部署前端网页）
	r.Static("/static", "./static")
	// 访问根路径时，返回静态网页
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	// 2. API 接口：获取实时股票数据（对接腾讯 API）
	r.GET("/api/stock", func(c *gin.Context) {
		// 获取请求参数（股票代码，默认上证指数）
		stockCode := c.DefaultQuery("code", "sh000001")

		// 调用腾讯财经 API 获取数据
		resp, err := http.Get("https://qt.gtimg.cn/q=" + stockCode)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "获取数据失败"})
			return
		}
		defer resp.Body.Close()

		// 读取并解析返回数据
		body, _ := ioutil.ReadAll(resp.Body)
		dataStr := string(body)
		dataArr := strings.Split(dataStr, "~")

		// 解析关键数据（注意：腾讯 API 返回的数组长度需足够，避免索引越界）
		if len(dataArr) < 9 {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "数据解析失败"})
			return
		}

		name := dataArr[1]
		price, _ := strconv.ParseFloat(dataArr[3], 64)
		change, _ := strconv.ParseFloat(dataArr[4], 64)
		open, _ := strconv.ParseFloat(dataArr[5], 64)
		high, _ := strconv.ParseFloat(dataArr[6], 64)
		low, _ := strconv.ParseFloat(dataArr[7], 64)
		prevClose, _ := strconv.ParseFloat(dataArr[8], 64)

		// 构造股票数据结构体
		stock := StockData{
			Code:      stockCode,
			Name:      name,
			Price:     price,
			Change:    change,
			Open:      open,
			High:      high,
			Low:       low,
			PrevClose: prevClose, // 现在字段存在，不会报错
			UpdateAt:  time.Now().Format("15:04:05"),
		}

		// 返回 JSON 格式结果
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    stock,
		})
	})

	// 3. 启动服务（端口 9090）
	err := r.Run(":9090")
	if err != nil {
		panic("服务启动失败：" + err.Error())
	}
}