package webhook

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gitlab-notificatons/models"
	"gitlab-notificatons/telegram"
	"log"
	"strings"
)

type MergeRequestHandler struct{}

func (m *MergeRequestHandler) HandleEvent(event map[string]interface{}, topicKey string) error {
	var mrEvent models.MergeRequestEvent
	if err := mapstructure.Decode(event, &mrEvent); err != nil {
		return fmt.Errorf("ошибка при обработке merge_request: %v", err)
	}

	var emoji, message string
	user := mrEvent.User.Name
	projectName := mrEvent.Project.Name
	mrTitle := mrEvent.ObjectAttributes.Title
	mrURL := mrEvent.ObjectAttributes.URL
	reviewers := []string{}
	for _, reviewer := range mrEvent.Reviewers {
		reviewers = append(reviewers, reviewer.Name)
	}

	// Логика обработки разных действий merge request
	switch mrEvent.ObjectAttributes.Action {
	case "open":
		emoji = "🚀"
		message = fmt.Sprintf("%s <b>Новый Merge Request</b> от <b>%s</b> в проекте <b>%s</b>\n📄 %s\n%s, пора смотреть говнокод!\n🔗 %s", emoji, user, projectName, mrTitle, strings.Join(reviewers, " "), mrURL)
		err := telegram.SendTelegramMessage(message, topicKey)
		if err != nil {
			return err
		}
	case "merge":
		emoji = "✅"
		message = fmt.Sprintf("%s <b>Merge Request</b> от <b>%s</b> был <b>замержен</b> в проекте <b>%s</b>\n📄 %s\n🔗 %s", emoji, user, projectName, mrTitle, mrURL)
		err := telegram.SendTelegramMessage(message, topicKey)
		if err != nil {
			return err
		}
	case "close":
		emoji = "❌"
		message = fmt.Sprintf("%s <b>Merge Request</b> от <b>%s</b> был <b>закрыт</b> в проекте <b>%s</b>\n📄 %s\n🔗 %s", emoji, user, projectName, mrTitle, mrURL)
		err := telegram.SendTelegramMessage(message, topicKey)
		if err != nil {
			return err
		}
	default:
		emoji = "📢"
		message = fmt.Sprintf("%s <b>Merge Request</b> от <b>%s</b> (%s) в проекте <b>%s</b>\n📄 %s\n🔗 %s", emoji, user, mrEvent.ObjectAttributes.Action, projectName, mrTitle, mrURL)
		err := telegram.SendTelegramMessage(message, topicKey)
		if err != nil {
			return err
		}
	}

	log.Println("✅ Уведомление отправлено для Merge Request")
	return nil
}
