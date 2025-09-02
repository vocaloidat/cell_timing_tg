package main

import (
	"cainiao/config"
	"cainiao/gotg"
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// parseWalletInfo 解析钱包信息，只提取钱包名称和CELL代币余额
func parseWalletInfo(output string) string {
	lines := strings.Split(output, "\n")
	var walletName string
	var cellBalance string

	// 查找钱包名称
	walletRegex := regexp.MustCompile(`wallet:\s*(\S+)`)
	for _, line := range lines {
		if match := walletRegex.FindStringSubmatch(line); match != nil {
			walletName = match[1]
			break
		}
	}

	// 查找CELL代币余额 - 修改逻辑以适应实际输出格式
	// 在输出中，coins出现在ticker之前，所以需要向前查找
	for i, line := range lines {
		line = strings.TrimSpace(line)
		// 找到CELL ticker
		if strings.Contains(line, "ticker: CELL") {
			// 向前搜索coins值
			for j := i - 1; j >= 0; j-- {
				prevLine := strings.TrimSpace(lines[j])
				if strings.HasPrefix(prevLine, "coins:") {
					cellBalance = strings.TrimPrefix(prevLine, "coins:")
					cellBalance = strings.TrimSpace(cellBalance)
					break
				}
				// 如果遇到另一个token块，停止搜索
				if strings.Contains(prevLine, "ticker:") && !strings.Contains(prevLine, "CELL") {
					break
				}
			}
			break
		}
	}

	// 格式化输出
	if walletName == "" {
		walletName = "未知"
	}
	if cellBalance == "" {
		cellBalance = "0.0"
	}

	return fmt.Sprintf("钱包名称: %s\nCELL代币余额: %s", walletName, cellBalance)
}

func main() {
	//读取钱包名称
	Myconfig := config.LodConfig()
	if Myconfig == nil {
		log.Panic("读取配置文件失败")
	}
	// 定义要执行的命令及其参数
	cmd := exec.Command("/opt/cellframe-node/bin/cellframe-node-cli", "wallet", "info", "-net", "Backbone", "-w", Myconfig.Cellframe.WalletName)

	// 捕获命令的标准输出和标准错误
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}

	// 解析输出，提取钱包名称和CELL代币余额
	walletInfo := parseWalletInfo(string(output))
	fmt.Printf("优化后的输出:\n%s\n", walletInfo)

	gotg.MyGOTG(walletInfo)
}
