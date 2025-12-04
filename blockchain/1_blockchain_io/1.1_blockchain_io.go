package blockchainio

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"
)

// ========== 1.1 区块链 I/O 操作 ==========
//
// 区块链系统中有大量的 I/O 操作，主要包括：
// 1. 网络 I/O：P2P 通信、RPC 调用、WebSocket 连接
// 2. 磁盘 I/O：存储区块数据、状态数据库、日志文件
// 3. 数据流 I/O：交易流、事件流、区块同步
// 4. 内存 I/O：状态缓存、交易池管理
//
// Go 语言在 I/O 密集型应用方面表现优秀，特别适合区块链开发。

// BlockchainIODemo 演示区块链中的 I/O 操作
func BlockchainIODemo() {
	fmt.Println("========== 1.1 区块链 I/O 操作 ==========")
	fmt.Println()
	fmt.Println("区块链系统中有大量的 I/O 操作，主要包括：")
	fmt.Println("1. 网络 I/O：P2P 通信、RPC 调用、WebSocket 连接")
	fmt.Println("2. 磁盘 I/O：存储区块数据、状态数据库、日志文件")
	fmt.Println("3. 数据流 I/O：交易流、事件流、区块同步")
	fmt.Println("4. 内存 I/O：状态缓存、交易池管理")
	fmt.Println()

	demonstrateNetworkIO()
	demonstrateDiskIO()
	demonstrateStreamIO()
	demonstrateConcurrentIO()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 区块链是 I/O 密集型应用，需要处理大量网络和磁盘操作")
	fmt.Println("✅ Go 的 goroutine 和 channel 非常适合处理并发 I/O")
	fmt.Println("✅ Go 的标准库提供了强大的 I/O 支持")
	fmt.Println("✅ 通过并发 I/O 可以大幅提升区块链节点的性能")
	fmt.Println()
}

// ========== 1. 网络 I/O ==========

// Block 表示一个区块
type Block struct {
	Number       int64         `json:"number"`
	Hash         string        `json:"hash"`
	ParentHash   string        `json:"parentHash"`
	Timestamp    int64         `json:"timestamp"`
	Transactions []Transaction `json:"transactions"`
}

// Transaction 表示一笔交易
type Transaction struct {
	Hash  string `json:"hash"`
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
}

// demonstrateNetworkIO 演示网络 I/O 操作
func demonstrateNetworkIO() {
	fmt.Println("=== 1.1.1 网络 I/O：区块链节点的核心操作 ===")
	fmt.Println()

	fmt.Println("区块链节点需要处理多种网络 I/O：")
	fmt.Println("1. P2P 网络通信：与其他节点交换区块和交易")
	fmt.Println("2. RPC 调用：提供 JSON-RPC API 给客户端")
	fmt.Println("3. WebSocket 连接：实时推送链上事件")
	fmt.Println("4. HTTP 请求：查询区块、交易、账户状态")
	fmt.Println()

	// 示例 1: 并发查询多个节点的区块数据
	fmt.Println("示例 1: 并发查询多个区块链节点（模拟）")
	demonstrateConcurrentRPC()

	// 示例 2: WebSocket 事件监听
	fmt.Println("示例 2: WebSocket 事件流（模拟）")
	demonstrateWebSocketStream()

	fmt.Println()
}

// demonstrateConcurrentRPC 演示并发 RPC 调用
func demonstrateConcurrentRPC() {
	fmt.Println("   场景：同时查询多个节点获取最新区块")
	fmt.Println()

	// 模拟多个节点地址
	nodes := []string{
		"https://node1.example.com",
		"https://node2.example.com",
		"https://node3.example.com",
	}

	// 使用 channel 收集结果
	resultCh := make(chan Block, len(nodes))
	var wg sync.WaitGroup

	// 并发查询所有节点
	for i, node := range nodes {
		wg.Add(1)
		go func(nodeURL string, nodeID int) {
			defer wg.Done()
			// 模拟 RPC 调用（实际中会使用真实的 HTTP 请求）
			block := fetchBlockFromNode(nodeURL, nodeID)
			resultCh <- block
			fmt.Printf("   [节点 %d] 获取到区块 #%d\n", nodeID, block.Number)
		}(node, i+1)
	}

	// 等待所有查询完成
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// 收集结果
	var blocks []Block
	for block := range resultCh {
		blocks = append(blocks, block)
	}

	fmt.Printf("   成功从 %d 个节点获取区块数据\n", len(blocks))
	fmt.Println()
}

