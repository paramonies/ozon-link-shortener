package controller

import (
	"bytes"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
	"github.com/paramonies/ozon-link-shortener/internal/app/mock"
	"github.com/paramonies/ozon-link-shortener/internal/app/model"
)

func TestController_shortLink(t *testing.T) {
	type mockBehaviorType func(s *mock.MockService, link model.ClientLink)

	tests := []struct {
		name                 string
		inputBody            string
		inputLink            model.ClientLink
		mockBehavior         mockBehaviorType
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK - link found",
			inputBody: `{"url": "http://test.ru"}`,
			inputLink: model.ClientLink{
				Url: "http://test.ru",
			},
			mockBehavior: func(s *mock.MockService, link model.ClientLink) {
				s.EXPECT().GetShortLink(link.Url).Return("beeLDlFcPz")
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"url":"beeLDlFcPz"}`,
		},
		{
			name:      "OK - link created",
			inputBody: `{"url": "http://test.ru"}`,
			inputLink: model.ClientLink{
				Url: "http://test.ru",
			},
			mockBehavior: func(s *mock.MockService, link model.ClientLink) {
				s.EXPECT().GetShortLink(link.Url).Return("")
				s.EXPECT().CreateLink(link.Url).Return(model.ClientLink{Url: "beeLDlFcPz"}, nil)
			},
			expectedStatusCode:   201,
			expectedResponseBody: `{"url":"beeLDlFcPz"}`,
		},
		{
			name:                 "Failure - Invalid response body",
			inputBody:            `{"urll":"http://test.ru"}`,
			inputLink:            model.ClientLink{},
			mockBehavior:         func(s *mock.MockService, link model.ClientLink) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid response body"}`,
		},
		{
			name:                 "Failure - Invalid http link format",
			inputBody:            `{"url":"test.ru"}`,
			inputLink:            model.ClientLink{},
			mockBehavior:         func(s *mock.MockService, link model.ClientLink) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid http link format"}`,
		},
		{
			name:      "Failure - Internal server error",
			inputBody: `{"url": "http://test.ru"}`,
			inputLink: model.ClientLink{
				Url: "http://test.ru",
			},
			mockBehavior: func(s *mock.MockService, link model.ClientLink) {
				s.EXPECT().GetShortLink(link.Url).Return("")
				s.EXPECT().CreateLink(link.Url).Return(model.ClientLink{Url: ""}, errors.New("internal server error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"internal server error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockService := mock.NewMockService(c)
			test.mockBehavior(mockService, test.inputLink)

			controller := NewController(mockService)
			router := gin.New()
			router.POST("/short", controller.shortLink)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/short", bytes.NewBufferString(test.inputBody))
			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}

func TestController_getLongLink(t *testing.T) {
	type mockBehaviorType func(s *mock.MockService, link model.ClientLink)

	tests := []struct {
		name                 string
		inputBody            string
		inputLink            model.ClientLink
		mockBehavior         mockBehaviorType
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "OK",
			inputBody: `{"url":"beeLDlFcPz"}`,
			inputLink: model.ClientLink{
				Url: "beeLDlFcPz",
			},
			mockBehavior: func(s *mock.MockService, link model.ClientLink) {
				s.EXPECT().GetLongLink(link.Url).Return("http://test.ru", nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"url":"http://test.ru"}`,
		},
		{
			name:                 "Failure - Invalid response body",
			inputBody:            `{"urll":"beeLDlFcPz"}`,
			inputLink:            model.ClientLink{},
			mockBehavior:         func(s *mock.MockService, link model.ClientLink) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid response body"}`,
		},
		{
			name:                 "Failure - Invalid short url id format",
			inputBody:            `{"url":"beeLDlFcPzZ"}`,
			inputLink:            model.ClientLink{},
			mockBehavior:         func(s *mock.MockService, link model.ClientLink) {},
			expectedStatusCode:   400,
			expectedResponseBody: `{"error":"invalid short url id format"}`,
		},
		{
			name:      "Failure - Internal server error",
			inputBody: `{"url":"beeLDlFcPz"}`,
			inputLink: model.ClientLink{
				Url: "beeLDlFcPz",
			},
			mockBehavior: func(s *mock.MockService, link model.ClientLink) {
				s.EXPECT().GetLongLink(link.Url).Return("", errors.New("internal server error"))
			},
			expectedStatusCode:   500,
			expectedResponseBody: `{"error":"internal server error"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockService := mock.NewMockService(c)
			test.mockBehavior(mockService, test.inputLink)

			controller := NewController(mockService)
			router := gin.New()
			router.POST("/long", controller.getLongLink)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/long", bytes.NewBufferString(test.inputBody))
			router.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
