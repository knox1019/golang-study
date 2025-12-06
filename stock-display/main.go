package main

import (
    "encoding/json"
    "fmt"
    "html/template"
    "io"
    "log"
    "net/http"
    "os"
    "sync"
    "time"
	"strings"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

// 股票数据结构
type StockData struct {
    Code         string  `json:"code"`
    Name         string  `json:"name"`
    LatestPrice  float64 `json:"latest_price"`
    Change       float64 `json:"change"`
    ChangePct    float64 `json:"change_pct"`
    MainNetInflow float64 `json:"main_net_inflow"`
    Market       string  `json:"market"`
    Timestamp    int64   `json:"timestamp"`
    Volume       int64   `json:"volume"`      // 成交量
    Turnover     float64 `json:"turnover"`    // 成交额
}

// 东方财富API响应结构
type EastMoneyRealtimeResp struct {
    Data struct {
        F43  float64 `json:"f43"`  // 最新价 (需要除以100)
        F57  string  `json:"f57"`  // 股票代码
        F58  string  `json:"f58"`  // 股票名称
        F169 float64 `json:"f169"` // 涨跌幅
        F170 float64 `json:"f170"` // 涨跌额
        F47  float64 `json:"f47"`  // 成交量 (手)
        F48  float64 `json:"f48"`  // 成交额 (元)
    } `json:"data"`
    Rc  int    `json:"rc"`
    Msg string `json:"msg"`
}

// 配置结构
type Config struct {
    Stocks []struct {
        Name string `json:"name"`
        Code string `json:"code"`
        Market string `json:"market"` // "sh" for Shanghai, "sz" for Shenzhen
    } `json:"stocks"`
    RefreshInterval int `json:"refresh_interval"` // seconds
    WebPort         int `json:"web_port"`
}

var (
    config         Config
    stockDataMap   = make(map[string]StockData)
    dataMutex      sync.RWMutex
    upgrader       = websocket.Upgrader{
        CheckOrigin: func(r *http.Request) bool {
            return true
        },
    }
    wsClients      = make(map[*websocket.Conn]bool)
    wsMutex        sync.RWMutex
)

func main() {
    // 加载配置
    loadConfig()
    
    // 初始化Gin
    if config.WebPort == 0 {
        config.WebPort = 9090
    }
    
    gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
    
    // 设置模板函数
    r.SetFuncMap(template.FuncMap{
        "formatPrice": func(price float64) string {
            return fmt.Sprintf("%.2f", price)
        },
        "formatChange": func(change float64) string {
            if change > 0 {
                return fmt.Sprintf("+%.2f", change)
            }
            return fmt.Sprintf("%.2f", change)
        },
        "formatChangePct": func(pct float64) string {
            if pct > 0 {
                return fmt.Sprintf("+%.2f%%", pct)
            }
            return fmt.Sprintf("%.2f%%", pct)
        },
        "formatInflow": func(inflow float64) string {
            if inflow > 0 {
                return fmt.Sprintf("+%.2f万", inflow)
            }
            return fmt.Sprintf("%.2f万", inflow)
        },
        "formatVolume": func(volume int64) string {
            if volume >= 100000000 {
                return fmt.Sprintf("%.2f亿手", float64(volume)/100000000)
            } else if volume >= 10000 {
                return fmt.Sprintf("%.2f万手", float64(volume)/10000)
            }
            return fmt.Sprintf("%d手", volume)
        },
        "formatTurnover": func(turnover float64) string {
            if turnover >= 100000000 {
                return fmt.Sprintf("%.2f亿元", turnover/100000000)
            } else if turnover >= 10000 {
                return fmt.Sprintf("%.2f万元", turnover/10000)
            }
            return fmt.Sprintf("%.2f元", turnover)
        },
        "getColorClass": func(change float64) string {
            if change > 0 {
                return "positive"
            } else if change < 0 {
                return "negative"
            }
            return "neutral"
        },
        "formatTime": func(t time.Time) string {
            return t.Format("2006-01-02 15:04:05")
        },
        "formatTimestamp": func(ts int64) string {
            return time.Unix(ts, 0).Format("15:04:05")
        },
        "upper": func(s string) string {
            return strings.ToUpper(s)
        },
    })
    
    // 加载模板
    r.LoadHTMLGlob("templates/*")
    
    // 静态文件服务
    r.Static("/static", "./static")
    
    // 路由设置
    r.GET("/", func(c *gin.Context) {
        dataMutex.RLock()
        stocks := make([]StockData, 0, len(stockDataMap))
        for _, stock := range stockDataMap {
            stocks = append(stocks, stock)
        }
        dataMutex.RUnlock()
        
        c.HTML(http.StatusOK, "index.html", gin.H{
            "Stocks": stocks,
            "Config": config,
            "Now":    time.Now(),
        })
    })
    
    r.GET("/api/stocks", getStocksAPI)
    r.GET("/api/stock/:market/:code", getSingleStockAPI)
    r.GET("/ws", handleWebSocket)
    
    // 启动数据更新协程
    go updateStockData()
    go broadcastUpdates()
    
    // 启动HTTP服务器
    log.Printf("服务器启动在 http://localhost:%d", config.WebPort)
    if err := r.Run(fmt.Sprintf(":%d", config.WebPort)); err != nil {
        log.Fatal("服务器启动失败:", err)
    }
}

func loadConfig() {
    // 默认配置
    config = Config{
        Stocks: []struct {
            Name   string `json:"name"`
            Code   string `json:"code"`
            Market string `json:"market"`
        }{
            {"贵州茅台", "600519", "sh"},
            {"腾讯控股", "00700", "hk"},
            {"宁德时代", "300750", "sz"},
            {"招商银行", "600036", "sh"},
            {"中国平安", "601318", "sh"},
        },
        RefreshInterval: 10,
        WebPort:         9090,
    }
    
    // 尝试从配置文件加载
    if file, err := os.Open("config.json"); err == nil {
        defer file.Close()
        if err := json.NewDecoder(file).Decode(&config); err != nil {
            log.Printf("读取配置文件失败，使用默认配置: %v", err)
        }
    }
}

func updateStockData() {
    ticker := time.NewTicker(time.Duration(config.RefreshInterval) * time.Second)
    defer ticker.Stop()
    
    // 立即执行一次
    fetchAllStocks()
    
    for range ticker.C {
        fetchAllStocks()
    }
}

func fetchAllStocks() {
    for _, stock := range config.Stocks {
        go func(market, code string) {
            if quote, err := getStockQuote(market, code); err == nil {
                dataMutex.Lock()
                stockDataMap[market+"-"+code] = quote
                dataMutex.Unlock()
            } else {
                log.Printf("获取股票数据失败 %s-%s: %v", market, code, err)
            }
        }(stock.Market, stock.Code)
    }
}

func getStockQuote(market, code string) (StockData, error) {
    // 转换为东方财富的市场代码
    var eastMoneyMarket string
    switch market {
    case "sh":
        eastMoneyMarket = "1"
    case "sz":
        eastMoneyMarket = "0"
    case "hk":
        eastMoneyMarket = "116"
    default:
        eastMoneyMarket = "1"
    }
    
    secid := fmt.Sprintf("%s.%s", eastMoneyMarket, code)
    url := "https://push2.eastmoney.com/api/qt/stock/get"
    
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return StockData{}, err
    }
    
    q := req.URL.Query()
    q.Add("secid", secid)
    q.Add("fields", "f43,f57,f58,f169,f170,f47,f48")
    q.Add("ut", "fa5fd1943c7b386f172d6893dbfba10b")
    q.Add("fltt", "2")
    q.Add("invt", "2")
    req.URL.RawQuery = q.Encode()
    
    req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
    req.Header.Set("Referer", "https://quote.eastmoney.com/")
    
    client := &http.Client{Timeout: 5 * time.Second}
    resp, err := client.Do(req)
    if err != nil {
        return StockData{}, err
    }
    defer resp.Body.Close()
    
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return StockData{}, err
    }
    
    var apiResp EastMoneyRealtimeResp
    if err := json.Unmarshal(body, &apiResp); err != nil {
        return StockData{}, err
    }
    
    if apiResp.Rc != 0 {
        return StockData{}, fmt.Errorf("API错误: %s", apiResp.Msg)
    }
    
    // 计算成交量和成交额
    volume := int64(apiResp.Data.F47) // 成交量 (手)
    turnover := apiResp.Data.F48 // 成交额 (元)
    
    return StockData{
        Code:        code,
        Name:        apiResp.Data.F58,
        LatestPrice: apiResp.Data.F43 / 100.0,
        Change:      apiResp.Data.F170 / 100.0,
        ChangePct:   apiResp.Data.F169,
        Market:      market,
        Timestamp:   time.Now().Unix(),
        Volume:      volume,
        Turnover:    turnover,
    }, nil
}