// fetchBlockFromNode 模拟从节点获取区块（实际中会使用 HTTP 请求）
func fetchBlockFromNode(nodeURL string, nodeID int) Block {
	// 模拟网络延迟
	time.Sleep(50 * time.Millisecond)
	return Block{
		Number:     int64(1000 + nodeID),
		Hash:       fmt.Sprintf("0x%x", nodeID),
		ParentHash: fmt.Sprintf("0x%x", nodeID-1),
		Timestamp:  time.Now().Unix(),
		Transactions: []Transaction{
			{Hash: fmt.Sprintf("tx%d", nodeID), From: "0x123", To: "0x456", Value: "100"},
		},
	}
}

// demonstrateWebSocketStream 演示 WebSocket 事件流
func demonstrateWebSocketStream() {
	fmt.Println("   场景：监听链上事件（新交易、新区块）")
	fmt.Println()

	// 模拟事件流
	eventCh := make(chan string, 10)

	// 模拟事件生产者
	go func() {
		events := []string{"新交易", "新区块", "合约事件", "状态更新"}
		for _, event := range events {
			eventCh <- event
			time.Sleep(100 * time.Millisecond)
		}
		close(eventCh)
	}()

	// 事件消费者
	fmt.Println("   监听事件流：")
	for event := range eventCh {
		fmt.Printf("   [事件] %s\n", event)
	}
	fmt.Println()
}

// ========== 2. 磁盘 I/O ==========

// demonstrateDiskIO 演示磁盘 I/O 操作
func demonstrateDiskIO() {
	fmt.Println("=== 1.1.2 磁盘 I/O：存储区块数据 ===")
	fmt.Println()

	fmt.Println("区块链节点需要大量磁盘 I/O：")
	fmt.Println("1. 存储区块数据：每个区块需要持久化到磁盘")
	fmt.Println("2. 状态数据库：存储账户余额、合约状态等")
	fmt.Println("3. 日志文件：记录节点运行日志")
	fmt.Println("4. 索引文件：快速查询区块和交易")
	fmt.Println()

	// 示例：写入区块数据到文件
	fmt.Println("示例：将区块数据写入文件（模拟）")
	demonstrateBlockStorage()

	fmt.Println()
}

// demonstrateBlockStorage 演示区块存储
func demonstrateBlockStorage() {
	block := Block{
		Number:     12345,
		Hash:       "0xabcdef123456",
		ParentHash: "0x1234567890ab",
		Timestamp:  time.Now().Unix(),
		Transactions: []Transaction{
			{Hash: "0xtx1", From: "0x111", To: "0x222", Value: "1000"},
			{Hash: "0xtx2", From: "0x333", To: "0x444", Value: "2000"},
		},
	}

	// 将区块序列化为 JSON
	blockJSON, err := json.MarshalIndent(block, "", "  ")
	if err != nil {
		fmt.Printf("   错误：序列化失败 %v\n", err)
		return
	}

	// 模拟写入文件（实际中会使用数据库）
	fmt.Printf("   区块 #%d 数据大小: %d 字节\n", block.Number, len(blockJSON))
	fmt.Println("   数据内容（前100字符）:")
	if len(string(blockJSON)) > 100 {
		fmt.Printf("   %s...\n", string(blockJSON)[:100])
	} else {
		fmt.Printf("   %s\n", string(blockJSON))
	}
	fmt.Println("   ✅ 区块数据已持久化到磁盘")
	fmt.Println()
}

// ========== 3. 数据流 I/O ==========

// demonstrateStreamIO 演示数据流 I/O
func demonstrateStreamIO() {
	fmt.Println("=== 1.1.3 数据流 I/O：处理连续的数据流 ===")
	fmt.Println()

	fmt.Println("区块链中的数据流：")
	fmt.Println("1. 交易流：持续接收和处理新交易")
	fmt.Println("2. 区块流：同步和验证新区块")
	fmt.Println("3. 事件流：监听智能合约事件")
	fmt.Println("4. 状态流：跟踪状态变化")
	fmt.Println()

	// 示例：处理交易流
	fmt.Println("示例：处理交易流")
	demonstrateTransactionStream()

	fmt.Println()
}

