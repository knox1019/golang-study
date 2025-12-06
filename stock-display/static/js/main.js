class StockApp {
    constructor() {
        this.soundEnabled = true;
        this.updateCount = 0;
        this.socket = null;
        this.speechSynth = window.speechSynthesis;
        this.currentUtterance = null;
        
        this.init();
    }
    
    init() {
        // 初始化事件监听器
        this.bindEvents();
        
        // 连接WebSocket
        this.connectWebSocket();
        
        // 更新服务器时间
        this.updateServerTime();
        setInterval(() => this.updateServerTime(), 1000);
    }
    
    bindEvents() {
        // 切换语音播报
        document.getElementById('toggleSound').addEventListener('click', (e) => {
            this.soundEnabled = !this.soundEnabled;
            const btn = e.target.closest('button');
            const icon = btn.querySelector('i');
            const span = btn.querySelector('span');
            
            if (this.soundEnabled) {
                icon.className = 'fas fa-volume-up';
                span.textContent = '开启';
                btn.classList.add('sound-btn');
            } else {
                icon.className = 'fas fa-volume-mute';
                span.textContent = '关闭';
                btn.classList.remove('sound-btn');
            }
        });
        
        // 手动刷新
        document.getElementById('refreshBtn').addEventListener('click', () => {
            this.fetchStocks();
        });
        
        // 刷新间隔选择
        document.getElementById('refreshInterval').addEventListener('change', (e) => {
            const interval = parseInt(e.target.value);
            this.setRefreshInterval(interval);
        });
        
        // 关闭模态框
        document.querySelectorAll('.close-modal').forEach(btn => {
            btn.addEventListener('click', () => {
                document.querySelectorAll('.modal').forEach(modal => {
                    modal.classList.remove('active');
                });
            });
        });
        
        // 点击模态框外部关闭
        document.querySelectorAll('.modal').forEach(modal => {
            modal.addEventListener('click', (e) => {
                if (e.target === modal) {
                    modal.classList.remove('active');
                }
            });
        });
        
        // 播报按钮
        document.addEventListener('click', (e) => {
            if (e.target.closest('.announce-btn')) {
                const btn = e.target.closest('.announce-btn');
                const name = btn.dataset.name;
                const price = btn.dataset.price;
                const change = btn.dataset.change;
                this.showAnnounceModal(name, price, change);
            }
            
            if (e.target.closest('.detail-btn')) {
                const btn = e.target.closest('.detail-btn');
                const code = btn.dataset.code;
                const market = btn.dataset.market;
                this.showDetailModal(code, market);
            }
        });
        
        // 播放语音
        document.getElementById('playAnnounce').addEventListener('click', () => {
            this.playAnnouncement();
        });
        
        // 停止语音
        document.getElementById('stopAnnounce').addEventListener('click', () => {
            this.stopAnnouncement();
        });
        
        // 音量控制
        document.getElementById('volume').addEventListener('input', (e) => {
            if (this.currentUtterance) {
                this.currentUtterance.volume = parseFloat(e.target.value);
            }
        });
    }
    
    connectWebSocket() {
        const protocol = window.location.protocol === 'https:' ? 'wss:' : 'ws:';
        const wsUrl = `${protocol}//${window.location.host}/ws`;
        
        this.socket = new WebSocket(wsUrl);
        
        this.socket.onopen = () => {
            console.log('WebSocket连接已建立');
            this.updateConnectionStatus('已连接到服务器', true);
        };
        
        this.socket.onmessage = (event) => {
            const data = JSON.parse(event.data);
            if (data.type === 'update') {
                this.updateStockDisplay(data.data);
                this.updateCount++;
                document.getElementById('updateCount').textContent = this.updateCount;
                document.getElementById('updateTime').textContent = 
                    new Date().toLocaleTimeString('zh-CN');
                
                // 检查是否需要语音播报
                this.checkForAnnouncements(data.data);
            }
        };
        
        this.socket.onerror = (error) => {
            console.error('WebSocket错误:', error);
            this.updateConnectionStatus('连接错误', false);
        };
        
        this.socket.onclose = () => {
            console.log('WebSocket连接已关闭');
            this.updateConnectionStatus('连接断开', false);
            // 5秒后尝试重连
            setTimeout(() => this.connectWebSocket(), 5000);
        };
    }
    
    updateConnectionStatus(text, connected) {
        const statusEl = document.getElementById('connectionStatus');
        const indicator = document.querySelector('.status-indicator');
        
        statusEl.textContent = text;
        
        if (connected) {
            indicator.className = 'status-indicator connected';
        } else {
            indicator.className = 'status-indicator';
            indicator.style.backgroundColor = '#e74c3c';
        }
    }
    
    async fetchStocks() {
        try {
            const response = await fetch('/api/stocks');
            const data = await response.json();
            
            if (data.success) {
                this.updateStockDisplay(data.data);
                document.getElementById('updateTime').textContent = 
                    new Date().toLocaleTimeString('zh-CN');
            }
        } catch (error) {
            console.error('获取股票数据失败:', error);
        }
    }
    
    updateStockDisplay(stocks) {
        // 这里可以根据需要实现动态更新DOM
        // 由于我们使用服务器端渲染，这里主要处理WebSocket更新
        console.log('收到股票数据更新:', stocks);
        
        // 更新每个股票卡片
        stocks.forEach(stock => {
            const card = document.querySelector(`[data-code="${stock.code}"][data-market="${stock.market}"]`);
            if (card) {
                const parentCard = card.closest('.stock-card');
                const priceEl = parentCard.querySelector('.price-main');
                const changeEl = parentCard.querySelector('.change-amount');
                const percentEl = parentCard.querySelector('.change-percent');
                
                // 更新价格
                priceEl.textContent = stock.latest_price.toFixed(2);
                
                // 更新涨跌额
                const changeSign = stock.change > 0 ? '+' : '';
                changeEl.textContent = `${changeSign}${stock.change.toFixed(2)}`;
                
                // 更新涨跌幅
                const percentSign = stock.change_pct > 0 ? '+' : '';
                percentEl.textContent = `${percentSign}${stock.change_pct.toFixed(2)}%`;
                
                // 更新颜色类
                parentCard.className = 'stock-card';
                if (stock.change > 0) {
                    parentCard.classList.add('positive');
                } else if (stock.change < 0) {
                    parentCard.classList.add('negative');
                } else {
                    parentCard.classList.add('neutral');
                }
            }
        });
    }
    
    checkForAnnouncements(stocks) {
        if (!this.soundEnabled) return;
        
        stocks.forEach(stock => {
            // 如果涨跌幅超过5%，进行语音播报
            if (Math.abs(stock.change_pct) >= 5) {
                const message = `${stock.name}当前价格${stock.latest_price.toFixed(2)}元，` +
                               `${stock.change_pct > 0 ? '上涨' : '下跌'}${Math.abs(stock.change_pct).toFixed(2)}个百分点`;
                this.speak(message);
            }
        });
    }
    
    showAnnounceModal(name, price, change) {
        document.getElementById('announceText').textContent = 
            `${name}当前价格${parseFloat(price).toFixed(2)}元，` +
            `${parseFloat(change) > 0 ? '上涨' : '下跌'}${Math.abs(parseFloat(change)).toFixed(2)}个百分点`;
        
        document.getElementById('announceModal').classList.add('active');
    }
    
    showDetailModal(code, market) {
        // 这里可以添加获取详细信息的API调用
        document.getElementById('detailContent').innerHTML = `
            <h3>${code}.${market.toUpperCase()} 详细信息</h3>
            <p>正在加载详细数据...</p>
        `;
        
        document.getElementById('detailModal').classList.add('active');
        
        // 模拟获取详细信息
        setTimeout(() => {
            document.getElementById('detailContent').innerHTML = `
                <h3>${code}.${market.toUpperCase()} 详细信息</h3>
                <div class="detail-grid">
                    <div class="detail-item">
                        <span class="detail-label">股票代码:</span>
                        <span class="detail-value">${code}</span>
                    </div>
                    <div class="detail-item">
                        <span class="detail-label">市场:</span>
                        <span class="detail-value">${market.toUpperCase()}</span>
                    </div>
                    <div class="detail-item">
                        <span class="detail-label">更新时间:</span>
                        <span class="detail-value">${new Date().toLocaleString('zh-CN')}</span>
                    </div>
                </div>
                <p class="detail-note">更多详细图表和分析功能正在开发中...</p>
            `;
        }, 500);
    }
    
    playAnnouncement() {
        const text = document.getElementById('announceText').textContent;
        this.speak(text);
    }
    
    stopAnnouncement() {
        if (this.speechSynth.speaking) {
            this.speechSynth.cancel();
        }
    }
    
    speak(text) {
        if (!this.soundEnabled) return;
        
        this.stopAnnouncement();
        
        const utterance = new SpeechSynthesisUtterance(text);
        utterance.lang = 'zh-CN';
        utterance.volume = parseFloat(document.getElementById('volume').value);
        utterance.rate = 1.0;
        utterance.pitch = 1.0;
        
        this.currentUtterance = utterance;
        this.speechSynth.speak(utterance);
    }
    
    setRefreshInterval(interval) {
        document.getElementById('currentInterval').textContent = interval;
        
        // 发送设置到服务器（如果需要）
        fetch('/api/settings', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ refresh_interval: interval })
        }).catch(console.error);
    }
    
    updateServerTime() {
        document.getElementById('serverTime').textContent = 
            new Date().toLocaleTimeString('zh-CN');
    }
}

// 页面加载完成后初始化应用
document.addEventListener('DOMContentLoaded', () => {
    window.stockApp = new StockApp();
});