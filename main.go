package main

import (
	"fmt"
	"log"
	"time"

	"dns-speedtest/internal/config"
	"dns-speedtest/internal/resolver"
	"dns-speedtest/internal/server"
	"dns-speedtest/internal/utils"
	"dns-speedtest/pkg/models"
)

func main() {
    fmt.Println("DNS测速工具 - Go版本")
    fmt.Println("初始化中...")

    // 加载配置
    cfg, err := config.LoadConfig("config.yaml")
    if (err != nil) {
        log.Fatalf("加载配置失败: %v", err)
    }
    
    // 设置全局配置
    resolver.SetConfig(cfg)

    // 加载DNS服务器列表
    dnsServers, err := server.LoadDNSList(cfg.DNSListPath)
    if (err != nil) {
        log.Fatalf("加载DNS服务器列表失败: %v", err)
    }
    
    fmt.Printf("成功加载 %d 个DNS服务器\n", len(dnsServers))

    var results []models.DNSTestResult
    
    // 显示测试菜单
    fmt.Println("\n请选择测试模式:")
    fmt.Println("1. 快速测试 (每个服务器测试1次)")
    fmt.Println("2. 平均值测试 (每个服务器测试多次取平均)")
    fmt.Println("3. 并发测试 (同时测试所有服务器)")
    
    menuSwitch := utils.GetUserInput("\n输入你的选项: ")
    
    startTime := time.Now()
    
    switch menuSwitch {
    case "1":
        fmt.Println("开始快速测试...")
        results = resolver.RunTest(dnsServers, 1)
    case "2":
        fmt.Printf("开始平均值测试，循环次数: %d...\n", cfg.TestRepeat)
        results = resolver.RunTest(dnsServers, cfg.TestRepeat)
    case "3":
        fmt.Printf("开始并发测试，循环次数: %d...\n", cfg.TestRepeat)
        cfg.Concurrent = true
        results = resolver.RunTest(dnsServers, cfg.TestRepeat)
    default:
        fmt.Println("未知选项，使用默认快速测试...")
        results = resolver.RunTest(dnsServers, 1)
    }
    
    elapsed := time.Since(startTime)
    
    utils.PrintResults(results)
    
    fmt.Printf("\n测试完成！总耗时: %.2f 秒\n", elapsed.Seconds())
    
    if utils.Confirm("按回车键退出，输入 s 将结果保存为CSV文件: ") {
        if err := utils.SaveResultsToCSV(results, "dns_test_results.csv"); err != nil {
            log.Fatalf("保存文件时发生错误: %v", err)
        }
    }
}