// demonstrateTransactionStream 演示交易流处理
func demonstrateTransactionStream() {
	// 交易输入 channel
	txCh := make(chan Transaction, 10)

	// 模拟交易生产者
	go func() {
		for i := 1; i <= 5; i++ {
			tx := Transaction{
				Hash:  fmt.Sprintf("0xtx%d", i),
				From:  fmt.Sprintf("0xfrom%d", i),
				To:    fmt.Sprintf("0xto%d", i),
				Value: fmt.Sprintf("%d", i*100),
			}
			txCh <- tx
			fmt.Printf("   [生产者] 产生交易: %s\n", tx.Hash)
			time.Sleep(50 * time.Millisecond)
		}
		close(txCh)
	}()

	// 交易处理者（验证、执行）
	fmt.Println("   处理交易流：")
	for tx := range txCh {
		// 模拟交易验证和处理
		fmt.Printf("   [处理者] 验证交易: %s (From: %s, To: %s, Value: %s)\n",
			tx.Hash, tx.From, tx.To, tx.Value)
		time.Sleep(30 * time.Millisecond)
		fmt.Printf("   [处理者] ✅ 交易已处理: %s\n", tx.Hash)
	}
	fmt.Println()
}

// ========== 4. 并发 I/O ==========

// demonstrateConcurrentIO 演示并发 I/O 操作
func demonstrateConcurrentIO() {
	fmt.Println("=== 1.1.4 并发 I/O：Go 的核心优势 ===")
	fmt.Println()

	fmt.Println("Go 的并发模型非常适合区块链 I/O：")
	fmt.Println("1. goroutine：轻量级，可以创建大量并发 I/O 操作")
	fmt.Println("2. channel：安全地在 goroutine 间传递数据")
	fmt.Println("3. select：多路复用，同时处理多个 I/O 操作")
	fmt.Println()

	// 示例：同时处理多种 I/O 操作
	fmt.Println("示例：同时处理网络 I/O、磁盘 I/O 和事件流")
	demonstrateMultiIO()

	fmt.Println()
}

// demonstrateMultiIO 演示多路 I/O 操作
func demonstrateMultiIO() {
	// 创建多个 channel
	networkCh := make(chan Block, 5)
	diskCh := make(chan Block, 5)
	eventCh := make(chan string, 5)

	// 模拟网络 I/O：接收新区块
	go func() {
		for i := 1; i <= 3; i++ {
			block := Block{Number: int64(i), Hash: fmt.Sprintf("0xblock%d", i)}
			networkCh <- block
			time.Sleep(100 * time.Millisecond)
		}
		close(networkCh)
	}()

	// 模拟磁盘 I/O：保存区块
	go func() {
		for i := 1; i <= 3; i++ {
			block := Block{Number: int64(i + 10), Hash: fmt.Sprintf("0xsaved%d", i)}
			diskCh <- block
			time.Sleep(80 * time.Millisecond)
		}
		close(diskCh)
	}()

	// 模拟事件流：产生事件
	go func() {
		events := []string{"新交易", "合约调用", "状态更新"}
		for _, event := range events {
			eventCh <- event
			time.Sleep(60 * time.Millisecond)
		}
		close(eventCh)
	}()

	// 使用 select 多路复用处理所有 I/O
	fmt.Println("   使用 select 同时处理多种 I/O：")
	networkDone := false
	diskDone := false
	eventDone := false

	for !networkDone || !diskDone || !eventDone {
		select {
		case block, ok := <-networkCh:
			if !ok {
				networkDone = true
			} else {
				fmt.Printf("   [网络 I/O] 接收到区块: #%d %s\n", block.Number, block.Hash)
			}
		case block, ok := <-diskCh:
			if !ok {
				diskDone = true
			} else {
				fmt.Printf("   [磁盘 I/O] 保存区块: #%d %s\n", block.Number, block.Hash)
			}
		case event, ok := <-eventCh:
			if !ok {
				eventDone = true
			} else {
				fmt.Printf("   [事件流] 收到事件: %s\n", event)
			}
		}
	}
	fmt.Println()
}

// ========== 5. 实际应用示例 ==========

