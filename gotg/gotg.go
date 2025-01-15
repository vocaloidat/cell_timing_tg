package gotg

import (
	"bytes"
	"cainiao/config"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const (
	botToken = ""
	chatID   = ""
)

type TelegramMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}

func sendMessage(token, chatID, message string) error {
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)

	msg := TelegramMessage{
		ChatID: chatID,
		Text:   message,
	}

	body, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to encode message: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

func MyGOTG(msg string) {
	message := msg

	err := sendMessage(config.Myconfig.Telegram.BotToken, config.Myconfig.Telegram.ChatID, message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "发送消息错误: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("发送消息完成")
}
