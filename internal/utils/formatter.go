package utils

import (
	"fmt"
	"strings"
	"time"

	"dns-speedtest/pkg/models"
)

func FormatTime(input float64) string {
    if input < 100 {
        return fmt.Sprintf("%.2f", input)
    } else if input < 1000 {
        return fmt.Sprintf("%.1f", input)
    } else {
        return fmt.Sprintf("%.0f", input)
    }
}

// PrintResults 打印测试结果
func PrintResults(results []models.DNSTestResult) {
    fmt.Println("\n测试结果 (按响应时间排序):")
    fmt.Println(strings.Repeat("-", 80))
    fmt.Printf("%-20s %-15s %-15s %-10s %s\n", "服务器名称", "服务器IP", "平均响应时间", "成功率", "错误信息")
    fmt.Println(strings.Repeat("-", 80))
    
    for _, result := range results {
        errorMsg := ""
        if result.SuccessRate < 100 {
            errorMsg = result.ErrorMessage
            if len(errorMsg) > 30 {
                errorMsg = errorMsg[:27] + "..."
            }
        }
        
        fmt.Printf("%-20s %-15s %-15s %-10.1f%% %s\n",
            truncate(result.ServerName, 20),
            result.ServerIP,
            formatDuration(result.AvgDuration),
            result.SuccessRate,
            errorMsg)
    }
    fmt.Println(strings.Repeat("-", 80))
}

// truncate 截断字符串到指定长度
func truncate(s string, maxLen int) string {
    if len(s) <= maxLen {
        return s
    }
    return s[:maxLen-3] + "..."
}

// formatDuration 格式化时间间隔为可读形式
func formatDuration(d time.Duration) string {
    if d < time.Millisecond {
        return fmt.Sprintf("%.2f μs", float64(d.Microseconds()))
    }
    return fmt.Sprintf("%.2f ms", float64(d.Microseconds())/1000)
}