// BlockchainNode 模拟区块链节点
type BlockchainNode struct {
	blockCh    chan Block
	txCh       chan Transaction
	eventCh    chan string
	peers      []string
	blockchain []Block
	mu         sync.RWMutex
}

// NewBlockchainNode 创建新的区块链节点
func NewBlockchainNode() *BlockchainNode {
	return &BlockchainNode{
		blockCh:    make(chan Block, 100),
		txCh:       make(chan Transaction, 1000),
		eventCh:    make(chan string, 100),
		peers:      []string{"peer1", "peer2", "peer3"},
		blockchain: make([]Block, 0),
	}
}

// Start 启动节点（处理各种 I/O）
func (n *BlockchainNode) Start() {
	fmt.Println("=== 实际应用：区块链节点的 I/O 处理 ===")
	fmt.Println()

	// 启动多个 goroutine 处理不同的 I/O 操作
	go n.handleNetworkIO()     // 网络 I/O：P2P 通信
	go n.handleTransactionIO() // 交易 I/O：接收和处理交易
	go n.handleBlockIO()       // 区块 I/O：同步和存储区块
	go n.handleEventIO()       // 事件 I/O：监听和广播事件

	// 模拟运行一段时间
	time.Sleep(500 * time.Millisecond)

	fmt.Println("   节点运行中，处理各种 I/O 操作...")
	fmt.Println("   ✅ 网络 I/O: 与 3 个节点保持连接")
	fmt.Println("   ✅ 交易 I/O: 交易池中有交易待处理")
	fmt.Println("   ✅ 区块 I/O: 同步最新区块")
	fmt.Println("   ✅ 事件 I/O: 监听链上事件")
	fmt.Println()
}

// handleNetworkIO 处理网络 I/O（P2P 通信）
func (n *BlockchainNode) handleNetworkIO() {
	for _, peer := range n.peers {
		// 模拟与每个节点建立连接
		go func(p string) {
			for {
				// 模拟接收区块
				time.Sleep(200 * time.Millisecond)
				// 实际中会从网络接收数据
			}
		}(peer)
	}
}

// handleTransactionIO 处理交易 I/O
func (n *BlockchainNode) handleTransactionIO() {
	for tx := range n.txCh {
		// 验证交易
		// 添加到交易池
		_ = tx
	}
}

// handleBlockIO 处理区块 I/O
func (n *BlockchainNode) handleBlockIO() {
	for block := range n.blockCh {
		n.mu.Lock()
		n.blockchain = append(n.blockchain, block)
		n.mu.Unlock()
		// 持久化到磁盘
	}
}

// handleEventIO 处理事件 I/O
func (n *BlockchainNode) handleEventIO() {
	for event := range n.eventCh {
		// 广播事件给订阅者
		_ = event
	}
}

// ========== 6. HTTP RPC 服务器示例 ==========

// startRPCServer 启动 RPC 服务器（演示 HTTP I/O）
func startRPCServer() {
	fmt.Println("=== HTTP RPC 服务器：处理客户端请求 ===")
	fmt.Println()

	http.HandleFunc("/rpc", func(w http.ResponseWriter, r *http.Request) {
		// 处理 JSON-RPC 请求
		w.Header().Set("Content-Type", "application/json")

		// 读取请求体
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "读取请求失败", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		// 解析 JSON-RPC 请求
		var rpcReq map[string]interface{}
		if err := json.Unmarshal(body, &rpcReq); err != nil {
			http.Error(w, "解析请求失败", http.StatusBadRequest)
			return
		}

		// 处理请求并返回响应
		response := map[string]interface{}{
			"jsonrpc": "2.0",
			"id":      rpcReq["id"],
			"result":  "处理成功",
		}

		json.NewEncoder(w).Encode(response)
	})

	fmt.Println("   RPC 服务器已启动（模拟）")
	fmt.Println("   端点: POST /rpc")
	fmt.Println("   支持的方法: eth_blockNumber, eth_getBlockByNumber, etc.")
	fmt.Println()
}

// ========== 辅助函数 ==========

// readBlockFromFile 从文件读取区块（演示文件 I/O）
func readBlockFromFile(filename string) (*Block, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var block Block
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&block); err != nil {
		return nil, err
	}

	return &block, nil
}

// writeBlockToFile 将区块写入文件（演示文件 I/O）
func writeBlockToFile(block *Block, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(block)
}
