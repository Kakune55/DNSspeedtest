package models

import (
	"time"
)

type DNSResult struct {
    Average float64 `json:"average"`
    StdDev   float64 `json:"std_dev"`
    Max      float64 `json:"max"`
    Min      float64 `json:"min"`
    ID       string  `json:"id"`
    IP       string  `json:"ip"`
}

// DNSTestResult 存储DNS服务器的测试结果
type DNSTestResult struct {
    ServerName   string
    ServerIP     string
    AvgDuration  time.Duration
    SuccessRate  float64
    ErrorMessage string
}