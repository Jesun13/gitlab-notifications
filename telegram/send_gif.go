package telegram

import (
	"fmt"
	"gitlab-notificatons/config"
	"net/url"
)

func SendTelegramGif(gifURL string, topicKey string) error {
	params := url.Values{}
	params.Set("chat_id", config.ChatID)
	params.Set("animation", gifURL)
	params.Set("message_thread_id", fmt.Sprintf("%d", config.Topics[topicKey]))

	return sendTelegramRequest("sendAnimation", params)
}
