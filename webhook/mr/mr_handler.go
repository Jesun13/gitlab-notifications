package mr

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"gitlab-notificatons/models"
	"gitlab-notificatons/telegram"
	"html/template"
	"log"
	"strings"
)

//go:embed mr.gohtml
var templateFS embed.FS
var mrTemplates = template.Must(template.New("").ParseFS(templateFS, "mr.gohtml"))

type MergeRequestHandler struct{}

type MergeRequestTemplateData struct {
	Title        string
	Description  string
	ProjectName  string
	SourceBranch string
	TargetBranch string
	Author       string
	Reviewers    string
	Action       string
	Buttons      []telegram.Button
}

func (m *MergeRequestHandler) HandleEvent(event map[string]interface{}, topicKey string) error {
	eventBytes, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("ошибка при сериализации события: %v", err)
	}

	var mrEvent models.MergeRequestEvent
	if err := json.Unmarshal(eventBytes, &mrEvent); err != nil {
		return fmt.Errorf("ошибка при парсинге merge_request: %v", err)
	}

	if mrEvent.ObjectAttributes.Title == "" || mrEvent.ObjectAttributes.URL == "" {
		return fmt.Errorf("неполные данные Merge Request")
	}

	data := MergeRequestTemplateData{
		Title:        mrEvent.ObjectAttributes.Title,
		Description:  mrEvent.ObjectAttributes.Description,
		ProjectName:  mrEvent.Project.Name,
		SourceBranch: mrEvent.ObjectAttributes.SourceBranch,
		TargetBranch: mrEvent.ObjectAttributes.TargetBranch,
		Author:       mrEvent.User.Name,
		Action:       mrEvent.ObjectAttributes.Action,
	}

	if len(mrEvent.Reviewers) > 0 {
		data.Reviewers = strings.Join(func() []string {
			var names []string
			for _, r := range mrEvent.Reviewers {
				names = append(names, r.Name)
			}
			return names
		}(), ", ")
	}

	data.Buttons = []telegram.Button{
		{
			Text: "Смотреть MR",
			URL:  mrEvent.ObjectAttributes.URL,
		},
	}

	var tplName string
	switch data.Action {
	case "open":
		tplName = "open"
	case "merge":
		tplName = "merge"
	case "close":
		tplName = "close"
	default:
		tplName = "default"
	}

	var buf bytes.Buffer
	if err := mrTemplates.ExecuteTemplate(&buf, tplName, data); err != nil {
		return fmt.Errorf("ошибка рендеринга шаблона: %v", err)
	}

	if err := telegram.SendTelegramMessage(buf.String(), topicKey, &data.Buttons); err != nil {
		return fmt.Errorf("ошибка отправки сообщения: %v", err)
	}

	log.Println("✅ Уведомление отправлено для Merge Request")
	return nil
}
