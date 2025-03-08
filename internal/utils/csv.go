package utils

import (
	"encoding/csv"
	"fmt"
	"os"

	"dns-speedtest/pkg/models"
)

// ReadCSV reads a CSV file and returns the records as a slice of string slices.
func ReadCSV(filePath string) ([][]string, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    // 允许每行有不同数量的字段
    reader.FieldsPerRecord = -1
    
    records, err := reader.ReadAll()
    if (err != nil) {
        return nil, err
    }

    return records, nil
}

// WriteCSV writes records to a CSV file.
func WriteCSV(filePath string, records [][]string) error {
    file, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    return writer.WriteAll(records)
}

// SaveResultsToCSV 将测试结果保存为CSV文件
func SaveResultsToCSV(results []models.DNSTestResult, filename string) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    
    writer := csv.NewWriter(file)
    defer writer.Flush()
    
    // 写入CSV头
    if err := writer.Write([]string{
        "服务器名称",
        "服务器IP",
        "平均响应时间(ms)",
        "成功率(%)",
        "错误信息",
    }); err != nil {
        return err
    }
    
    // 写入测试结果
    for _, result := range results {
        record := []string{
            result.ServerName,
            result.ServerIP,
            fmt.Sprintf("%.2f", float64(result.AvgDuration.Microseconds())/1000),
            fmt.Sprintf("%.1f", result.SuccessRate),
            result.ErrorMessage,
        }
        if err := writer.Write(record); err != nil {
            return err
        }
    }
    
    fmt.Printf("结果已保存到 %s\n", filename)
    return nil
}