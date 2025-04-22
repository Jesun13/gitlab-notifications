package helper

import (
	"encoding/json"
	"fmt"
	"gitlab-notificatons/webhook"
	"gitlab-notificatons/webhook/mr"
	"net/http"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	var event map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "не удалось разобрать тело запроса", http.StatusBadRequest)
		return
	}

	// Определяем тип события
	eventKind := event["object_kind"].(string)
	// Получаем topic из заголовков
	topicKey := r.Header.Get("x-topic")

	// Создаем диспетчер событий
	dispatcher := NewEventDispatcher()
	dispatcher.RegisterHandler("merge_request", &mr.MergeRequestHandler{})
	dispatcher.RegisterHandler("pipeline", &webhook.PipelineHandler{})

	// Распределяем событие
	if err := dispatcher.Dispatch(eventKind, event, topicKey); err != nil {
		http.Error(w, fmt.Sprintf("ошибка при обработке события: %v", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
