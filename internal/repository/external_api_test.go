package repository_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/karrless/em-interview/internal/models"
	"github.com/karrless/em-interview/internal/repository"
)

func TestGetSongInfo(t *testing.T) {
	expectedResponse := map[string]string{
		"releaseDate": "16.07.2006",
		"text":        "Ooh baby, don't you know I suffer?\\nOoh baby, can you hear me moan?\\nYou caught me under false pretenses\\nHow long before you let me go?\\n\\nOoh\\nYou set my soul alight\\nOoh\\nYou set my soul alight",
		"link":        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем параметры запроса
		query := r.URL.Query()
		if query.Get("song") != "Starlight" {
			t.Errorf("expected song=Starlight, got %s", query.Get("song"))
		}
		if query.Get("group") != "Muse" {
			t.Errorf("expected group=Muse, got %s", query.Get("group"))
		}

		// Устанавливаем статус OK и возвращаем тестовый ответ
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(expectedResponse)
	}))
	defer server.Close()
	// Настраиваем Song и r.config для теста
	testSong := &models.Song{
		Title: "Starlight",
		Group: "Muse",
	}

	repo := repository.NewExtarnalAPIRepository(&repository.ExternalAPIConfig{ExtarnalAPIURL: server.URL})

	// Выполняем тестируемую функцию
	err := repo.UpdateSongInfo(testSong)

	// Проверяем, что ошибок нет
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Проверяем значения в результате
	expectedSong := &models.Song{
		Title:       "Starlight",
		Group:       "Muse",
		ReleaseDate: "16.07.2006",
		Text:        "Ooh baby, don't you know I suffer?\\nOoh baby, can you hear me moan?\\nYou caught me under false pretenses\\nHow long before you let me go?\\n\\nOoh\\nYou set my soul alight\\nOoh\\nYou set my soul alight",
		Link:        "https://www.youtube.com/watch?v=Xsp3_a-PMTw",
	}

	if !reflect.DeepEqual(testSong, expectedSong) {
		t.Errorf("expected %+v, got %+v", expectedSong, testSong)
	}

	testSong = &models.Song{
		Title: "Я Русский",
		Group: "SHAMAN",
	}

	err = repo.UpdateSongInfo(testSong)
	if err != repository.ErrBadRequest {
		t.Errorf("expected ErrBadRequest, got %v", err)
	}

}
