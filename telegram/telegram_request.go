package telegram

import (
	"fmt"
	"gitlab-notificatons/config"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func sendTelegramRequest(endpoint string, params url.Values) error {
	apiURL := fmt.Sprintf("https://api.telegram.org/bot%s/%s", config.TelegramBotToken, endpoint)

	resp, err := http.Post(apiURL, "application/x-www-form-urlencoded", strings.NewReader(params.Encode()))
	if err != nil {
		log.Printf("❌ Ошибка при отправке запроса в Telegram: %v", err)
		return err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Printf("❌ Telegram API вернул ошибку: %s, %s", resp.Status, apiURL)
		log.Printf("📦 Ответ Telegram: %s", string(body)) // <-- ЛОГ ТЕЛА ОШИБКИ
		return fmt.Errorf("не удалось отправить запрос в Telegram. Статус: %s", resp.Status)
	}

	log.Printf("✅ Запрос успешно отправлен в Telegram (endpoint: %s)", endpoint)
	return nil
}
