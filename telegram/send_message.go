package telegram

import (
	"encoding/json"
	"fmt"
	"gitlab-notificatons/config"
	"net/url"
)

type Button struct {
	Text string `json:"text"`
	URL  string `json:"url"`
}

func SendTelegramMessage(message string, topicKey string, mrURLs *[]Button) error {
	// Параметры для сообщения
	params := url.Values{}
	params.Set("chat_id", config.ChatID)
	params.Set("text", message)
	params.Set("parse_mode", "HTML")
	params.Set("message_thread_id", fmt.Sprintf("%d", config.Topics[topicKey]))

	// Если mrURLs не nil и содержит элементы, добавляем кнопки
	if mrURLs != nil && len(*mrURLs) > 0 {
		// Создаём клавиатуру с кнопками в правильном формате
		inlineKeyboard := struct {
			InlineKeyboard [][]Button `json:"inline_keyboard"`
		}{
			InlineKeyboard: [][]Button{
				*mrURLs, // Разворачиваем слайс кнопок
			},
		}

		// Преобразуем структуру в JSON
		inlineKeyboardJSON, err := json.Marshal(inlineKeyboard)
		if err != nil {
			return fmt.Errorf("ошибка сериализации клавиатуры: %v", err)
		}

		// Добавляем клавиатуру в параметры
		params.Set("reply_markup", string(inlineKeyboardJSON))
	}

	// Отправляем запрос
	return sendTelegramRequest("sendMessage", params)
}
