package handlers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/shortener/internal/app/storage"
)

func TestSaveURL(t *testing.T) {
	// определяем структуру теста
	type want struct {
		code     int
		response string
	}
	// создаём массив тестов: имя и желаемый результат
	tests := []struct {
		name string
		url  string
		want want
	}{
		// определяем все тесты
		{
			name: "[Positive] Сохранение URL - получаю 200; данные сохранены",
			url:  "https://www.yandex.ru",
			want: want{
				code: http.StatusCreated,
			},
		},
	}
	for _, tt := range tests {
		// запускаем каждый тест
		t.Run(tt.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, "/", bytes.NewBuffer([]byte(tt.url)))
			request.Header.Set("Content-Type", "application/json")
			// создаём новый Recorder
			w := httptest.NewRecorder()
			// определяем хендлер
			r := gin.New()
			storage := storage.NewStorage()
			rg := NewRouterGroup(&r.RouterGroup, &storage)
			rg.Routes()

			// запускаем сервер
			r.ServeHTTP(w, request)
			res := w.Result()

			// проверяем код ответа
			if res.StatusCode != tt.want.code {
				t.Errorf("Expected status code %d, got %d", tt.want.code, w.Code)
			}

			assert.Truef(t, func() bool {
				for _, v := range storage.URL {
					if v.URL == tt.url {
						return true
					}
				}
				return false
			}(), "Can't save metric")

			// получаем и проверяем тело запроса
			defer res.Body.Close()
		})
	}
}
