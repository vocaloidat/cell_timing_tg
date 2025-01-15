package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Cellframe Cellframe `yaml:"Cellframe"`
	Telegram  Telegram  `yaml:"Telegram"`
}

type Cellframe struct {
	WalletName string `yaml:"walletName"`
}

type Telegram struct {
	BotToken string `yaml:"botToken"`
	ChatID   string `yaml:"chatID"`
}

// 配置对象
var Myconfig = &Config{
	Cellframe: Cellframe{
		WalletName: "",
	},
	Telegram: Telegram{
		BotToken: "",
		ChatID:   "",
	},
}

func LodConfig() *Config {
	// 打开 YAML 文件
	file, err := os.Open("./config.yaml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}

	defer file.Close()

	// 创建解析器
	decoder := yaml.NewDecoder(file)

	// 解析 YAML 数据
	err = decoder.Decode(&Myconfig)
	if err != nil {
		fmt.Println("Error decoding YAML:", err)
		return nil
	}
	fmt.Println("读取配置文件：")
	fmt.Printf("%+v \n", Myconfig)
	return Myconfig
}
