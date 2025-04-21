package telegram

import (
	"fmt"
	"gitlab-notificatons/config"
	"log"
	"net/url"
)

func SendTelegramMessage(message string, topicKey string) error {
	params := url.Values{}
	params.Set("chat_id", config.ChatID)
	params.Set("text", message)
	params.Set("parse_mode", "HTML")
	params.Set("message_thread_id", fmt.Sprintf("%d", config.Topics[topicKey]))

	log.Default().Println(fmt.Sprintf("%d", config.Topics[topicKey]))
	return sendTelegramRequest("sendMessage", params)
}
