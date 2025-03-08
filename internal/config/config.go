package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config 保存程序配置
type Config struct {
    TestRepeat  int    `yaml:"test_repeat"`
    DNSListPath string `yaml:"dns_list_path"`
    TestDomain  string `yaml:"test_domain"`
    Timeout     int    `yaml:"timeout"`
    Concurrent  bool   `yaml:"concurrent"`
}

// 创建默认配置
func createDefaultConfig() *Config {
    return &Config{
        TestRepeat:  3,
        DNSListPath: "dns_list.csv",
        TestDomain:  "cloudflare.com",
        Timeout:     5,
        Concurrent:  false,
    }
}

// LoadConfig 从YAML文件加载配置，如果文件不存在则创建默认配置文件
func LoadConfig(path string) (*Config, error) {
    // 尝试读取配置文件
    data, err := os.ReadFile(path)
    
    // 如果文件不存在，创建默认配置文件
    if os.IsNotExist(err) {
        config := createDefaultConfig()
        
        // 确保目录存在
        dir := filepath.Dir(path)
        if dir != "." && dir != "" {
            if err := os.MkdirAll(dir, 0755); err != nil {
                return nil, err
            }
        }
        
        // 将配置序列化为YAML
        data, err = yaml.Marshal(config)
        if err != nil {
            return nil, err
        }
        
        // 写入配置文件
        if err := os.WriteFile(path, data, 0644); err != nil {
            return nil, err
        }
        
        return config, nil
    } else if err != nil {
        return nil, err
    }
    
    // 如果文件存在，解析它
    config := &Config{}
    if err := yaml.Unmarshal(data, config); err != nil {
        return nil, err
    }
    
    // 补充缺失的默认值
    if config.TestRepeat <= 0 {
        config.TestRepeat = 3
    }
    if config.DNSListPath == "" {
        config.DNSListPath = "dns_list.csv"
    }
    if config.TestDomain == "" {
        config.TestDomain = "www.google.com"
    }
    if config.Timeout <= 0 {
        config.Timeout = 5 // 默认超时5秒
    }
    
    return config, nil
}