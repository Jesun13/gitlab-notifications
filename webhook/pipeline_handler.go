package webhook

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gitlab-notificatons/models"
	"gitlab-notificatons/telegram"
	"log"
)

type PipelineHandler struct{}

func (p *PipelineHandler) HandleEvent(event map[string]interface{}, topicKey string) error {
	var pipelineEvent models.PipelineEvent
	if err := mapstructure.Decode(event, &pipelineEvent); err != nil {
		return fmt.Errorf("ошибка при обработке merge_request: %v", err)
	}
	emoji := ""
	status := pipelineEvent.ObjectAttributes.Status
	message := fmt.Sprintf("<b>Pipeline событие</b> в проекте <b>%s</b> от <b>%s</b>\nСтатус: <b>%s</b>\nСсылка: %s", pipelineEvent.Project.Name, pipelineEvent.User.Name, status, pipelineEvent.ObjectAttributes.URL)

	fmt.Println(pipelineEvent)
	switch status {
	case "success":
		emoji = "✅"
	case "failed":
		emoji = "❌"
	default:
		emoji = "📢"
	}

	message = emoji + " " + message

	// Отправка уведомления в Telegram
	err := telegram.SendTelegramMessage(message, topicKey, nil)
	if err != nil {
		return err
	}

	log.Println("✅ Уведомление отправлено для Pipeline")
	return nil
}
