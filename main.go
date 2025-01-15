package main

import (
	"cainiao/config"
	"cainiao/gotg"
	"fmt"
	"log"
	"os/exec"
)

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

	// 打印命令输出
	fmt.Printf("Command Output: \n%s\n", string(output))

	gotg.MyGOTG(string(output))
}
