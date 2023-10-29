package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Nexadis/alice-skill/internal/api"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWebhook(t *testing.T) {
	// описываем ожидаемое тело ответа при успешном запросе
	successBody, err := json.Marshal(
		&api.API{
			Response: &api.Response{
				Text: api.CanNothing,
			},
			Version: api.Version,
		},
	)
	require.NoError(t, err)

	// описываем набор данных: метод запроса, ожидаемый код ответа, ожидаемое тело
	testCases := []struct {
		method       string
		expectedCode int
		expectedBody string
	}{
		{method: http.MethodGet, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodPut, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodDelete, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodPost, expectedCode: http.StatusOK, expectedBody: string(successBody)},
	}
	e := newServer()
	srv := httptest.NewServer(e.Server.Handler)
	defer srv.Close()

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			req := resty.New().R()
			req.Method = tc.method
			req.URL = srv.URL

			resp, err := req.Send()

			assert.NoError(t, err, "error making HTTP request")

			// вызовем хендлер как обычную функцию, без запуска самого сервера

			assert.Equal(t, tc.expectedCode, resp.StatusCode(), "Код ответа не совпадает с ожидаемым")
			// проверим корректность полученного тела ответа, если мы его ожидаем
			if tc.expectedBody != "" {
				// assert.JSONEq помогает сравнить две JSON-строки
				assert.JSONEq(t, tc.expectedBody, string(resp.Body()), "Тело ответа не совпадает с ожидаемым")
			}
		})
	}
}