func getStocksAPI(c *gin.Context) {
    dataMutex.RLock()
    stocks := make([]StockData, 0, len(stockDataMap))
    for _, stock := range stockDataMap {
        stocks = append(stocks, stock)
    }
    dataMutex.RUnlock()
    
    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "data":    stocks,
        "time":    time.Now().Unix(),
    })
}

func getSingleStockAPI(c *gin.Context) {
    market := c.Param("market")
    code := c.Param("code")
    
    dataMutex.RLock()
    stock, exists := stockDataMap[market+"-"+code]
    dataMutex.RUnlock()
    
    if exists {
        c.JSON(http.StatusOK, gin.H{
            "success": true,
            "data":    stock,
        })
    } else {
        c.JSON(http.StatusNotFound, gin.H{
            "success": false,
            "error":   "股票不存在",
        })
    }
}

func handleWebSocket(c *gin.Context) {
    conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Printf("WebSocket升级失败: %v", err)
        return
    }
    defer conn.Close()
    
    wsMutex.Lock()
    wsClients[conn] = true
    wsMutex.Unlock()
    
    // 发送当前数据
    dataMutex.RLock()
    stocks := make([]StockData, 0, len(stockDataMap))
    for _, stock := range stockDataMap {
        stocks = append(stocks, stock)
    }
    dataMutex.RUnlock()
    
    if err := conn.WriteJSON(gin.H{
        "type": "init",
        "data": stocks,
    }); err != nil {
        log.Printf("发送初始数据失败: %v", err)
    }
    
    // 保持连接
    for {
        messageType, _, err := conn.ReadMessage()
        if err != nil {
            break
        }
        if messageType == websocket.CloseMessage {
            break
        }
    }
    
    wsMutex.Lock()
    delete(wsClients, conn)
    wsMutex.Unlock()
}

func broadcastUpdates() {
    ticker := time.NewTicker(time.Duration(config.RefreshInterval) * time.Second)
    defer ticker.Stop()
    
    for range ticker.C {
        dataMutex.RLock()
        stocks := make([]StockData, 0, len(stockDataMap))
        for _, stock := range stockDataMap {
            stocks = append(stocks, stock)
        }
        dataMutex.RUnlock()
        
        wsMutex.RLock()
        for client := range wsClients {
            go func(conn *websocket.Conn) {
                if err := conn.WriteJSON(gin.H{
                    "type": "update",
                    "data": stocks,
                    "time": time.Now().Unix(),
                }); err != nil {
                    log.Printf("广播数据失败: %v", err)
                }
            }(client)
        }
        wsMutex.RUnlock()
    }
}
