package resolver

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"dns-speedtest/internal/config"
	"dns-speedtest/internal/server"
	"dns-speedtest/pkg/models"
)

var globalConfig *config.Config

// SetConfig 设置全局配置
func SetConfig(cfg *config.Config) {
	globalConfig = cfg
}

// RunTest 对所有DNS服务器运行测试
func RunTest(servers []server.DNSServer, repeat int) []models.DNSTestResult {
	var results []models.DNSTestResult
	
	if globalConfig == nil {
		globalConfig = &config.Config{
			TestDomain: "www.google.com",
			Timeout:    5,
		}
	}
	
	fmt.Printf("将测试 %d 个DNS服务器，每个重复 %d 次\n", len(servers), repeat)
	
	// 是否并发测试
	if globalConfig.Concurrent {
		results = runConcurrentTest(servers, repeat)
	} else {
		results = runSequentialTest(servers, repeat)
	}
	
	// 排序结果
	sortResults(results)
	
	return results
}

// runSequentialTest 顺序测试所有DNS服务器
func runSequentialTest(servers []server.DNSServer, repeat int) []models.DNSTestResult {
	var results []models.DNSTestResult
	
	for _, srv := range servers {
		fmt.Printf("测试 %s (%s)...\n", srv.ID, srv.IP)
		result := testSingleServer(srv, repeat)
		results = append(results, result)
	}
	
	return results
}

// runConcurrentTest 并发测试所有DNS服务器
func runConcurrentTest(servers []server.DNSServer, repeat int) []models.DNSTestResult {
	results := make([]models.DNSTestResult, len(servers))
	var wg sync.WaitGroup
	
	for i, srv := range servers {
		wg.Add(1)
		go func(index int, server server.DNSServer) {
			defer wg.Done()
			fmt.Printf("测试 %s (%s)...\n", server.ID, server.IP)
			results[index] = testSingleServer(server, repeat)
		}(i, srv)
	}
	
	wg.Wait()
	return results
}

// testSingleServer 测试单个DNS服务器
func testSingleServer(srv server.DNSServer, repeat int) models.DNSTestResult {
	var durations []time.Duration
	var failCount int
	var lastError string
	
	for i := 0; i < repeat; i++ {
		duration, err := testDNS(srv.IP, globalConfig.TestDomain)
		if err != nil {
			failCount++
			lastError = err.Error()
		} else {
			durations = append(durations, duration)
		}
	}
	
	// 计算成功率和平均响应时间
	successRate := 100.0
	if repeat > 0 {
		successRate = float64(repeat-failCount) / float64(repeat) * 100
	}
	
	var avgDuration time.Duration
	if len(durations) > 0 {
		var total time.Duration
		for _, d := range durations {
			total += d
		}
		avgDuration = total / time.Duration(len(durations))
	}
	
	return models.DNSTestResult{
		ServerName:   srv.ID,
		ServerIP:     srv.IP,
		AvgDuration:  avgDuration,
		SuccessRate:  successRate,
		ErrorMessage: lastError,
	}
}

// testDNS 测试单个DNS查询的响应时间
func testDNS(serverIP, domain string) (time.Duration, error) {
	timeout := time.Duration(globalConfig.Timeout) * time.Second
	
	r := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			d := net.Dialer{
				Timeout: timeout,
			}
			return d.DialContext(ctx, "udp", serverIP+":53")
		},
	}
	
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	
	start := time.Now()
	_, err := r.LookupHost(ctx, domain)
	duration := time.Since(start)
	
	if err != nil {
		return duration, err
	}
	
	return duration, nil
}

// sortResults 按响应时间对结果进行排序
func sortResults(results []models.DNSTestResult) {
	for i := 0; i < len(results); i++ {
		for j := i + 1; j < len(results); j++ {
			if results[i].AvgDuration > results[j].AvgDuration {
				results[i], results[j] = results[j], results[i]
			}
		}
	}
}