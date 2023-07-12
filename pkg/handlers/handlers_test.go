package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/temathc/news-aggregator/models"
)

func TestHandlers_ListPublication(t *testing.T) {
	time := time.Now()
	publ := []models.Publications{
		{
			Title:       "test_title",
			Description: "test_description",
			Link:        "test_link",
			PubTime:     time,
		},
	}
	publicationsJson, _ := json.Marshal(publ)

	type fields struct {
		publicationsRepo *MockpublicationsRepository
	}
	type args struct {
		w *httptest.ResponseRecorder
		r *http.Request
	}
	tests := []struct {
		name                string
		setUp               func(f *fields)
		args                args
		expectedRequestBody []byte
		expectedStatusCode  int
	}{
		{
			name: "success",
			setUp: func(f *fields) {
				f.publicationsRepo.EXPECT().GetPublicationsWithLimit(1).Return([]models.Publications{
					{
						Title:       "test_title",
						Description: "test_description",
						Link:        "test_link",
						PubTime:     time,
					},
				}, nil).Times(1)
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/publications?limit=1", nil),
			},
			expectedRequestBody: publicationsJson,
			expectedStatusCode:  200,
		},
		{
			name: "not_found",
			setUp: func(f *fields) {
				f.publicationsRepo.EXPECT().GetPublicationsWithLimit(1).Return([]models.Publications{}, nil).Times(1)
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/publications?limit=1", nil),
			},
			expectedStatusCode: 404,
		},
		{
			name: "failed_to_get_list_of_news",
			setUp: func(f *fields) {
				f.publicationsRepo.EXPECT().GetPublicationsWithLimit(1).Return([]models.Publications{}, errors.New("Error")).Times(1)
			},
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/publications?limit=1", nil),
			},
			expectedStatusCode: 500,
		},
		{
			name: "bad_request",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "/publications?limцапit=s", nil),
			},
			expectedStatusCode: 400,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				publicationsRepo: NewMockpublicationsRepository(ctrl),
			}
			if tt.setUp != nil {
				tt.setUp(&f)
			}
			h := NewPublication(f.publicationsRepo)
			h.ListPublication(tt.args.w, tt.args.r)

			assert.Equal(t, tt.expectedStatusCode, tt.args.w.Code)
			assert.Equal(t, tt.expectedRequestBody, tt.args.w.Body.Bytes())
		})
	}
}
