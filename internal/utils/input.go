package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetUserInput 从控制台读取用户输入
func GetUserInput(prompt string) string {
    fmt.Print(prompt)
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    return strings.TrimSpace(input)
}

// Confirm 询问用户确认
func Confirm(prompt string) bool {
    input := GetUserInput(prompt)
    return strings.ToLower(input) == "s"
}