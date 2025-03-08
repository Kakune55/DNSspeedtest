package server

import (
	"encoding/csv"
	"net/http"
	"os"
	"strings"
)

// DNSServer represents a DNS server with an ID and IP address.
type DNSServer struct {
    ID string  // 服务器名称
    IP string  // IP地址
}

// LoadServersFromFile loads DNS servers from a local CSV file.
func LoadServersFromFile(filePath string) ([]DNSServer, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    // 添加这一行允许不同行有不同字段数
    reader.FieldsPerRecord = -1
    
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var servers []DNSServer
    for _, record := range records {
        if len(record) < 2 {
            continue // 跳过无效记录
        }
        
        // 添加主 IP 地址
        servers = append(servers, DNSServer{ID: record[0], IP: record[1]})
        
        // 如果有备用 IP 地址，也添加到列表中
        for i := 2; i < len(record); i++ {
            if record[i] != "" {
                backupID := record[0] + " (备用)"
                servers = append(servers, DNSServer{ID: backupID, IP: record[i]})
            }
        }
    }
    return servers, nil
}

// LoadServersFromURL loads DNS servers from an online CSV file.
func LoadServersFromURL(url string) ([]DNSServer, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    reader := csv.NewReader(resp.Body)
    // 添加这一行允许不同行有不同字段数
    reader.FieldsPerRecord = -1
    
    records, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var servers []DNSServer
    for _, record := range records {
        if len(record) < 2 {
            continue // 跳过无效记录
        }
        
        // 添加主 IP 地址
        servers = append(servers, DNSServer{ID: record[0], IP: record[1]})
        
        // 如果有备用 IP 地址，也添加到列表中
        for i := 2; i < len(record); i++ {
            if record[i] != "" {
                backupID := record[0] + " (备用)"
                servers = append(servers, DNSServer{ID: backupID, IP: record[i]})
            }
        }
    }
    return servers, nil
}

// LoadDNSList 从文件或URL加载DNS服务器列表
func LoadDNSList(path string) ([]DNSServer, error) {
    // 检查是否是URL
    if strings.HasPrefix(path, "http://") || strings.HasPrefix(path, "https://") {
        return LoadServersFromURL(path)
    }
    return LoadServersFromFile(path)
}